package main

import (
	"fmt"
	"github.com/ActiveState/tail"
	"github.com/codegangsta/cli"
	"os"
	"github.com/Shopify/sarama"
)

func main() {
	app := cli.NewApp()
	app.Name = "tail-kafka"
	app.Version = "0.1"
	app.Usage = "Tail a file (like a log file) and send the output to a Kafka topic"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:      "tail",
			ShortName: "t",
			Usage:     "tail log file and send to kafka",
			Flags: []cli.Flag{
				cli.BoolFlag{"debug", "Print tail lines"},
				cli.StringFlag{"logdir", "/var/log/apache2/access_log", "log file (absolute path)"},
				cli.StringFlag{"server", "", "Kafka server location with port `localhost:9092`"},
				cli.StringFlag{"topic", "apache", "Kafka queue topic"},
			},
			Action: func(c *cli.Context) {
				run(c)
			},
		},
	}
	app.Run(os.Args)
}

func run(c *cli.Context) {
	client, err := sarama.NewClient("client_id", []string{c.String("server")}, sarama.NewClientConfig())
	if err != nil {
		panic(err)
	} else {
		fmt.Println("> connected")
	}
	defer client.Close()

	producer, err := sarama.NewProducer(client, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	t, err := tail.TailFile(c.String("logdir"), tail.Config{Follow: true})
	for line := range t.Lines {
		fmt.Println(line.Text)
		go sendLineToKafka(line.Text, producer, c.String("topic"), c.Bool("debug"))
	}
	if (err != nil) {
		panic(err)
	}
}

func sendLineToKafka(line string, producer *sarama.Producer, topic string, debug bool) {
	err := producer.SendMessage(topic, nil, sarama.StringEncoder(line))
	if err != nil {
		panic(err)
	} else {
		if debug {
			fmt.Println(line + " > message sent")
		}
	}
}

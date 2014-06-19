==
Tail-Kafka
==

Program to tail a file and send each line to a Kafka topic.  Written in go.

Example usage:

```sh
./tail-kafka t --debug --logdir /var/log/apache2/access_log --server localhost:9092 --topic apache --client "client_id"
> connected
::1 - - [16/May/2014:21:39:25 -0400] "HEAD / HTTP/1.1" 200 - > message sent
::1 - - [16/May/2014:22:40:02 -0400] "HEAD / HTTP/1.1" 200 - > message sent
```

Configuration:
```sh
./tail-kafka t -h
NAME:
   tail - tail log file and send to kafka

USAGE:
   command tail [command options] [arguments...]

DESCRIPTION:


OPTIONS:
   --debug                                      Print tail lines
   --logdir '/var/log/apache2/access_log'       log file (absolute path)
   --server                                     Kafka server location with port `localhost:9092`
   --topic 'apache'                             Kafka queue topic
   --client 'client_id'                         Client ID for Kafka
```

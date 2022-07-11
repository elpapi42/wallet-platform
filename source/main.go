package main

import (
	"os"
	"os/signal"
	"syscall"
	"wallet/source/infrastructure/gin"
	"wallet/source/infrastructure/kafka"
)

func main() {
	consumer := kafka.NewKafkaConsumer(
		[]string{"localhost:9095"},
		"monolog",
		"monolog-consumer",
	)
	server := gin.NewGinServer(8080)

	consumer.Start()
	kafka.InitWriter()
	server.Start()

	waitTerminationSignal()

	server.Stop()
	kafka.CloseWriter()
	consumer.Close()
}

func waitTerminationSignal() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}

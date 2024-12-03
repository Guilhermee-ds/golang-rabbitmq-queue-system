package main

import (
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

const queueName = "Service1Queue"

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	Channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer Channel.Close()

	_, err = Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/send", func(c *gin.Context) {
		msg := c.Query("msg")
		if msg == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Message is required"})
			return
		}

		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		}

		//send message for the queue
		err = Channel.Publish("", queueName, false, false, message)
		if err != nil {
			log.Printf("Failed to publish message: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish message"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
	})

	log.Fatal(router.Run(":8080"))
}

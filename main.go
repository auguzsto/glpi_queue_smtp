package main

import (
	"log"

	"glpi_queue_smtp/modules/queues"
	"glpi_queue_smtp/modules/smtp"

	"github.com/gofor-little/env"
)

func main() {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}

	fineshed := make(chan queues.Queue)
	queues := queues.FindAll(12)
	if len(queues) > 0 {
		for _, queue := range queues {
			queue := queue
			go func() {
				smtp.Send(queue.To, queue.Name, queue.BodyHTML, &queue)
				fineshed <- queue
			}()
		}

		for range queues {
			queue := <-fineshed
			log.Println("Enviado: ", queue.ID, " | Name: ", queue.Name)
		}
	}
}

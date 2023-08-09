package main

import (
	"fmt"
	"log"
	"time"

	"glpi_queue_smtp/modules/queues"
	"glpi_queue_smtp/modules/smtp"

	"github.com/gofor-little/env"
)

func main() {
	fmt.Println("Starting")
	if err := env.Load(".env"); err != nil {
		panic(err)
	}

	for {

		fineshed := make(chan queues.Queue)
		getQueue := queues.FindAll(12)
		start := time.Now()

		if len(getQueue) > 0 {
			for _, queue := range getQueue {
				queue := queue
				go func() {
					smtp.Send(queue.To, queue.Name, queue.BodyHTML, queue)
					fineshed <- queue
				}()
			}

			for range getQueue {
				queue := <-fineshed
				log.Println("Enviado: ", queue.ID, " | Name: ", queue.Name)
			}
		}
		elapsed := time.Since(start)
		queues.CreateCronTaskLogs(getQueue, elapsed)
		time.Sleep(time.Second * 15)
	}

}

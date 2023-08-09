package smtp

import (
	"glpi_queue_smtp/modules/queues"
	"html"
	"log"
	"net/smtp"

	"github.com/gofor-little/env"
)

func Send(to, subject, body string, queue *queues.Queue) {
	SMTPFROM := env.Get("SMTPFROM", "DEFAULT_VALUE")
	SMTPPASS := env.Get("SMTPPASS", "DEFAULT_VALUE")
	SMTPADDRESS := env.Get("SMTPADDRESS", "DEFAULT_VALUE")
	SMTPHOST := env.Get("SMTPHOST", "DEFAULT_VALUE")

	auth := smtp.PlainAuth("", SMTPFROM, SMTPPASS, SMTPHOST)
	msg := "From: " + queue.Sendername + "<" + SMTPFROM + ">" + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + " \n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
		html.UnescapeString(body)
	err := smtp.SendMail(SMTPADDRESS, auth, SMTPFROM, []string{to}, []byte(msg))
	if err != nil {
		queues.IncrementSentTryCaseErrorSmtp(queue)
		log.Printf("Smtp error: %s", err)
		return
	}

	queues.Fineshed(queue)
}

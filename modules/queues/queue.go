package queues

import (
	"glpi_queue_smtp/modules/database"
	"time"
)

type Queue struct {
	ID         int64
	Name       string
	ItemsID    int64
	From       string
	Sendername string
	To         string
	BodyHTML   string
	IsDeleted  string
	SentTime   time.Time
	SentTry    int64
}

func FindAll(lenght int) []Queue {
	db := database.Con()
	var queues []Queue
	rows, err := db.Query("SELECT id, NAME, items_id, sender, sendername, recipient, body_html, is_deleted, sent_time, sent_try FROM glpi_queuednotifications WHERE is_deleted = 0 LIMIT ?", lenght)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var q Queue
		rows.Scan(&q.ID, &q.Name, &q.ItemsID, &q.From, &q.Sendername, &q.To, &q.BodyHTML, &q.IsDeleted, &q.SentTime, &q.SentTry)
		queues = append(queues, q)
	}

	return queues
}

func Fineshed(queue *Queue) {
	db := database.Con()
	sentTime := time.Now().Format(time.DateTime)
	sentTry := queue.SentTry + 1
	_, err := db.Exec("UPDATE glpi_queuednotifications SET sent_try = ?, is_deleted = 1, sent_time = ? WHERE id = ?", sentTry, sentTime, &queue.ID)
	if err != nil {
		panic(err)
	}

}

func IncrementSentTryCaseErrorSmtp(queue *Queue) {
	db := database.Con()
	sentTry := queue.SentTry + 1
	_, err := db.Exec("UPDATE glpi_queuednotifications SET sent_try = ? WHERE id = ?", sentTry, &queue.ID)
	if err != nil {
		panic(err)
	}
}

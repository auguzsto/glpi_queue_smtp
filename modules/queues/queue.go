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

func FindAll(length int) []Queue {
	db := database.Con()
	var queues []Queue
	rows, err := db.Query("SELECT id, NAME, items_id, sender, sendername, recipient, body_html, is_deleted, sent_time, sent_try FROM glpi_queuednotifications WHERE is_deleted = 0 LIMIT ?", length)
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

func Fineshed(queue Queue) {
	db := database.Con()
	sentTime := time.Now().Format(time.DateTime)
	queue.SentTry++
	_, err := db.Exec("UPDATE glpi_queuednotifications SET sent_try = ?, is_deleted = 1, sent_time = ? WHERE id = ?", queue.SentTry, sentTime, queue.ID)
	defer db.Close()
	if err != nil {
		panic(err)
	}

}

func IncrementSentTryCaseErrorSmtp(queue Queue) {
	db := database.Con()
	queue.SentTry++
	_, err := db.Exec("UPDATE glpi_queuednotifications SET sent_try = ? WHERE id = ?", queue.SentTry, queue.ID)
	defer db.Close()
	if err != nil {
		panic(err)
	}
}

func CreateCronTaskLogs(queues []Queue, elapsed time.Duration) {
	currentTime := time.Now().Format(time.DateTime)
	lastIdCronTasksLogs := getLastIdCronTaskLogs()
	lastIdCronTasksLogs++
	db := database.Con()
	_, err := db.Exec("INSERT INTO glpi_crontasklogs (id, crontasks_id, crontasklogs_id, date, state, elapsed, volume, content) VALUES (?, 22, ?, ?, 2, ?, ?, 'Action completed, fully processed')", lastIdCronTasksLogs, lastIdCronTasksLogs, currentTime, elapsed.Seconds(), len(queues))
	defer db.Close()
	if err != nil {
		panic(err)
	}
}

func getLastIdCronTaskLogs() int64 {
	db := database.Con()
	var lastId int64
	rows, err := db.Query("SELECT id FROM glpi_crontasklogs ORDER BY id DESC LIMIT 1")
	defer db.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&lastId)
	}

	return lastId
}

func HealthCronTaskLogs() {
	db := database.Con()
	_, err := db.Exec("DELETE FROM glpi_crontasklogs WHERE crontasks_id = 22")
	defer db.Close()
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Minute * 4320) // three days.
}

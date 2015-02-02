package models

import ()

type Message struct {
	Id         int64  // message id
	UserId     int64  // user id
	Body       string // message body
	CommitTime int64  //
}

func CreateMessage(m *Message) int64 {
	db := database()
	result, err := db.Exec("INSERT INTO messages (user_id, message, commit_time) VALUES (?, ?, ?)", m.UserId, m.Body, m.CommitTime)

	defer db.Close()
	if err != nil {
		panic(err)
		return -1
	}

	id, _ := result.LastInsertId()
	return id
}

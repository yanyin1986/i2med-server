package models

const (
	FriendRequest = iota // 发起请求
	FriendAccepted
	FriendDeclined
) // FriendStatus

type Friend struct {
	Id         int64
	UserId     int64
	FriendId   int64
	Status     int32 // FriendStatus
	CreateTime int64
	ModifyTime int64
}

func CreateFriend(f *Friend) (int64, error) {
	db := database()
	defer db.Close()

	result, err := db.Exec("INSERT INTO friends (user_id, friend_id, status, create_time, modify_time) VALUES (?, ?, ?, ?, ?)", f.UserId, f.FriendId, f.Status, f.CreateTime, f.ModifyTime)
	id, err := result.LastInsertId()
	return id, err
}

func GetFriendsByUserId(userId int64) ([]Friend, error) {
	db := database()
	defer db.Close()

	rows, err := db.Query("SELECT id, friend_id, status, create_time, modify_time FROM friends WHERE user_id = ?", userId)
	defer rows.Close()

	array := []Friend{}
	for rows.Next() {
		var (
			id         int64
			friendId   int64
			status     int32
			createTime int64
			modifyTime int64
		)

		rows.Scan(&id, &friendId, &status, &createTime, &modifyTime)
		f := Friend{
			Id:         id,
			UserId:     userId,
			FriendId:   friendId,
			Status:     status,
			CreateTime: createTime,
			ModifyTime: modifyTime,
		}

		array = append(array, f)
	}

	return array, err
}

func GetFriendByUserIdAndFriendId(userId int64, friendId int64) (Friend, error) {
	db := database()
	defer closeDatabase(db)

	var (
		id         int64
		status     int32
		createTime int64
		modifyTime int64
	)
	row := db.QueryRow("SELECT id, status, create_time, modify_time FROM friends WHERE user_id = ? AND friend_id = ?", userId, friendId)
	err := row.Scan(&id, &status, &createTime, &modifyTime)
	f := Friend{
		Id:         id,
		UserId:     userId,
		FriendId:   friendId,
		Status:     status,
		CreateTime: createTime,
		ModifyTime: modifyTime,
	}
	return f, err
}

func UpdateFriend(f *Friend) error {
	db := database()
	defer db.Close()

	_, err := db.Exec("UPDATE friends SET status = ?, modify_time = ? WHERE id = ?)", f.Status, f.ModifyTime, f.Id)
	return err
}

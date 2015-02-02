package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	DB_HOST = "localhost"
	DB_PORT = "3306"
	DB_NAME = "i2med"
	DB_USER = "root"
	DB_PWD  = ""
)

type User struct {
	Id         int64
	Name       string
	Age        int32
	Hospital   Hospital
	Title      string
	Department string
	Province   string
	City       string
	Education  []Education
}

type Hospital struct {
	Name string
}

type Education struct {
	SchoolName string
	Collage    string
}

func init() {

}

func database() *sql.DB {
	conn := DB_USER + ":" + DB_PWD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Println(err)
		return nil
	}

	return db
}

func closeDatabase(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}

func CreateUser(u *User) (int64, error) {
	db := database()
	result, err := db.Exec("insert into users(name, age, title, department, province, city) values(?, ?, ?, ?, ?, ?)", u.Name, u.Age, u.Title, u.Department, u.Province, u.City)
	id, err := result.LastInsertId()
	db.Close()
	return id, err
}

func GetUserByName(name string) *User {
	return GetUserBy("name", name)
}

func GetUserById(id int64) *User {
	return GetUserBy("id", id)
}

func GetUserBy(key string, value interface{}) *User {
	db := database()
	var (
		id         int64
		name       string
		age        int32
		title      string
		department string
		province   string
		city       string
	)

	sql := "SELECT id, name, age, title, department, province, city FROM users WHERE " + key + "=?"

	err := db.QueryRow(sql, value).Scan(&id, &name, &age, &title, &department, &province, &city)
	if err != nil {
		panic(err)
		return nil
	}

	user := User{
		Id:         id,
		Name:       name,
		Age:        age,
		Title:      title,
		Department: department,
		Province:   province,
		City:       city,
	}

	return &user
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	router.GET("/question/", returnQuestion)
	router.POST("/question/", addQuestion)
	router.Run("localhost:8080")
}

func dbInit() (*sql.DB, error) {

	var db *sql.DB
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "Quizzie",
	}

	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, fmt.Errorf("error opening database: error %v", err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		return nil, fmt.Errorf("ping error %v", pingErr)
	}

	fmt.Println("Successfully connected to the Quizzie Database")

	return db, err
}

func returnQuestion(c *gin.Context) {

	db, err := dbInit()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM Questions WHERE user_id = ?", 1)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("Get request to returnQuestion recieved")

}

func addQuestion(c *gin.Context) {

	db, err := dbInit()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	question := c.PostForm("question")
	optionOne := c.PostForm("optionOne")
	optionTwo := c.PostForm("optionTwo")
	optionThree := c.PostForm("optionThree")
	optionFour := c.PostForm("optionFour")
	genre := c.PostForm("genre")

	_, dbErr := db.Exec("INSERT INTO Questions (question,optionOne,optionTwo,optionThree,optionFour,genre,user_id) VALUES (?,?,?,?,?,?,?)", question, optionOne, optionTwo, optionThree, optionFour, genre, 1)

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	fmt.Println("Question added successfully")

}

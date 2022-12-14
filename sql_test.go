package errflow

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type User struct {
	id   int    `sql:"id"`
	name string `sql:"name"`
}

type Post struct {
	id     int    `sql:"id"`
	userID string `sql:"userID"`
	by     string `sql:"by"`
	title  string `sql:"title"`
}

var client *http.Client = &http.Client{}

func test() (User, []Post, error) {
	var db *sqlx.DB
	var user User
	var posts []Post
	var logID int64
	flow := NewDBx(db)
	flow.
		Beginx().
		Get(&user, "SELECT * FROM usesr WHERE `name`=?;", "Bobby").
		Get(&posts, "SELECT * FROM posts WHERE `userID`=?;", user.id).
		ExecFlow("INSERT INTO logs").
		Run(func(vr *Result) { logID = vr.LastInsertId() }).
		Commit()

	if err := flow.Err(); err != nil {
		err = fmt.Errorf("failed to fetch posts: %v", err)
		return user, nil, err
	}

	log.Printf("Fetched posts with log: %d", logID)

	return user, posts, nil
}

package errflow

import (
	"fmt"
	"log"

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

func test() (User, []Post, error) {
	var db *sqlx.DB
	var user User
	var posts []Post
	var logID int64
	flow := DB(db)
	flow.
		Beginx().
		Get(&user, "SELECT * FROM usesr WHERE `name`=?;", "Bobby").
		Get(&posts, "SELECT * FROM posts WHERE `userID`=?;", user.id).
		ExecFlow("INSERT INTO logs").
		Run(func(vr *VirtualResult) { logID, _ = vr.LastInsertId() }).
		Commit()

	if err := flow.Err(); err != nil {
		err = fmt.Errorf("failed to fetch posts: %v", err)
		return user, nil, err
	}

	log.Printf("Fetched posts with log: %d", logID)

	return user, posts, nil
}

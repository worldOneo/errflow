# ErrFlow

This is an experiment. We are here to see, how delayed error handling could improve code quality or degrade maintainability.

Example:

```go
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
```

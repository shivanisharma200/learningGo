package restAPIs

import (
	"database/sql"
)

func GetAll(db *sql.DB) (posts []Post, err error) {
	rows, err := db.Query("SELECT * from posts")
	if err != nil {
		return []Post{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		_ = rows.Scan(&post.ID, &post.Title, &post.Body)
		posts = append(posts, post)
	}
	return posts, nil
}
func GetById(db *sql.DB, id string) (post Post, err error) {

	err = db.QueryRow("SELECT * from posts WHERE id=?", id).Scan(&post.ID, &post.Title, &post.Body)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}
func InsertRow(db *sql.DB, post Post) (err error) {
	_, err = db.Exec("INSERT INTO posts (id,title,body) VALUES(?,?,?)", post.ID, post.Title, post.Body)
	return err
}
func UpdateById(db *sql.DB, id string, post Post) (err error) {
	_, err = db.Exec("UPDATE posts SET title = ?, body = ? WHERE id = ?", post.Title, post.Body, id)
	return err
}
func DeleteById(db *sql.DB, id string) (err error){
	_, err = db.Exec("DELETE FROM posts WHERE id = ?", id)
	return err
	
}

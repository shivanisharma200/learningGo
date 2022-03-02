package restAPIs

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/post")
	log.Print(err)
}

// post structure

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// slice of posts
var Posts []Post

// defining functions
// getPosts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(Posts)
	// experimenting
	posts, err := GetAll(db)
	if err != nil{
		log.Println("error:", err)
		w.WriteHeader(500)
		_, _ = w.Write([]byte("something unexpected happened"))
		return
	}
	resp, _ := json.Marshal(posts)
	_, _ = w.Write(resp)
}
// GetPost
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	post, err := GetById(db, vars["id"])
	if err != nil{
		log.Println("error:", err)
		w.WriteHeader(500)
		_, _ = w.Write([]byte("something unexpected happened"))
		return
	}
	resp, _ := json.Marshal(post)
	_, _ = w.Write(resp)
}

// createPost
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// experimenting
	var post Post
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &post)
	err := InsertRow(db, post)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(500)
		_, _ = w.Write([]byte("something unexpected happened"))
		return
	}
	_, _ = w.Write([]byte("success"))
}

// updatePost
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var post Post
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &post)
	err := UpdateById(db, vars["id"], post)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(500)
		_, _ = w.Write([]byte("something unexpected happened"))
		return
	}
	_, _ = w.Write([]byte("success"))
}

// deletePost
func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	err := DeleteById(db, vars["id"])
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(500)
		_, _ = w.Write([]byte("something unexpected happened"))
		return
	}
	_, _ = w.Write([]byte("success"))

}

package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type post struct {
	Id     *string `json:"id"`
	Titulo string  `json:"titulo"`
	Image  *string `json:"image"`
	UserId string  `json:"userId"`
	User   *user   `json:"user"`
}

func getPosts(c *gin.Context) {
	var posts []post
	rows, err := db.Query("SELECT * from posts")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()
	for rows.Next() {
		var post post
		if err := rows.Scan(&post.Id, &post.Titulo, &post.Image, &post.UserId); err != nil {

			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		post.User.Id = &post.UserId
		post.User.Find()

		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, posts)
}

func createPost(c *gin.Context) {
	var post post

	if err := c.BindJSON(&post); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err := db.Exec("INSERT INTO posts(titulo,userId) values(?,?)", post.Titulo, post.UserId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "post creado",
	})
}

func getPostsByUserId(c *gin.Context) {
	// An album to hold data from the returned row.
	var post post
	var id string
	id = c.Param("id")
	row := db.QueryRow("SELECT * FROM posts WHERE userId = ?", id)

	if err := row.Scan(&post.Id, &post.Titulo, &post.Image, &post.UserId); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"err": err.Error(),
			})
			return

		}
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, post)

}

package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id     *string `json:"id"`
	Nombre string  `json:"nombre"`
	Correo string  `json:"correo"`
	Clave  string  `json:"clave"`
	Image  *string `json:"image"`
}

func (u *user) Find() {

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", u.Id)

	if err := row.Scan(u.Id, u.Nombre, u.Correo, u.Clave, u.Image); err != nil {

		fmt.Println(err)
	}
	// return fileds
}

func getUsers(c *gin.Context) {
	var users []user
	rows, err := db.Query("SELECT * from users")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()
	for rows.Next() {
		var user user
		if err := rows.Scan(&user.Id, &user.Nombre, &user.Correo, &user.Clave, &user.Image); err != nil {

			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func postUser(c *gin.Context) {
	var user user

	if err := c.BindJSON(&user); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err := db.Exec("INSERT INTO users(nombre,correo,clave)  values(?,?,?)", user.Nombre, user.Correo, user.Clave)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "usuario creado",
	})
}

func getUserByID(c *gin.Context) {
	// An album to hold data from the returned row.
	var user user
	var id string
	id = c.Param("id")
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	if err := row.Scan(&user.Id, &user.Nombre, &user.Correo, &user.Clave, &user.Image); err != nil {
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
	c.IndentedJSON(http.StatusOK, user)

}

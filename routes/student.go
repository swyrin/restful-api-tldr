package routes

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-client/db"
)

type StudentT struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func AddStudentRoutes(rg *gin.RouterGroup, client *db.PrismaClient, ctx context.Context) {
	rg.POST("/student", func(c *gin.Context) {
		var body StudentT

		_ = c.BindJSON(&body)

		created, err := client.Student.CreateOne(
			db.Student.ID.Set(body.ID),
			db.Student.Name.Set(body.Name),
			db.Student.Email.Set(body.Email),
		).Exec(ctx)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		c.IndentedJSON(http.StatusOK, &created)
	})

	rg.GET("/student/:id", func(c *gin.Context) {
		found, err := client.Student.FindUnique(
			db.Student.ID.Equals(c.Param("id")),
		).Exec(ctx)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		c.IndentedJSON(http.StatusOK, found)
	})

	rg.GET("/student/all", func(c *gin.Context) {
		foundStudents, err := client.Student.FindMany().Exec(ctx)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		c.JSON(http.StatusOK, foundStudents)
	})

	rg.PUT("/student/:id", func(c *gin.Context) {
		var body StudentT
		_ = c.BindJSON(&body)

		model, err := client.Student.FindUnique(
			db.Student.ID.Equals(c.Param("id")),
		).Update(
			db.Student.ID.Set(body.ID),
			db.Student.Name.Set(body.Name),
			db.Student.Email.Set(body.Email),
		).Exec(ctx)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		c.IndentedJSON(http.StatusOK, model)
	})

	rg.DELETE("/student/:id", func(c *gin.Context) {
		found, err := client.Student.FindUnique(
			db.Student.ID.Equals(c.Param("id")),
		).Delete().Exec(ctx)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		c.JSON(http.StatusOK, found)
	})
}

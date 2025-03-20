package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"

	"rest-client/prisma/db"
)

type StudentT struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func AddStudentRoutes(rg *gin.RouterGroup) {
	client := db.NewClient()
	_ = client.Prisma.Connect()

	defer func() {
		_ = client.Prisma.Disconnect()
	}()

	ctx := context.Background()

	rg.POST("/", func(c *gin.Context) {
		var body StudentT

		_ = c.BindJSON(&body)

		created, _ := client.Student.CreateOne(
			db.Student.ID.Set(body.ID),
			db.Student.Name.Set(body.Name),
			db.Student.Email.Set(body.Email),
		).Exec(ctx)

		/*
			_, _ = client.Student.CreateOne(
				db.Student.ID.Set("ITITIU21172"),
				db.Student.Name.Set("Pham Tien Dat"),
				db.Student.Email.Set("dat20036@gmail.com"),
			).Exec(ctx)

			_, _ = client.Student.CreateOne(
				db.Student.ID.Set("ITITIU21263"),
				db.Student.Name.Set("Nguyen Trong Nguyen"),
				db.Student.Email.Set("ITITIU21263@student.hcmiu.edu.vn"),
			).Exec(ctx)
		*/

		c.IndentedJSON(http.StatusOK, created)
	})

	rg.GET("/:id", func(c *gin.Context) {
		found, _ := client.Student.FindUnique(
			db.Student.ID.Equals(c.Param("id")),
		).Exec(ctx)

		c.IndentedJSON(http.StatusOK, found)
	})

	rg.GET("/", func(c *gin.Context) {
		foundStudents, _ := client.Student.FindMany().Exec(ctx)

		c.JSON(http.StatusOK, foundStudents)
	})

	rg.PUT("/:id", func(c *gin.Context) {
		var body StudentT
		_ = c.BindJSON(&body)

		model, _ := client.Student.FindUnique(
			db.Student.ID.Equals(c.Param("id")),
		).Update(
			db.Student.Name.Set(body.Name),
			db.Student.Email.Set(body.Email),
		).Exec(ctx)

		c.IndentedJSON(http.StatusOK, model)
	})

	rg.DELETE("/:id", func(c *gin.Context) {
		found, _ := client.Student.FindUnique(
			db.Student.ID.Equals(c.Param("id")),
		).Delete().Exec(ctx)

		c.JSON(http.StatusOK, found)
	})
}

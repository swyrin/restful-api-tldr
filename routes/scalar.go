package routes

import (
	"fmt"
	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
)

func AddScalarDocRoute(rg *gin.RouterGroup) {
	rg.GET("/docs", func(c *gin.Context) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "https://raw.githubusercontent.com/swyrin/restful-api-tldr/refs/heads/main/docs/openapi.yml",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "API Demo",
			},
			DarkMode: true,
		})

		if err != nil {
			fmt.Printf("%v", err)
		}

		c.Data(200, "text/html; charset=utf-8", []byte(htmlContent))
	})
}

package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Setup() {
	var routegroup = router.Group("/")

	AddStudentRoutes(routegroup)
	_ = router.Run(":7554")
}

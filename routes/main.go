package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"rest-client/db"
)

var router = gin.Default()

func Setup(db *db.PrismaClient, ctx context.Context) {
	var rg = router.Group("/")

	AddStudentRoutes(rg, db, ctx)
	AddScalarDocRoute(rg)

	_ = router.Run(":7554")
}

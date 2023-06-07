package route

import (
	"time"

	"github.com/averagedo/go-project-template/bootstrap"
	"github.com/averagedo/go-project-template/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRoute(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
}

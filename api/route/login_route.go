package route

import (
	"time"

	"github.com/averagedo/go-project-template/api/controller"
	"github.com/averagedo/go-project-template/bootstrap"
	"github.com/averagedo/go-project-template/domain"
	"github.com/averagedo/go-project-template/mongo"
	"github.com/averagedo/go-project-template/repository"
	"github.com/averagedo/go-project-template/usecase"
	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}

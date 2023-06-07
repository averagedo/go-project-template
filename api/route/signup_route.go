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

func NewSignupRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}

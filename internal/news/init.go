package news

import (
	"github.com/Mufidzz/bareksa-test/internal/news/delivery/rest"
	"github.com/Mufidzz/bareksa-test/internal/news/usecase"
	"github.com/Mufidzz/bareksa-test/internal/repository/postgre"
	"github.com/Mufidzz/bareksa-test/internal/repository/redis"
	"github.com/gin-gonic/gin"
)

type Domain struct {
	Usecase *usecase.Usecase
}

func StartHTTP(router *gin.Engine, postgre *postgre.Postgre, redisRepo *redis.Redis) *Domain {
	uc := usecase.New(&usecase.Repositories{
		NewsDataRepository:        postgre,
		NewsTopicDataRepository:   postgre,
		NewsTagDataRepository:     postgre,
		AssignNewsAssocRepository: postgre,
		NewsRedisRepository:       redisRepo,
	})

	httpHandler := rest.NewHTTP(router, uc, uc, uc)
	httpHandler.SetRoutes()

	return &Domain{
		Usecase: uc,
	}
}

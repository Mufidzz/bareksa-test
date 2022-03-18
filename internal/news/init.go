package news

import (
	"github.com/Mufidzz/bareksa-test/internal/news/delivery/rest"
	"github.com/Mufidzz/bareksa-test/internal/news/usecase"
	"github.com/Mufidzz/bareksa-test/internal/repository/postgre"
	"github.com/gin-gonic/gin"
)

type Domain struct {
	Usecase *usecase.Usecase
}

func StartHTTP(router *gin.Engine, postgre *postgre.Postgre) *Domain {
	uc := usecase.New(&usecase.Repositories{
		NewsDataRepository: postgre,
	})

	httpHandler := rest.NewHTTP(router, uc)
	httpHandler.SetRoutes()

	return &Domain{
		Usecase: uc,
	}
}

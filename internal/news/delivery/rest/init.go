package rest

import "github.com/gin-gonic/gin"

type Usecases struct {
	NewsDataUC
}

type HTTPHandler struct {
	router   *gin.Engine
	usecases Usecases
}

func NewHTTP(
	router *gin.Engine,
	newsDataUC NewsDataUC,
) *HTTPHandler {
	return &HTTPHandler{
		router: router,
		usecases: Usecases{
			NewsDataUC: newsDataUC,
		},
	}
}

func (handler *HTTPHandler) SetRoutes() {
	router := handler.router
	news := router.Group("/news")
	{
		news.GET("/", handler.HandleGetNews)
		news.GET("/:id", handler.HandleGetSingleNews)
		news.PUT("/:newsId", handler.HandleUpdateSingleNews)
		news.POST("/", handler.HandleCreateSingleNews)
		news.DELETE("/:newsId", handler.HandleDeleteSingleNews)
	}
}

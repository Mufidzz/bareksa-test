package rest

import "github.com/gin-gonic/gin"

type Usecases struct {
	NewsDataUC
	NewsTopicDataUC
	NewsTagDataUC
}

type HTTPHandler struct {
	router   *gin.Engine
	usecases Usecases
}

func NewHTTP(
	router *gin.Engine,
	newsDataUC NewsDataUC,
	newsTopicDataUC NewsTopicDataUC,
	newsTagDataUC NewsTagDataUC,
) *HTTPHandler {
	return &HTTPHandler{
		router: router,
		usecases: Usecases{
			NewsDataUC:      newsDataUC,
			NewsTopicDataUC: newsTopicDataUC,
			NewsTagDataUC:   newsTagDataUC,
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

	newsTopic := router.Group("/news-topic")
	{
		newsTopic.GET("/", handler.HandleGetNewsTopic)
		newsTopic.PUT("/", handler.HandleUpdateNewsTopic)
		newsTopic.POST("/", handler.HandleCreateNewsTopic)
		newsTopic.DELETE("/", handler.HandleDeleteNewsTopic)
	}

	newsTag := router.Group("/news-tag")
	{
		newsTag.GET("/", handler.HandleGetNewsTag)
		newsTag.PUT("/", handler.HandleUpdateNewsTag)
		newsTag.POST("/", handler.HandleCreateNewsTag)
		newsTag.DELETE("/", handler.HandleDeleteNewsTag)
	}
}

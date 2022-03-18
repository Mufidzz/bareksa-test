package main

import (
	"github.com/Mufidzz/bareksa-test/internal/news"
	"github.com/Mufidzz/bareksa-test/internal/repository/postgre"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	postgreRepo, err := postgre.New("user=tes_bareksa password=tes_bareksa dbname=tes_bareksa host=139.162.36.125 sslmode=disable")
	if err != nil {
		log.Printf("[DB Init] error initialize database, trace %v", err)
	}

	StartREST(postgreRepo)
}

func StartREST(pg *postgre.Postgre) {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	news.StartHTTP(router, pg)

	router.Run(":4456")
}

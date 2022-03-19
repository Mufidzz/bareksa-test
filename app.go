package main

import (
	"github.com/Mufidzz/bareksa-test/internal/news"
	"github.com/Mufidzz/bareksa-test/internal/repository/postgre"
	redisRepository "github.com/Mufidzz/bareksa-test/internal/repository/redis"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	postgreRepo, err := postgre.New("user=tes_bareksa password=tes_bareksa dbname=tes_bareksa host=139.162.36.125 sslmode=disable")
	if err != nil {
		log.Printf("[DB Init] error initialize database, trace %v", err)
	}

	redisRepo, err := redisRepository.New(redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	StartREST(postgreRepo, redisRepo)
}

func StartREST(pg *postgre.Postgre, redisRepo *redisRepository.Redis) {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	news.StartHTTP(router, pg, redisRepo)

	router.Run(":4456")
}

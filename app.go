package main

import (
	"Test_Bareksa/internal/repository/postgre"
	"Test_Bareksa/presentation"
	"fmt"
	"log"
)

func main() {
	postgreRepo, err := postgre.New("user=tes_bareksa password=tes_bareksa dbname=tes_bareksa host=139.162.36.125 sslmode=disable")
	if err != nil {
		log.Printf("[DB Init] error initialize database, trace %v", err)
	}

	res, err := postgreRepo.GetNews(presentation.Pagination{
		Offset: 0,
		Count:  2,
	}, &presentation.NewsFilter{
		Status: 1,
		Topics: []int{1, 2, 3},
	})

	fmt.Println(res, err)
}

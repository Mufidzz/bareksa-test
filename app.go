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

	res, err := postgreRepo.CreateBulkNews([]presentation.CreateBulkNewsRequest{
		{
			Title:   "A",
			Content: "B",
			Status:  1,
		},
		{
			Title:   "D",
			Content: "E",
			Status:  2,
		},
	})

	fmt.Println(res, err)
}

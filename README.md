# Bareksa THT

### Database Model
![alt text](https://i.ibb.co/QPgmkZg/Screenshot-7.png)

### HOW TO TEST

``` go test -cover -v ./... ```

### HOW TO RUN

1. Setup Postgre Connection String -- app.go
``` 
postgreRepo, err := postgre.New("user=tes_bareksa password=tes_bareksa dbname=tes_bareksa host=localhost sslmode=disable")
	if err != nil {
		log.Printf("[DB Init] error initialize database, trace %v", err)
	} 
```
3. Setup Redis Address app.go
``` 
redisRepo, err := redisRepository.New(redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}) 
```


5. Run Syntax
``` go run app.go ```

### API DOCUMENTATION
https://documenter.getpostman.com/view/5872118/UVsPPk7z

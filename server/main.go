package main

import (
	"fmt"
	"net/http"

	"log"

	"./router"
)

func main() {
	// uri := "mongodb+srv://admin:0uRsrOjvWTHMEKMx@gotodo1.6wst9.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	// uri := "mongodb://localhost:27017"
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	// if err != nil {
	// 	panic(err)
	// }

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// //ping the primary
	// if err := client.Ping(ctx, readpref.Primary()); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Successfully connected and pinged!")

	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package server

import (
	"net/http"
)

// const productManagerAddr = "localhost:8081"

func index(w http.ResponseWriter, r *http.Request) {
	// conn, err := grpc.Dial(
	// 	productManagerAddr,

	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// 	grpc.WithBlock(),
	// )
	// defer func() {
	// 	if connErr := conn.Close(); connErr != nil {
	// 		log.Fatalf("failed to close connection. error: %v", connErr)
	// 	}
	// }()

	// if err != nil {
	// 	log.Fatalf("failed to connect puroduct-manager. error: %v", err)
	// }

	// req := &hellopb.HelloRequest{
	// 	Name: "sample",
	// }

	// client := hellopb.NewGreetingServiceClient(conn)
	// res, err := client.Hello(context.Background(), req)
	// if err != nil {
	// 	log.Fatalf("failed client.Hello. error: %v", err)
	// }

	// fmt.Println(res.GetMessage())
	generateHTML(w, "", "layout", "header", "index", "footer")
}

package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	hellopb "yorushika-store/store-controller/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (ps *ProductServer) index(w http.ResponseWriter, r *http.Request) {
	productManagerAddr := fmt.Sprintf("%s:%s", ps.psc.Host, ps.psc.Port)

	conn, err := grpc.Dial(
		productManagerAddr,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	defer func() {
		if connErr := conn.Close(); connErr != nil {
			log.Fatalf("failed to close connection. error: %v", connErr)
		}
	}()

	if err != nil {
		log.Fatalf("failed to connect puroduct-manager. error: %v", err)
	}

	req := &hellopb.HelloRequest{
		Name: "sample",
	}

	client := hellopb.NewGreetingServiceClient(conn)
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("failed client.Hello. error: %v", err)
	}

	fmt.Println(res.GetMessage())
	generateHTML(w, "", "layout", "header", "index", "footer")
}

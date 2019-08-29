package main

import (
	"log"
	"net/http"
)

type InMemoryPlayStore struct{}

func (i *InMemoryPlayStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	// 用来封装PlayServer 函数

	server := &PlayerServer{&InMemoryPlayStore{}}
	handler := http.HandlerFunc(server.ServeHTTP)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}

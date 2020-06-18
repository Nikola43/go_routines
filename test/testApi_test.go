package test

import (
	"log"
	"testing"
	"time"
)

type Value struct {
	Param  string
}

func TestGetAll(t *testing.T) {

	url := "http://localhost:8080/reverse"

	a := Value{Param:"hola men"}

	i := 0
	start := time.Now()
	for i = 0; i < 10000; i++ {
		GetRequest(url, "", a)
	}
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
)

var ctx = context.Background()
var redisClient *redis.Client

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	redisClient = rdb

	http.HandleFunc("/insert", HandleInsertValue)
	http.HandleFunc("/get", HandleGetValue)

	http.ListenAndServe(":9000", nil)
}

func HandleGetValue(writer http.ResponseWriter, request *http.Request) {

	key, ok := request.URL.Query()["key"]
	if ok {
		val, _ := redisClient.Get(ctx, key[0]).Result()
		writer.Write([]byte(val))
		return
	}
}

func HandleInsertValue(writer http.ResponseWriter, request *http.Request) {

	key, ok := request.URL.Query()["key"]
	value, valueOk := request.URL.Query()["value"]
	if ok && valueOk {
		redisClient.Set(ctx, key[0], value[0], 0)
	}

	writer.Write([]byte(fmt.Sprintf("key %s already insert", key)))
	return
}

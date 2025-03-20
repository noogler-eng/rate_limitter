package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/noogler-eng/rate-limiter/limitter"
	"github.com/noogler-eng/rate-limiter/redisdb"
)


func main(){
	// connecting to the redis
	redisdb.InitRedis()
	defer redisdb.RedisClient.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var res_string string = "this is route for testing"
		// Encode writes the JSON encoding of v to the stream, 
		// followed by a newline character.
		json.NewEncoder(w).Encode(res_string)
	})

	myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a rate-limited endpoint"))
	})

	http.HandleFunc("/ping", limitter.RateLimitter(myHandler))

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
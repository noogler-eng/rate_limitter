package limitter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/noogler-eng/rate-limiter/redisdb"
	"github.com/redis/go-redis/v9"
)

var LimitInMinutes int64 = 1
var RequestLimit int64 = 5

// wrapping the rate limtter middleare with the req, res function or HandlerFunc
func RateLimitter(originalHandler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware ratelimiting")
		
		//func net.SplitHostPort(hostport string) (host string, port string, err error)
		// getting thr ip from this
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		fmt.Fprintf(w, "Client IP: %s", ip)

		ctx := context.Background()
		now := time.Now().Unix()
		minTimestamp := now - int64(LimitInMinutes*60)

		// func (c redis.cmdable) ZCount(ctx context.Context, key string, min string, max string) *redis.IntCmd
		count, err := redisdb.RedisClient.ZCount(ctx, ip, fmt.Sprintf("%d", minTimestamp), fmt.Sprintf("%d", now)).Result()
		if err != nil {
			http.Error(w, "Redis error", http.StatusInternalServerError)
			return
		}

		if (count < RequestLimit){
			// func (c redis.cmdable) ZAdd(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd
			// here score is for the sorting ....
			// member is for sorting...
			_, err = redisdb.RedisClient.ZAdd(ctx, ip, redis.Z{
				Score:  float64(now),
				Member: now,
			}).Result()
			if err != nil {
				http.Error(w, "Redis error", http.StatusInternalServerError)
				return
			}
	
			// Remove old entries beyond the limit window, from left to minimum
			redisdb.RedisClient.ZRemRangeByScore(ctx, ip, "-inf", fmt.Sprintf("%d", minTimestamp))
			redisdb.RedisClient.ZAdd(ctx, ip, )
			originalHandler.ServeHTTP(w, r);
		}else{
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]string{"error": "Rate limit exceeded"})
		}
	})
}


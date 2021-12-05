package main

import (
	"fmt"
	"time"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/gofiber/fiber/v2"
)

var cache ttlcache.SimpleCache = ttlcache.NewCache()

func main() {
	cache.SetTTL(time.Duration(10 * time.Second))

	app := fiber.New()

	app.Get("/test",)
	cache.Set("1", 6868985)
	fmt.Println(cache.Get("1"))
}

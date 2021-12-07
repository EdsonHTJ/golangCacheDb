package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/gofiber/fiber/v2"
)

var cache ttlcache.SimpleCache = ttlcache.NewCache()

type writeReq struct {
	Hash  string
	Name  string
	Value interface{}
}

func main() {
	cache.SetTTL(time.Duration(10 * time.Second))

	app := fiber.New()

	app.Post("/write", func(c *fiber.Ctx) error {
		// Get raw body from POST request
		c.Body() // user=john
		fmt.Println(c.Body())

		var req writeReq

		json.Unmarshal(c.Body(), &req)

		hash := sha256.Sum256(append([]byte(req.Name), []byte(req.Hash)...))
		cache.Set(hex.EncodeToString(hash[:]), req.Value)

		return nil
	})

	app.Post("/read", func(c *fiber.Ctx) error {
		// Get raw body from POST request
		c.Body() // user=john
		fmt.Println(c.Body())

		var req writeReq
		fmt.Println(string(c.Body()))
		json.Unmarshal(c.Body(), &req)

		hash := sha256.Sum256(append([]byte(req.Name), []byte(req.Hash)...))

		req.Value, _ = cache.Get(hex.EncodeToString(hash[:]))

		res, _ := json.Marshal(req)
		fmt.Println(string(res))

		c.Status(200).Write(res)
		return nil
	})

	app.Listen(":5000")
}

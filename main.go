package main

import (
	crand "crypto/rand"
	"log"
	"math"
	"math/big"
	"math/rand"

	server "github.com/meihei3/portfolio-server/app"
)

var (
	frontendContentsPath = "./local" // 本番環境では`-ldflags`フラグで上書きする
)

func main() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64()) // math/rand は seed 値が固定なので、生成する必要がある。

	config := server.WithFrontendContentsPathConfig(frontendContentsPath)
	app := server.NewAppWithConfig(config)
	app.Use(server.NewLoggerMiddleware())

	if err := app.Run(); err != nil {
		log.Fatalf("failed server: %v", err)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/j3yzz/snapbuy-api-gateway/pkg/auth"
	"github.com/j3yzz/snapbuy-api-gateway/pkg/config"
	"github.com/j3yzz/snapbuy-api-gateway/pkg/order"
	"github.com/j3yzz/snapbuy-api-gateway/pkg/product"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at load config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	if err = r.Run(c.Port); err != nil {
		log.Fatalln("Failed at run web server", err)
	}
}

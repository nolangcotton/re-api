package api

import (
	"log"

	"github.com/nolangcotton/resrvc/cmd/api/router"

	"github.com/gin-gonic/gin"
)

func Server() {
	r := gin.Default()

	// Get requests
	r.GET("/ping", router.Ping)
	r.GET("/get_rent_avg/:city/:state", router.GetForRentAverages)
	r.GET("/get_forsale_avg/:city/:state", router.GetForSaleAverages)

	// Post to scrape new city / town
	r.POST("/add/:city/:state", router.AddCity)
	r.SetTrustedProxies(nil)

	err := r.Run(":2001")
	if err != nil {
		log.Panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}

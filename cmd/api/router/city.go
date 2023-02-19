package router

import (
	"log"
	"strings"

	"github.com/nolangcotton/resrvc/cmd/service"
	"github.com/nolangcotton/resrvc/db"

	"github.com/gin-gonic/gin"
)

type respAvgRent struct {
	City         string  `json:"city"`
	Zipcode      string  `json:"zipcode"`
	Price        float32 `json:"price"`
	Bedrooms     float32 `json:"Bedrooms"`
	Bathrooms    float32 `json:"Bathrooms"`
	Sqft         float32 `json:"Sqft"`
	DaysOnMarket float32 `json:"DaysOnMarket"`
	YearBuilt    float32 `json:"YearBuilt"`
}

func AddCity(c *gin.Context) {

	city := strings.Title(c.Param("city"))
	state := strings.ToUpper(c.Param("state"))

	service.LoadRentals(city, state)
	service.LoadSales(city, state)

}

func GetAvgRent(c *gin.Context) {

	db := db.Conn()

	city := strings.Title(c.Param("city"))
	state := strings.ToUpper(c.Param("state"))

	avgRentalSql := `
	select 
		city, 
		zipcode, 
		round(avg(price), 2) as price, 
		round(avg(bedrooms), 2) as bedrooms, 
		round(avg(bathrooms), 2) as bathrooms, 
		round(avg(squarefootage), 2) as sqft, 
		round(avg(daysonmarket), 2) as days_on_market, 
		round(avg(yearbuilt), 2) as year_built
	from dwh.for_rent 
	where city = '$1' and state = '$2'
	group by city, zipcode
	order by price desc
	`

	rows, err := db.Query(avgRentalSql, city, state)
	if err != nil {
		log.Fatal("Failed to get rental data")
	}
	defer rows.Close()

	for rows.Next() {

	}

}

package router

import (
	"log"
	"strings"

	"github.com/nolangcotton/resrvc/cmd/service"
	"github.com/nolangcotton/resrvc/db"

	"github.com/gin-gonic/gin"
)

type avgRow struct {
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

func GetForRentAverages(c *gin.Context) {

	db := db.Conn()

	city := strings.Title(c.Param("city"))
	state := strings.ToUpper(c.Param("state"))

	resultSet := []avgRow{}
	var currRow avgRow

	avgRentalSql := `
	select 
		city, 
		zipcode, 
		round(avg(price), 2) as price, 
		round(avg(bedrooms), 2) as bedrooms, 
		round(avg(bathrooms)) as bathrooms, 
		round(avg(squarefootage), 2) as sqft, 
		round(avg(daysonmarket), 2) as days_on_market, 
		round(avg(yearbuilt), 2) as year_built
	from dwh.for_rent 
	where city = initcap($1) and state = upper($2)
	group by city, zipcode
	order by price desc
	`

	rows, err := db.Query(avgRentalSql, city, state)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&currRow.City, &currRow.Zipcode, &currRow.Price, &currRow.Bedrooms, &currRow.Bathrooms, &currRow.Sqft, &currRow.DaysOnMarket, &currRow.YearBuilt)
		if err != nil {
			log.Panic(err)
		}
		resultSet = append(resultSet, currRow)
	}

	c.JSON(200, resultSet)

}

func GetForSaleAverages(c *gin.Context) {

	db := db.Conn()

	city := strings.Title(c.Param("city"))
	state := strings.ToUpper(c.Param("state"))

	resultSet := []avgRow{}
	var currRow avgRow

	avgSalesSql := `
	select 
		city, 
		zipcode, 
		round(avg(price), 2) as price, 
		round(avg(bedrooms), 2) as bedrooms, 
		round(avg(bathrooms)) as bathrooms, 
		round(avg(squarefootage), 2) as sqft, 
		round(avg(daysonmarket), 2) as days_on_market, 
		round(avg(yearbuilt), 2) as year_built
	from dwh.for_sale 
	where city = initcap($1) and state = upper($2)
	group by city, zipcode
	order by price desc
	`

	rows, err := db.Query(avgSalesSql, city, state)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&currRow.City, &currRow.Zipcode, &currRow.Price, &currRow.Bedrooms, &currRow.Bathrooms, &currRow.Sqft, &currRow.DaysOnMarket, &currRow.YearBuilt)
		if err != nil {
			log.Panic(err)
		}
		resultSet = append(resultSet, currRow)
	}

	c.JSON(200, resultSet)

}

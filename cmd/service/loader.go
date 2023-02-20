package service

import (
	"log"

	"github.com/nolangcotton/resrvc/db"
)

func LoadRentals(city string, state string) {
	log.Println("Loading Rental Data")
	db := db.Conn()
	defer db.Close()
	for _, rental_listing := range GetRentalData(city, state, 500) {
		sqlStatement := `
		INSERT INTO dwh.for_rent
		(county, bedrooms, bathrooms, squarefootage, yearbuilt, lotsize, propertytype, price, listeddate, addressline1, city, state, zipcode, formattedaddress, lastseen, status, removeddate, daysonmarket, createddate, id, latitude, longitude)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
		`
		_, err := db.Exec(sqlStatement,
			rental_listing.County,
			rental_listing.Bedrooms,
			rental_listing.Bathrooms,
			rental_listing.SquareFootage,
			rental_listing.YearBuilt,
			rental_listing.LotSize,
			rental_listing.PropertyType,
			rental_listing.Price,
			rental_listing.ListedDate,
			rental_listing.AddressLine1,
			rental_listing.City,
			rental_listing.State,
			rental_listing.ZipCode,
			rental_listing.FormattedAddress,
			rental_listing.LastSeen,
			rental_listing.Status,
			rental_listing.RemovedDate,
			rental_listing.DaysOnMarket,
			rental_listing.CreatedDate,
			rental_listing.Id,
			rental_listing.Latitude,
			rental_listing.Longitude,
		)
		if err != nil {
			log.Println("Failed to insert row -> ", rental_listing.Id)
			log.Panic(err)
		}
	}
	log.Println("Loaded " + city + " Rentals successfully")
}

func LoadSales(city string, state string) {
	log.Println("Loading Sales Data")
	db := db.Conn()
	defer db.Close()
	for _, sale_listing := range GetSalesData(city, state, 500) {
		sqlStatement := `
		INSERT INTO dwh.for_sale
		(bedrooms, bathrooms, price, squarefootage, county, propertytype, addressline1, addressline2, city, state, zipcode, formattedaddress, lastseen, listeddate, status, removeddate, daysonmarket, createddate, yearbuilt, lotsize, id, latitude, longitude)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
		`
		_, err := db.Exec(sqlStatement,
			sale_listing.Bedrooms,
			sale_listing.Bathrooms,
			sale_listing.Price,
			sale_listing.SquareFootage,
			sale_listing.County,
			sale_listing.PropertyType,
			sale_listing.AddressLine1,
			sale_listing.AddressLine2,
			sale_listing.City,
			sale_listing.State,
			sale_listing.ZipCode,
			sale_listing.FormattedAddress,
			sale_listing.LastSeen,
			sale_listing.ListedDate,
			sale_listing.Status,
			sale_listing.RemovedDate,
			sale_listing.DaysOnMarket,
			sale_listing.CreatedDate,
			sale_listing.YearBuilt,
			sale_listing.LotSize,
			sale_listing.Id,
			sale_listing.Latitude,
			sale_listing.Longitude,
		)
		if err != nil {
			log.Println("Failed to insert row -> ", sale_listing.Id)
			log.Panic(err)
		}
	}
	log.Println("Loaded " + city + " Sales successfully")
}

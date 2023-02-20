package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type salesListing struct {
	Bedrooms         int     `json:"bedrooms"`
	Bathrooms        float32 `json:"bathrooms"`
	Price            int     `json:"price"`
	SquareFootage    int     `json:"squareFootage"`
	County           string  `json:"county"`
	PropertyType     string  `json:"propertyType"`
	AddressLine1     string  `json:"addressLine1"`
	AddressLine2     string  `json:"addressLine2"`
	City             string  `json:"city"`
	State            string  `json:"state"`
	ZipCode          string  `json:"zipCode"`
	FormattedAddress string  `json:"formattedAddress"`
	LastSeen         string  `json:"lastSeen"`
	ListedDate       string  `json:"listedDate"`
	Status           string  `json:"status"`
	RemovedDate      string  `json:"removedDate"`
	DaysOnMarket     int     `json:"daysOnMarket"`
	CreatedDate      string  `json:"createdDate"`
	YearBuilt        int     `json:"yearBuilt"`
	LotSize          int     `json:"lotSize"`
	Id               string  `json:"id"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
}

func GetSalesData(city string, state string, limit int) []salesListing {

	listings := []salesListing{}

	url := fmt.Sprintf("https://realty-mole-property-api.p.rapidapi.com/saleListings?city=%s&state=%s&limit=%d", city, state, limit)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", os.Getenv("RE_API_KEY"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("RE_API_HOST"))

	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		log.Panic(resErr)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	jsonErr := json.Unmarshal(body, &listings)
	if jsonErr != nil {
		log.Println(jsonErr)
		log.Println("Req Body: ", string(body))
		log.Panic(jsonErr)
	}
	return listings

}

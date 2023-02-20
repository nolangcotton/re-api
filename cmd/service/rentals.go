package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type rentalListing struct {
	County           string  `json:"county"`
	Bedrooms         int     `json:"bedrooms"`
	Bathrooms        float32 `json:"bathrooms"`
	SquareFootage    int     `json:"squareFootage"`
	YearBuilt        int     `json:"yearBuilt"`
	LotSize          int     `json:"lotSize"`
	PropertyType     string  `json:"propertyType"`
	Price            int     `json:"price"`
	ListedDate       string  `json:"listedDate"`
	AddressLine1     string  `json:"addressLine1"`
	City             string  `json:"city"`
	State            string  `json:"state"`
	ZipCode          string  `json:"zipCode"`
	FormattedAddress string  `json:"formattedAddress"`
	LastSeen         string  `json:"lastSeen"`
	Status           string  `json:"status"`
	RemovedDate      string  `json:"removedDate"`
	DaysOnMarket     int     `json:"daysOnMarket"`
	CreatedDate      string  `json:"createdDate"`
	Id               string  `json:"id"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
}

func GetRentalData(city string, state string, limit int) []rentalListing {

	listings := []rentalListing{}

	url := fmt.Sprintf("https://realty-mole-property-api.p.rapidapi.com/rentalListings?city=%s&state=%s&&limit=%d", city, state, limit)

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
		log.Println("Req Body:", string(body))
		log.Panic(jsonErr)
	}

	return listings

}

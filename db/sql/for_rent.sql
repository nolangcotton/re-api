/*
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
*/
CREATE TABLE DWH.FOR_RENT (
    County           varchar,
	Bedrooms         int,
	Bathrooms        float,
	SquareFootage    int,
	YearBuilt        int,
	LotSize          int,
	PropertyType     varchar,
	Price            int,
	ListedDate       varchar,
	AddressLine1     varchar,
	City             varchar,
	State            varchar,
	ZipCode          varchar,
	FormattedAddress varchar,
	LastSeen         varchar,
	Status           varchar,
	RemovedDate      varchar,
	DaysOnMarket     int,
	CreatedDate      varchar,
	Id               varchar,
	Latitude         float,
	Longitude        float
);

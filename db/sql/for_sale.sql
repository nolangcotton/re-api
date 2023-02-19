/*
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
	Latitude             float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
}
*/
CREATE TABLE DWH.FOR_SALE (
    Bedrooms         int,
	Bathrooms        float,
	Price            int,
	SquareFootage    int,
	County           varchar,
	PropertyType     varchar,
	AddressLine1     varchar,
	AddressLine2     varchar,
	City             varchar,
	State            varchar,
	ZipCode          varchar,
	FormattedAddress varchar,
	LastSeen         varchar,
	ListedDate       varchar,
	Status           varchar,
	RemovedDate      varchar,
	DaysOnMarket     int,
	CreatedDate      varchar,
	YearBuilt        int,
	LotSize          int,
	Id               varchar,
	Latitude         float,
	Longitude        float
);
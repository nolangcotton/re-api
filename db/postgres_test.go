package db

import (
	"log"
	"testing"
)

func TestConn(t *testing.T) {
	db := Conn()
	err := db.Ping()
	if err != nil {
		log.Panic(err)
	}
}

func TestCheckForSchema(t *testing.T) {
	var schemaCount int
	db := Conn()
	defer db.Close()

	count, err := db.Query("select count(*) from information_schema.schemata where schema_name = 'dwh' and catalog_name = 'realestate'")
	if err != nil {
		log.Panic(err)
	}

	for count.Next() {
		err := count.Scan(&schemaCount)
		if err != nil {
			log.Println(schemaCount)
			log.Panic(err)
		} else if schemaCount != 1 {
			createSchema := "create schema dwh"
			_, err := db.Exec(createSchema)
			if err != nil {
				log.Panic(err)
			}
			log.Println("Created dwh schema")
		} else {
			log.Println("dwh schema already exists")
		}
	}
}

func TestCheckForTables(t *testing.T) {

	var schema string
	var table string
	var tableCount int

	db := Conn()
	defer db.Close()

	rows, err := db.Query("select table_schema, table_name from information_schema.tables where table_schema = 'dwh'")
	if err != nil {
		log.Panic(err)
	}

	count, err := db.Query("select count(*) from information_schema.tables where table_schema = 'dwh'")
	if err != nil {
		log.Panic(err)
	}

	for count.Next() {
		err := count.Scan(&tableCount)
		if err != nil {
			log.Println(tableCount)
			log.Panic(err)
		} else if tableCount != 2 {
			createRentTable := `CREATE TABLE DWH.FOR_RENT (
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
				Longitude        float, 
				PRIMARY KEY (city, state, id)
			)`
			createSalesTable := `CREATE TABLE DWH.FOR_SALE (
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
				Longitude        float,
				PRIMARY KEY (city, state, id)
			)`
			_, err := db.Exec(createRentTable)
			if err != nil {
				log.Panic(err)
			}
			_, err2 := db.Exec(createSalesTable)
			if err2 != nil {
				log.Panic(err2)
			}
			log.Println("Created dwh tables successfully")
		} else {
			log.Println("dwh tables already exist")
		}
	}

	for rows.Next() {
		err := rows.Scan(&schema, &table)
		if err != nil {
			log.Panic(err)
		}
		log.Println("Table -> ", schema, ".", table, " Retrieved Successfully")
	}
	rows.Close()
}

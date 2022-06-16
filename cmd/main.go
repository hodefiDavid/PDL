package main

import (
	"PDL"
	"fmt"
	//"github.com/353solutions/pdl"
)

func main() {

	const testURL = "https://s3.amazonaws.com/nyc-tlc/trip+data/green_tripdata_2018-03.parquet"
	PDL.Download(testURL, "")

	fmt.Printf("Hello G")
}

package main

import (
	"PDL"
	//"github.com/353solutions/pdl"
)

const testURL = "https://s3.amazonaws.com/nyc-tlc/trip+data/green_tripdata_2018-03.parquet"

func main() {
	PDL.Download(testURL, "")
}

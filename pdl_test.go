package PDL

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadSize(t *testing.T) {
	const testURL = "https://s3.amazonaws.com/nyc-tlc/trip+data/green_tripdata_2018-03.parquet"
	size, err := downloadSize(testURL)
	require.NoError(t, err)
	/*
		if err != nil {
			t.Fatalf("%s: %s", testURL, err) // will exit current test
			// t.Errorf("%s: %s", testURL, err)
		}
	*/
	expected := int64(13304186)
	require.Equal(t, expected, size)
	/*
		// expected := 13304186 // compilation error
		if size != expected {
			t.Fatalf("size: got %d, expected %d", size, expected)
		}
	*/
}

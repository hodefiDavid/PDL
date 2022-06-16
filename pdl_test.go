package PDL

import "testing"

func TestDownloadSize(t *testing.T) {
	const testURL = ""
	size, err := downloadSize(testURL)
	if err != nil {
		t.Fatalf("%s : %s", testURL, err)
	}
	const expected = 133 //need to fix the num
	if size != expected {
		t.Fatalf("size got %d, expected %d", size, expected)
	}

}

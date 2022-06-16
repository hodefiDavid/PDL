package PDL

import (
	"fmt"
	"net/http"
)

func downloadSize(url string) (int64, error) {

	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("%s bad status - %s", url, resp.Status)

	}
	return 0, err //delete

}

func Download(url, fileName string) error {
	size, err := downloadSize(url)
	fmt.Println(size, err)
	// TODO
	return nil
}

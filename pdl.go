package PDL

import "net/http"

func downloadSize(url string) (int64, err error) {

	resp, err := http.head(url)
	if err != nil {
		return 0, err
	}

}

func download(url, fileName string) error {

	return nil
}

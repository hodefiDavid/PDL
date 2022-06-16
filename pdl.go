package PDL

import (
"fmt"
"io"
"net/http"
"os"
)

func downloadSize(url string) (int64, error) {
resp, err := http.Head(url)
if err != nil {
return 0, err
}

if resp.StatusCode != http.StatusOK {
return 0, fmt.Errorf("%s: bad status - %s", url, resp.Status)
}

return resp.ContentLength, nil
}

func Download(url, fileName string) error {
size, err := downloadSize(url)
fmt.Println(size, err)
// TODO
return nil
}

// createEmptyFile creates an empty file in given size
func createEmptyFile(path string, size int64) error {
file, err := os.Create(path)
if err != nil {
return err
}
defer file.Close()

file.Seek(size-1, io.SeekStart)
file.Write([]byte{0})

return nil
}

func downloadPart(url string, start, end int64, path string) error {
req, err := http.NewRequest(http.MethodGet, url, nil)
if err != nil {
return err
}

req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
resp, err := http.DefaultClient.Do(req)
}

// writeAt writes data a location of file
func writeAt(path string, offset int64, r io.Reader) error {
file, err := os.OpenFile(path, os.O_RDWR, 0755)
if err != nil {
return err
}
defer file.Close()

if _, err := file.Seek(offset, io.SeekStart); err != nil {
return err
}

_, err = io.Copy(file, r)
return err
}


// createEmptyFile creates an empty file in given size
func createEmptyFile(path string, size int64) error {
file, err := os.Create(path)
if err != nil {
return err
}
defer file.Close()

file.Seek(size-1, io.SeekStart)
file.Write([]byte{0})

return nil
}

}

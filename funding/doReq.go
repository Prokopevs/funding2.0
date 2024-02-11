package funding

import (
	"fmt"
	"io"
	"net/http"
)

func DoReq(url string) (content *[]byte) {
	resp, err := http.Get(url)
	if err != nil {
		// errStr := fmt.Sprintf("Error requesting data from %s: %s\n", url, err.Error())
		// errW.ErrorHandler(errStr)
		fmt.Println(err)

		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// errStr := fmt.Sprintf("Error reading body from %s: %s\n", *url, err.Error())
		// errW.ErrorHandler(errStr)
		fmt.Println(err)

		return nil
	}

	return &body
}

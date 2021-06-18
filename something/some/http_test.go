package some

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {
	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}

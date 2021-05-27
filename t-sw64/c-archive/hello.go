package main

import "fmt"
import "C"
import "net/http"
import "bytes"
import "io"

func main() {}

//export Hello
func Hello() {
	fmt.Println("Hello From Go")
}

//export Browser
func Browser(url string) *C.char {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b := bytes.NewBuffer(nil)
	io.Copy(b, resp.Body)
	return C.CString(b.String())
}

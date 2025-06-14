package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, Go Standard Library")
	resp, err := http.Get("http://tilltoss.net:5000/tahsil/getsettings2")

	if err != nil {
		fmt.Println("The error is", err)
	}
	defer resp.Body.Close()

	fmt.Println("This is response from GET request Status", resp.Status)
	fmt.Println("This is response from GET request StatusCode", resp.StatusCode)
	fmt.Println("This is response from GET request Body", resp.Body)
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//010 OMIT
func main() {
	fmt.Println("REST example")
	v := url.Values{}
	v.Set("a", "3")
	v.Add("b", "4")
	q := v.Encode() // HL

	resp, err := http.Get("http://localhost:8080/sum?" + q)
	if err != nil {
		log.Fatalf("could not GET from server: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("unable to read response body: %v", err)
	}
	fmt.Printf("the sum of 3 and 4 is: %s", string(body))
}

//020 OMIT

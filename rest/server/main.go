package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//030 OMIT
func sumStr(as, bs string) (int, error) {
	var (
		a, b int
		err  error
	)
	if a, err = strconv.Atoi(as); err != nil {
		return 0, fmt.Errorf("a param bad: %v", err)
	}
	if b, err = strconv.Atoi(bs); err != nil {
		fmt.Println(a)
		return 0, fmt.Errorf("b param bad: %v", err)
	}
	return a + b, nil
}

//040 OMIT

//010 OMIT
func main() {
	fmt.Println("REST server")
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		aStr := r.FormValue("a")
		bStr := r.FormValue("b")
		sum, err := sumStr(aStr, bStr) // HL
		if err != nil {
			fmt.Fprintf(w, "ERROR: %v", err)
			return
		}
		fmt.Fprintf(w, "%v", sum)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//020 OMIT

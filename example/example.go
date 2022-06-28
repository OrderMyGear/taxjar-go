package main

import (
	"fmt"
	"os"

	"github.com/OrderMyGear/taxjar-go"
)

func main() {
	c := taxjar.NewClient(os.Getenv("API_TOKEN"), os.Getenv("API_URI"))
	c.Debug = true

	// Get rates at specific ZIP with an optional city specifier
	// rate, err := c.Rates.Get("12901", taxjar.RateCity("Plattsburgh"))
	// if nil != err {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("%+v\n", rate)

	from := taxjar.Address{
		Street:  "2211 Commerce St.",
		City:    "Dallas",
		State:   "TX",
		Zip:     "75201",
		Country: "US",
	}
	to := taxjar.Address{
		State:   "TX",
		Zip:     "75206",
		Country: "US",
	}
	taxes, err := c.Taxes.Calculate(from, to, 10.00, 100.00)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", taxes)
}

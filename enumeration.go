package main

import "fmt"

type SiteAddress string

const (
	Site_Google SiteAddress = "google.com"
	Site_Bing   SiteAddress = "bing.com"
	Site_Yahoo  SiteAddress = "yahoo.com"
)

func (s SiteAddress) getIP() string {
	switch s {
	case Site_Google:
		return "4.2.2.4"
	case Site_Bing:
		return "6.6.6.6"
	case Site_Yahoo:
		return "9.5.4.8"
	default:
		return "NOT_FOUND"
	}
}

func (s SiteAddress) getCountry() string {
	switch s {
	case Site_Google:
		return "US"
	case Site_Bing:
		return "US"
	case Site_Yahoo:
		return "England"
	default:
		return "NOT_FOUND"
	}
}

func main() {
	fmt.Println(Site_Google.getIP())
	fmt.Println(Site_Google.getCountry())
	fmt.Println(Site_Google)
	fmt.Println(Site_Google)
}

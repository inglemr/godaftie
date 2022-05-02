package daft

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
)

const (
	LISTINGS_PATH = "/listings"
)

type Ad struct {
	Listing Listing `json:"listing"`
	SavedAd bool    `json:"savedAd"`
}

type Listing struct {
	ID               int          `json:"id"`
	Title            string       `json:"title"`
	SeoTitle         string       `json:"seoTitle"`
	Sections         []string     `json:"sections"`
	SaleType         []string     `json:"saleType"`
	FeaturedLevel    string       `json:"featuredLevel"`
	PublishDate      int64        `json:"publishDate"`
	Price            string       `json:"price"`
	AbbreviatedPrice string       `json:"abbreviatedPrice"`
	NumBedrooms      string       `json:"numBedrooms"`
	NumBathrooms     string       `json:"numBathrooms"`
	PropertyType     string       `json:"propertyType"`
	DaftShortcode    string       `json:"daftShortcode"`
	Seller           Seller       `json:"seller"`
	Media            Media        `json:"media"`
	Ber              Ber          `json:"ber"`
	Point            Point        `json:"point"`
	SeoFriendlyPath  string       `json:"seoFriendlyPath"`
	PageBranding     PageBranding `json:"pageBranding"`
	Category         string       `json:"category"`
	State            string       `json:"state"`
	PremierPartner   bool         `json:"premierPartner"`
	Prs              Prs          `json:"prs"`
	Sticker          string       `json:"sticker"`
}

type Listings struct {
	Listings           []Ad               `json:"listings"`
	ShowcaseListings   []interface{}      `json:"showcaseListings"`
	Paging             Paging             `json:"paging"`
	DfpTargetingValues DfpTargetingValues `json:"dfpTargetingValues"`
	Breadcrumbs        []Breadcrumbs      `json:"breadcrumbs"`
	CanonicalURL       string             `json:"canonicalUrl"`
	MapView            bool               `json:"mapView"`
	SavedSearch        bool               `json:"savedSearch"`
}

type DfpTargetingValues struct {
	RentalPriceFrom []string `json:"rentalPriceFrom"`
	PageType        []string `json:"pageType"`
	CityName        []string `json:"cityName"`
	SearchPageNo    []string `json:"searchPageNo"`
	AreaName        []string `json:"areaName"`
	AdState         []string `json:"adState"`
	DistilledBrand  []string `json:"distilledBrand"`
	RentalPriceTo   []string `json:"rentalPriceTo"`
	Section         []string `json:"section"`
	CountyName      []string `json:"countyName"`
	IsUserLoggedIn  []string `json:"isUserLoggedIn"`
}
type Breadcrumbs struct {
	DisplayValue string `json:"displayValue"`
	URL          string `json:"url"`
}

type ListingsRequest struct {
	Section   string       `json:"section"`
	Filters   []Filters    `json:"filters"`
	Ranges    []Ranges     `json:"ranges"`
	Paging    PagingFilter `json:"paging"`
	GeoFilter GeoFilter    `json:"geoFilter"`
	Terms     string       `json:"terms"`
}

//There is no easy way to collect data from the daft.ie internal api and the
//webpage itself. This method attempts to collect some additional data about
//the listing via web scrapping
//Thus an additional data structure is used to return this data
type ListingData struct {
	Facilities  []string
	Features    []string
	Description string
	Location    string
}

func (cl *Client) GetListingData(path string) (ListingData, error) {
	///.//*[@data-testid='description']
	///.//*[@data-testid='location']
	listingData := ListingData{}
	doc, err := htmlquery.LoadURL(DAFT_URL + path)
	if err != nil {
		fmt.Printf("Error requesting %v got %v\n", path, err)
		log.Fatal(err)
		return listingData, err
	}
	facilities, err := htmlquery.QueryAll(doc, ".//*[@data-testid='facilities']/ul/li/text()")
	if err != nil {
		return listingData, err
	}

	for _, n := range facilities {
		listingData.Facilities = append(listingData.Facilities, n.Data)
	}

	location, err := htmlquery.QueryAll(doc, ".//div[@data-testid='location-text']/text()")
	if err != nil {
		return listingData, err
	}
	if len(location) > 0 {
		listingData.Location = location[0].Data
	}

	description, err := htmlquery.QueryAll(doc, ".//div[@data-testid='description']/text()")
	if err != nil {
		return listingData, err
	}
	if len(description) > 0 {
		listingData.Description = description[0].Data
	}

	features, err := htmlquery.QueryAll(doc, ".//*[@data-testid='features']/ul/li/text()")
	if err != nil {
		return listingData, err
	}
	for _, n := range features {
		listingData.Features = append(listingData.Features, n.Data)
	}
	return listingData, nil
}

func (cl *Client) GetListings(options ListingsRequest) (Listings, error) {
	var listings Listings
	b, er := json.Marshal(options)
	if er != nil {
		return listings, er
	}
	resp, err := cl.newRequest("POST", LISTINGS_PATH, b, nil)
	if err != nil {
		return listings, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return listings, err
	}
	json.Unmarshal(body, &listings)
	return listings, nil
}

func (listing *Listing) GetPriceAndCadence() (int, string) {
	price := 0
	cadence := "Price On Application"
	if listing.Price != "Price on Applicaton" {
		priceString := strings.ReplaceAll(listing.Price, "€", "")
		priceString = strings.ToLower(strings.ReplaceAll(priceString, ",", ""))
		price, _ = getIntFromString(priceString)
		if strings.Contains(priceString, "week") {
			cadence = "weekly"
		} else if strings.Contains(priceString, "month") {
			cadence = "monthly"
		} else {
			cadence = "onetime"
		}
	}
	return price, cadence
}

func (listing *Listing) GetBedCount() int {
	beds, _ := getIntFromString(listing.NumBedrooms)
	return beds
}

func (listing *Listing) GetBathCount() int {
	baths, _ := getIntFromString(listing.NumBathrooms)
	return baths
}

func getIntFromString(target string) (int, error) {
	match := regexp.MustCompile(`\d+`).FindString(target)
	return strconv.Atoi(match)
}

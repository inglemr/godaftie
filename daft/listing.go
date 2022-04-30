package daft

import (
	"encoding/json"
	"io/ioutil"
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

func (cl *Client) GetListings(options ListingsRequest) Listings {
	var listings Listings
	b, er := json.Marshal(options)
	if er != nil {
		panic(er)
	}
	resp, err := cl.newRequest("POST", LISTINGS_PATH, b, nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &listings)
	return listings
}

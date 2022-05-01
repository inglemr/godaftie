package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/inglemr/godaftie/daft"
)

func main() {
	dafClient := daft.New()
	dafClient.Init()
	listReq := daft.ListingsRequest{
		Paging: daft.PagingFilter{
			PageSize: 50,
		},
	}
	listings := dafClient.GetListings(listReq)
	// spew.Dump(listings.Paging)
	// spew.Dump(listings.Listings[0].Listing)
	// spew.Dump(listings.Listings[0].Listing.GetPriceAndCadence())
	// spew.Dump(listings.Listings[0].Listing.GetBathCount())
	// spew.Dump(listings.Listings[0].Listing.GetBedCount())
	fmt.Println(listings.Listings[0].Listing.SeoFriendlyPath)
	listingData := dafClient.GetListingData("/for-rent/griffith-wood-griffith-avenue-drumcondra-dublin-9/3523580")
	spew.Dump(listingData)
}

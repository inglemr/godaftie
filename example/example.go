package main

import (
	"github.com/inglemr/godaftie/daft"

	"github.com/davecgh/go-spew/spew"
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
	spew.Dump(listings.Paging)
	spew.Dump(listings.Listings[0].Listing)
	spew.Dump(listings.Listings[0].Listing.GetPriceAndCadence())
	spew.Dump(listings.Listings[0].Listing.GetBathCount())
	spew.Dump(listings.Listings[0].Listing.GetBedCount())
}

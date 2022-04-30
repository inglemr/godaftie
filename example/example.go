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
			PageSize: 100,
		},
	}
	listings := dafClient.GetListings(listReq)
	spew.Dump(listings.Paging)
	spew.Dump(listings.Listings[0])
}

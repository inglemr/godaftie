package daft

type Seller struct {
	SellerID         int    `json:"sellerId"`
	Name             string `json:"name"`
	Branch           string `json:"branch"`
	Address          string `json:"address"`
	Phone            string `json:"phone"`
	PhoneWhenToCall  string `json:"phoneWhenToCall"`
	BackgroundColour string `json:"backgroundColour"`
	SellerType       string `json:"sellerType"`
	ShowContactForm  bool   `json:"showContactForm"`
	SquareLogo       string `json:"squareLogo"`
	LicenceNumber    string `json:"licenceNumber"`
	StandardLogo     string `json:"standardLogo"`
}

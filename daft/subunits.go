package daft

type SubUnits struct {
	ID              int    `json:"id"`
	Price           string `json:"price"`
	NumBedrooms     string `json:"numBedrooms"`
	NumBathrooms    string `json:"numBathrooms"`
	PropertyType    string `json:"propertyType"`
	DaftShortcode   string `json:"daftShortcode"`
	SeoFriendlyPath string `json:"seoFriendlyPath"`
	Category        string `json:"category"`
	Image           Image  `json:"image"`
	Media           Media  `json:"media"`
	Ber             Ber    `json:"ber"`
}

type Prs struct {
	TotalUnitTypes   int        `json:"totalUnitTypes"`
	SubUnits         []SubUnits `json:"subUnits"`
	TagLine          string     `json:"tagLine"`
	Location         string     `json:"location"`
	AboutDevelopment string     `json:"aboutDevelopment"`
}

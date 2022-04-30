package daft

type Images struct {
	Size720X480 string `json:"size720x480"`
	Size600X600 string `json:"size600x600"`
	Size400X300 string `json:"size400x300"`
	Size360X240 string `json:"size360x240"`
	Size300X200 string `json:"size300x200"`
	Size320X280 string `json:"size320x280"`
	Size72X52   string `json:"size72x52"`
	Size680X392 string `json:"size680x392"`
	Caption     string `json:"caption,omitempty"`
}

type Image struct {
	Caption       string `json:"caption"`
	Size1440X960  string `json:"size1440x960"`
	Size1200X1200 string `json:"size1200x1200"`
	Size720X480   string `json:"size720x480"`
	Size600X600   string `json:"size600x600"`
	Size400X300   string `json:"size400x300"`
	Size360X240   string `json:"size360x240"`
	Size300X200   string `json:"size300x200"`
	Size320X280   string `json:"size320x280"`
	Size72X52     string `json:"size72x52"`
	Size680X392   string `json:"size680x392"`
}

type Media struct {
	Images         []Images `json:"images"`
	TotalImages    int      `json:"totalImages"`
	HasVideo       bool     `json:"hasVideo"`
	HasVirtualTour bool     `json:"hasVirtualTour"`
	HasBrochure    bool     `json:"hasBrochure"`
}

type PageBranding struct {
	StandardLogo     string   `json:"standardLogo"`
	SquareLogo       string   `json:"squareLogo"`
	BackgroundColour string   `json:"backgroundColour"`
	SquareLogos      []string `json:"squareLogos"`
	RectangleLogo    string   `json:"rectangleLogo"`
}

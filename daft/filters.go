package daft

type Filters struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
type Ranges struct {
	From string `json:"from"`
	To   string `json:"to"`
	Name string `json:"name"`
}
type PagingFilter struct {
	From     string `json:"from"`
	PageSize int    `json:"pageSize"`
}
type GeoFilter struct {
	StoredShapeIds []string `json:"storedShapeIds"`
	GeoSearchType  string   `json:"geoSearchType"`
}

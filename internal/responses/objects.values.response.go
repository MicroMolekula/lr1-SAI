package responses

type ObjectValuesResponse struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

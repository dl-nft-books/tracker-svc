package opensea

type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	BannerURL   string `json:"external_url"`
}

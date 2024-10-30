package dto

type CreateShortUrlPayload struct {
	OriginalUrl    string `json:"original_url"`

}

type UrlPayload struct {
	ShortCode    string `json:"short_code"`
}

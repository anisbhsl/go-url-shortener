package model

//URL defines
type URL struct {
	ID       string `json:"id,omitempty"`
	LongURL  string `json:"longURL,omitempty"`
	ShortURL string `json:"shortURL,omitempty"`
}

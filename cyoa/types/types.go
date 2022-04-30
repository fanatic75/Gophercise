package types

type Book map[string]Arc
type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Arc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

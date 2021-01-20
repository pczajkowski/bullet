package bullet

type Push struct {
	ID       string  `json:"iden"`
	Active   bool    `json:"active"`
	Created  float64 `json:"created"`
	Type     string  `json:"type"`
	Title    string  `json:"title"`
	Body     string  `json:"body"`
	URL      string  `json:"url"`
	FileName string  `json:"file_name"`
	FileType string  `json:"file_type"`
	FileURL  string  `json:"file_url"`
	ImageURL string  `json:"image_url"`
}

type Pushes struct {
	Items  []Push `json:"pushes"`
	Cursor string `json:"cursor"`
}

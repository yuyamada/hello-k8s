package article

// Item is an article
type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Articles is a slice of articles
type Articles struct {
	Items []Item
}

// New creates new insrance of Articles
func New() *Articles {
	return &Articles{}
}

// Add appends new item to articles
func (r *Articles) Add(item Item) {
	r.Items = append(r.Items, item)
}

// GetAll returns a slice of articles
func (r *Articles) GetAll() []Item {
	return r.Items
}

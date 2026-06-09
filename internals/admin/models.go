package admin

type CreateBlogPostRequestType struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
	Cover string   `json:"cover"`
}

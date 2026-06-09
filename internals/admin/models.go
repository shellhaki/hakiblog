package admin

type CreateBlogPostRequestType struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
	Cover string   `json:"cover"`
}

type BlogPost struct {
	ID    int64    `json:"id"`
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
	Cover string   `json:"cover"`
	
}

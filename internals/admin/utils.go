package admin

import (
	"errors"
	"haki/blog/internals/config"
)

func CreateBlogPost(title string, body string, tags []string, cover string) error {
	db := config.DB
	defer db.Close()
	if title == "" && body == "" && len(tags) == 0 && cover == "" {
		return errors.New("title,body,tags,cover required!")
	}
	query := "insert into blog_post(title,body,tags,cover) values($1,$2,$3,$4)"
	_, err := db.Exec(query, title, body, tags, cover)
	if err != nil {
		return err
	}
	return nil
}

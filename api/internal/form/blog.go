package form

import "gopkg.in/guregu/null.v3"

type Blog struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateBlog struct {
	Title       null.String `json:"title"`
	Description null.String `json:"description"`
}

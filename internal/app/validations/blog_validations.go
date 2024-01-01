// blog_validations.go
package validations

type CreateBlogInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

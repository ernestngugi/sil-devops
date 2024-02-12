package controller

import (
	"context"

	"github.com/ernestngugi/sil-devops/internal/db"
	"github.com/ernestngugi/sil-devops/internal/form"
	"github.com/ernestngugi/sil-devops/internal/model"
	"github.com/ernestngugi/sil-devops/internal/repos"
)

type (
	BlogController interface {
		CreateBlog(ctx context.Context, dB db.DB, blog *form.Blog) (*model.Blog, error)
		BlogByID(ctx context.Context, dB db.DB, blogID int64) (*model.Blog, error)
		AllBlogs(ctx context.Context, dB db.DB) ([]*model.Blog, error)
		DeleteBlog(ctx context.Context, dB db.DB, blogID int64) error
		UpdateBlog(ctx context.Context, dB db.DB, blogID int64, form *form.UpdateBlog) (*model.Blog, error)
	}

	AppBlogController struct {
		blogRepository repos.BlogRepository
	}
)

func NewBlogRepository(
	blogRepository repos.BlogRepository,
) BlogController {
	return &AppBlogController{
		blogRepository: blogRepository,
	}
}

func (c *AppBlogController) DeleteBlog(
	ctx context.Context,
	dB db.DB,
	blogID int64,
) error {

	blog, err := c.blogRepository.BlogByID(ctx, dB, blogID)
	if err != nil {
		return err
	}

	return c.blogRepository.DeleteBlog(ctx, dB, blog.ID)
}

func (c *AppBlogController) AllBlogs(
	ctx context.Context,
	dB db.DB,
) ([]*model.Blog, error) {
	return c.blogRepository.ListBlogs(ctx, dB)
}

func (c *AppBlogController) UpdateBlog(
	ctx context.Context,
	dB db.DB,
	blogID int64,
	form *form.UpdateBlog,
) (*model.Blog, error) {

	blog, err := c.blogRepository.BlogByID(ctx, dB, blogID)
	if err != nil {
		return &model.Blog{}, err
	}

	if form.Title.Valid {
		blog.Title = form.Title.String
	}

	if form.Description.Valid {
		blog.Description = form.Description.String
	}

	err = c.blogRepository.Save(ctx, dB, blog)
	if err != nil {
		return &model.Blog{}, err
	}

	return blog, nil
}

func (c *AppBlogController) BlogByID(
	ctx context.Context,
	dB db.DB,
	blogID int64,
) (*model.Blog, error) {
	return c.blogRepository.BlogByID(ctx, dB, blogID)
}

func (c *AppBlogController) CreateBlog(
	ctx context.Context,
	dB db.DB,
	form *form.Blog,
) (*model.Blog, error) {

	blog := &model.Blog{
		Title:       form.Title,
		Description: form.Description,
	}

	err := c.blogRepository.Save(ctx, dB, blog)
	if err != nil {
		return &model.Blog{}, err
	}

	return blog, nil
}

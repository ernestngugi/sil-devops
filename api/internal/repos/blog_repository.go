package repos

import (
	"context"
	"time"

	"github.com/ernestngugi/sil-devops/api/internal/db"
	"github.com/ernestngugi/sil-devops/api/internal/model"
)

const (
	createBlog     = "INSERT INTO blogs (title, description, date_created, date_modified) VALUES ($1, $2, $3, $4)"
	updateBlog     = "UPDATE blogs SET title = ?, description = ?, date_modified = ? WHERE id = ?"
	selectBlog     = "SELECT id, title, description, date_created, date_modified FROM blogs"
	selectBlogByID = selectBlog + " WHERE id = ?"
	deleteBlog     = "DELETE blogs WHERE id = ?"
)

type (
	BlogRepository interface {
		ListBlogs(ctx context.Context, dB db.DB) ([]*model.Blog, error)
		Save(ctx context.Context, dB db.DB, blog *model.Blog) error
		BlogByID(ctx context.Context, dB db.DB, blogID int64) (*model.Blog, error)
		DeleteBlog(ctx context.Context, dB db.DB, blogID int64) error
	}

	AppBlogRepository struct{}
)

func NewBlogRepository() BlogRepository {
	return &AppBlogRepository{}
}

func (r *AppBlogRepository) DeleteBlog(
	ctx context.Context,
	dB db.DB,
	blogID int64,
) error {

	_, err := dB.ExecContext(
		ctx,
		deleteBlog,
		blogID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *AppBlogRepository) ListBlogs(
	ctx context.Context,
	dB db.DB,
) ([]*model.Blog, error) {

	rows, err := dB.QueryContext(
		ctx,
		selectBlog,
	)
	if err != nil {
		return []*model.Blog{}, err
	}

	defer rows.Close()

	blogs := make([]*model.Blog, 0)

	for rows.Next() {

		var blog model.Blog

		err := rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Description,
			&blog.DateCreated,
			&blog.DateModified,
		)
		if err != nil {
			return []*model.Blog{}, err
		}

		blogs = append(blogs, &blog)
	}

	return blogs, nil
}

func (r *AppBlogRepository) BlogByID(
	ctx context.Context,
	dB db.DB,
	blogID int64,
) (*model.Blog, error) {

	var blog model.Blog

	err := dB.QueryRowContext(
		ctx,
		selectBlogByID,
		blogID,
	).Scan(
		&blog.ID,
		&blog.Title,
		&blog.Description,
		&blog.DateCreated,
		&blog.DateModified,
	)
	if err != nil {
		return &model.Blog{}, err
	}

	return &blog, nil
}

func (r *AppBlogRepository) Save(
	ctx context.Context,
	dB db.DB,
	blog *model.Blog,
) error {

	blog.DateCreated = time.Now()
	blog.DateModified = time.Now()

	if blog.ID == 0 {

		_, err := dB.ExecContext(
			ctx,
			createBlog,
			blog.Title,
			blog.Description,
			blog.DateCreated,
			blog.DateModified,
		)
		if err != nil {
			return err
		}
	}

	_, err := dB.ExecContext(
		ctx,
		updateBlog,
		blog.Title,
		blog.Description,
		blog.DateModified,
		blog.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

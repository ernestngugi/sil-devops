package repos

import (
	"context"
	"time"

	"github.com/ernestngugi/sil-devops/internal/db"
	"github.com/ernestngugi/sil-devops/internal/model"
)

const (
	createBlogSQL     = "INSERT INTO blogs (title, description, date_created, date_modified) VALUES ($1, $2, $3, $4) RETURNING id"
	updateBlogSQL     = "UPDATE blogs SET title = $1, description = $2, date_modified = $3 WHERE id = $4"
	selectBlogSQL     = "SELECT id, title, description, date_created, date_modified FROM blogs"
	selectBlogByIDSQL = selectBlogSQL + " WHERE id = $1"
	deleteBlogSQL     = "DELETE FROM blogs WHERE id = $1"
)

type (
	BlogRepository interface {
		ListBlogs(ctx context.Context, operations db.SQLOperations) ([]*model.Blog, error)
		Save(ctx context.Context, operations db.SQLOperations, blog *model.Blog) error
		BlogByID(ctx context.Context, operations db.SQLOperations, blogID int64) (*model.Blog, error)
		DeleteBlog(ctx context.Context, operations db.SQLOperations, blogID int64) error
	}

	AppBlogRepository struct{}
)

func NewBlogRepository() BlogRepository {
	return &AppBlogRepository{}
}

func (r *AppBlogRepository) DeleteBlog(
	ctx context.Context,
	operations db.SQLOperations,
	blogID int64,
) error {

	_, err := operations.ExecContext(
		ctx,
		deleteBlogSQL,
		blogID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *AppBlogRepository) ListBlogs(
	ctx context.Context,
	operations db.SQLOperations,
) ([]*model.Blog, error) {

	rows, err := operations.QueryContext(
		ctx,
		selectBlogSQL,
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
	operations db.SQLOperations,
	blogID int64,
) (*model.Blog, error) {

	var blog model.Blog

	err := operations.QueryRowContext(
		ctx,
		selectBlogByIDSQL,
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
	operations db.SQLOperations,
	blog *model.Blog,
) error {

	blog.DateCreated = time.Now()
	blog.DateModified = time.Now()

	if blog.ID == 0 {

		err := operations.QueryRowContext(
			ctx,
			createBlogSQL,
			blog.Title,
			blog.Description,
			blog.DateCreated,
			blog.DateModified,
		).Scan(&blog.ID)
		if err != nil {
			return err
		}

		return nil
	}

	_, err := operations.ExecContext(
		ctx,
		updateBlogSQL,
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

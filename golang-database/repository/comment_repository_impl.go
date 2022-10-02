package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepositoryImpl(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (reposotory *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	//TODO insert to table comments
	script := "INSERT INTO comments (email, comment) value (?, ?)"
	result, err := reposotory.DB.ExecContext(ctx, script, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil

}

func (reposotory *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	//TODO Select from table
	script := "SELECT id, email, comment FROM comments where id = ? LIMIT 1"
	rows, err := reposotory.DB.QueryContext(ctx, script, id)
	defer rows.Close()

	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("id not found")
	}
}

func (reposotory *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	//TODO Select all from table
	script := "SELECT id, email, comment FROM comments"
	rows, err := reposotory.DB.QueryContext(ctx, script)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		comments = append(comments, comment)
	}
	return comments, nil
}

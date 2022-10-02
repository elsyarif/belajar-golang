package repository

import (
	"context"
	"golang-database/entity"
)

//CommentRepository TODO: Comment Rpository
//function atau metohod yang berhubungan langsung dengan database
type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}

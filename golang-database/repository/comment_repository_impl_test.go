package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "haidar@gmail.com",
		Comment: "komen haidar",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(golang_database.GetConnection())
	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 35)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(golang_database.GetConnection())
	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}

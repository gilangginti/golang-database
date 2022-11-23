package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	golangdatabase "golang-database"
	"golang-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "testReposito@gmail.com",
		Comment: "Test Rep",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println("Berhasil Insert Comment", result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	result, err := commentRepository.FindById(ctx, 33)
	if err != nil {
		panic(err)
	}
	fmt.Println("Berhasil mengambil data by id", result)
}
func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Berhasil mengambil data", result)
}

package page_test

import (
	"context"
	"database/sql"
	"testing"

	repo "github.com/Hives-project/HivePages/pkg/storage/mysql/page"

	"github.com/Hives-project/HivePages/pkg/page"
)

func TestHelloEmpty(t *testing.T) {
	repo := repo.NewPageRepository(&sql.DB{})
	service := page.NewPageService(repo)
	resp, err := service.GetPages(context.Background(), "1")
	if resp != nil && err != nil {
		t.Fatalf(`did not hello world`)
	}
}

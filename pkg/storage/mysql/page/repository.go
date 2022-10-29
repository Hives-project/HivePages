package page

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Hives-project/HivePages/pkg/page"
	"github.com/go-sql-driver/mysql"
)

type pageRepository struct {
	db *sql.DB
}

func NewPageRepository(sql *sql.DB) page.PageRepository {
	return &pageRepository{
		db: sql,
	}
}

func (r *pageRepository) GetPages(ctx context.Context, pageId string) ([]page.GetPage, error) {
	var pages []page.GetPage
	result, err := r.db.Query("SELECT `firstname`, `lastname` from `pages`")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var page page.GetPage
		err := result.Scan(&page.Firstname, &page.Lastname)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}
	return pages, nil
}

func (r *pageRepository) CreatePage(ctx context.Context, page page.CreatePage) error {
	stmt, err := r.db.Prepare("INSERT INTO pages(id, firstname, lastname) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(page.Uuid, page.Firstname, page.Lastname)
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == 1062 {
			return errors.New("this page already exists")
		}
	} else if err != nil {
		return err
	}
	return nil
}

package page

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Hives-project/HivePages/pkg/page"
	"github.com/Hives-project/HivePages/pkg/util"
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

func (r *pageRepository) GetPages(ctx context.Context) ([]page.GetPage, error) {
	var pages []page.GetPage
	result, err := r.db.Query("SELECT `uuid`, `pageName`, `description` from `pages`")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var page page.GetPage
		err := result.Scan(&page.Uuid, &page.PageName, &page.Description)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}
	return pages, nil
}

func (r *pageRepository) CreatePage(ctx context.Context, page page.CreatePage) error {
	stmt, err := r.db.Prepare("INSERT INTO pages(uuid, pageName, description) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(page.Uuid, page.PageName, page.Description)
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == 1062 {
			return errors.New("this page already exists")
		}
	} else if err != nil {
		return err
	}
	return nil
}

func (r *pageRepository) GetPageById(ctx context.Context, pageId string) (page.GetPage, error) {
	var page page.GetPage
	row := r.db.QueryRow("SELECT `uuid`, `pageName`, `description` FROM `pages` WHERE uuid = ?", pageId)
	err := row.Scan(&page.Uuid, &page.PageName, &page.Description)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return page, util.NewErrorf(err, util.ErrorCodeNotFound, "page with pageid: %s does not exist", pageId)
	case err != nil:
		return page, util.NewErrorf(err, util.ErrorCodeInternal, "internal server error")
	default:
		return page, nil
	}
}

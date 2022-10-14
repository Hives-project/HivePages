package page

import (
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

func (r *pageRepository) GetPages() ([]page.GetPage, error) {
	var pages []page.GetPage
	result, err := r.db.Query("SELECT `id`, `firstname`, `lastname` from `pages`")
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

func (r *pageRepository) GetPageByUuid(uuid string) (page.GetPage, error) {
	var brand page.GetPage

	row := r.db.QueryRow("SELECT `id`, `firstname`, `lastname` from `pages` WHERE id = ?", uuid)
	err := row.Scan(&brand.Uuid, &brand.Firstname, &brand.Lastname)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return brand, util.NewErrorf(err, util.ErrorCodeNotFound, "brand with id: %s does not exist", uuid)
	case err != nil:
		return brand, util.NewErrorf(err, util.ErrorCodeInternal, "internal server error")
	default:
		return brand, nil
	}
}

func (r *pageRepository) CreatePage(page page.GetPage) error {
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

func (r *pageRepository) DeletePage(uuid string) error {
	stmt, err := r.db.Prepare("DELETE FROM pages WHERE uuid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(uuid)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("page does not exist")
	}
	return nil
}

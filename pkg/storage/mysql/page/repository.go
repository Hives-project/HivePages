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

// // Gets all pages from database and returns array of pages
func (r *pageRepository) GetPages(ctx context.Context, pageId string) ([]page.GetPage, error) {
	var pages []page.GetPage
	result, err := r.db.Query("SELECT `firstname`, `lastname` from `pages`;")
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

// func (r *pageRepository) GetPageByUuid(uuid string) error {
// 	var page models.Page

// 	row := r.db.QueryRow("SELECT {ELEMENTS} from {TABLENAME} WHERE uuid = ?;", uuid)
// 	err := row.Scan(&page.ID, &page.UUID, &page.Email, &page.Subscribed)

// 	switch {
// 	case errors.Is(err, sql.ErrNoRows):
// 		return errors.New("page does not exist")
// 	case err != nil:
// 		return errors.New("internal server error")
// 	default:
// 		return nil
// 	}
// }

// Creates a new record in database
// If already exists, returns error that page exists
func (r *pageRepository) CreatePage(ctx context.Context, page page.CreatePage) error {
	stmt, err := r.db.Prepare("INSERT INTO pages(id, firstname, lastname) VALUES(?, ?, ?);")
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

// Deletes a record from database
// If record doesn't exist, returns not exist error
// func (r *pageRepository) DeletePage(uuid string) error {
// 	stmt, err := r.db.Prepare("DELETE FROM pages WHERE uuid = ?;")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	result, err := stmt.Exec(uuid)
// 	if err != nil {
// 		return err
// 	}

// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}

// 	if rows != 1 {
// 		return errors.New("Page does not exist")
// 	}
// 	return nil
// }

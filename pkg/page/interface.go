package page

import "context"

//go:generate mockgen --source=logic.go --destination=./mock/mock_logic.go
type PageService interface {
	CreatePage(ctx context.Context, page CreatePage) error
	GetPages(ctx context.Context, uuid string) ([]GetPage, error)
}

//go:generate mockgen --source=../storage/mysql/page/repository.go --destination=./mock/mock_repository.go
type PageRepository interface {
	CreatePage(ctx context.Context, page CreatePage) error
	GetPages(ctx context.Context, pageId string) ([]GetPage, error)
}

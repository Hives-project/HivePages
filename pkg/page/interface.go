package page

import "context"

//go:generate mockgen --source=logic.go --destination=./mock/mock_logic.go
type PageService interface {
	CreatePage(ctx context.Context, page Page) error
	GetPages(ctx context.Context) ([]Page, error)
	GetPageById(ctx context.Context, pageId string) (Page, error)
}

//go:generate mockgen --source=../storage/mysql/page/repository.go --destination=./mock/mock_repository.go
type PageRepository interface {
	CreatePage(ctx context.Context, page Page) error
	GetPages(ctx context.Context) ([]Page, error)
	GetPageById(ctx context.Context, pageId string) (Page, error)
}

package page

//go:generate mockgen --source=interface.go --destination=./mock/mock_logic.go
type PageService interface {
	CreatePage(page GetPage) error
	GetPages() ([]GetPage, error)
	GetPageByUuid(uuid string) (GetPage, error)
}

//go:generate mockgen --source=interface.go --destination=./mock/mock_repository.go
type PageRepository interface {
	CreatePage(page GetPage) error
	GetPages() ([]GetPage, error)
	GetPageByUuid(uuid string) (GetPage, error)
}

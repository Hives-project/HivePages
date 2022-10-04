package page

import (
	"context"

	"github.com/Hives-project/HivePages/pkg/util"

	"github.com/google/uuid"
)

type pageService struct {
	pageRepository PageRepository
}

func NewPageService(u PageRepository) PageService {
	return &pageService{
		pageRepository: u,
	}
}

func (u *pageService) CreatePage(ctx context.Context, page CreatePage) error {
	page.Uuid = uuid.New().String()
	if err := u.pageRepository.CreatePage(ctx, page); err != nil {
		return util.NewErrorf(err, util.ErrorCodeInternal, "%s", "could not create page")
	}
	return nil
}

func (u *pageService) GetPages(ctx context.Context, uuid string) ([]GetPage, error) {
	pages, err := u.pageRepository.GetPages(ctx, uuid)
	if err != nil {
		return nil, util.NewErrorf(err, util.ErrorCodeInternal, "%s", "could not get pages")
	}
	return pages, nil
}

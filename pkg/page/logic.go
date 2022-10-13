package page

import (
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

func (u *pageService) CreatePage(page GetPage) error {
	page.Uuid = uuid.New().String()
	if err := u.pageRepository.CreatePage(page); err != nil {
		return util.NewErrorf(err, util.ErrorCodeInternal, "%s", "could not create page")
	}
	return nil
}

func (u *pageService) GetPages() ([]GetPage, error) {
	pages, err := u.pageRepository.GetPages()
	if err != nil {
		return nil, util.NewErrorf(err, util.ErrorCodeInternal, "%s", "could not get pages")
	}
	return pages, nil
}

func (u *pageService) GetPageByUuid(uuid string) (GetPage, error) {
	page, err := u.pageRepository.GetPageByUuid(uuid)
	if err != nil {
		return GetPage{}, util.NewErrorf(err, util.ErrorCodeInternal, "%s", "could not get page")
	}
	return page, util.NewErrorf(err, util.ErrorCodeInternal, "could not get page")
}

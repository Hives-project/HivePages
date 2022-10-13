package page_test

import (
	"testing"

	"github.com/Hives-project/HivePages/pkg/page"
	repo "github.com/Hives-project/HivePages/pkg/page/mock"
	"github.com/golang/mock/gomock"
)

func TestHelloEmpty(t *testing.T) {
	controller := gomock.NewController(t)
	mockRepo := repo.NewMockPageRepository(controller)
	service := page.NewPageService(mockRepo)

	mockRepo.EXPECT().GetPages()

	resp, err := service.GetPages()
	if resp != nil && err != nil {
		t.Fatalf(`did not succeed`)
	}
}

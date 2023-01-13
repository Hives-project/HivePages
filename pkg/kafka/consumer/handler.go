package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Hives-project/HivePages/pkg/kafka/producer"
	"github.com/Hives-project/HivePages/pkg/page"
	"github.com/segmentio/kafka-go"
)

var (
	pageService page.PageService
	ctx         context.Context = context.Background()
)

func RouteMessagesFromTopics(service page.PageService, m kafka.Message) {
	pageService = service

	switch m.Topic {
	case "sndosdzx-createPage":
		createPage(m.Value)
	case "sndosdzx-getUsername":
		getUsername(m.Value)
	}
}

func createPage(message []byte) {
	var page page.Page
	if err := json.Unmarshal(message, &page); err != nil {
		log.Printf("Invalid request payload- createPage()")
		return
	}

	if err := pageService.CreatePage(ctx, page); err != nil {
		log.Printf("internal server error: %s", err.Error())
		return
	}
}

func getUsername(message []byte) {
	var req pageWithKrabbelid
	if err := json.Unmarshal(message, &req); err != nil {
		log.Printf("Invalid request payload - getUsername()")
		return
	}

	page, err := pageService.GetPageById(ctx, req.PageId)
	if err != nil {
		log.Printf("internal server error: %s", err.Error())
		return
	}

	if page.UserName == "" || req.KrabbelId == "" {
		log.Printf("username or krabbelid empty: %s, %s", page.UserName, req.KrabbelId)
		return
	}

	if err = producer.UpdateUsername(page.UserName, req.KrabbelId); err != nil {
		log.Printf("error producing message: %s", err.Error())
	}
}

type pageWithKrabbelid struct {
	PageId    string `json:"page_id"`
	KrabbelId string `json:"krabbel_id"`
}

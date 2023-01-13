package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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
	case "createPage":
		createPage(m.Value)
	case "getUsername":
		getUsername(m.Value)
	}
}

func createPage(value []byte) {
	var page page.Page
	if err := json.Unmarshal(value, &page); err != nil {
		log.Printf("Invalid request payload")
		return
	}
	if err := pageService.CreatePage(ctx, page); err != nil {
		log.Printf("internal server error: %s", err.Error())
		return
	}
}

func getUsername(value []byte) {
	var req page.Page
	if err := json.Unmarshal(value, &req); err != nil {
		log.Printf("Invalid request payload")
		return
	}
	page, err := pageService.GetPageById(ctx, req.Uuid)
	if err != nil {
		log.Printf("internal server error: %s", err.Error())
		return
	}
	fmt.Println(page.UserName)
	// Todo: push username to topic for krabbels to pick up
}

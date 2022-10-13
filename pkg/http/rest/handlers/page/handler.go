package page

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Hives-project/HivePages/pkg/http/rest/handlers"
	"github.com/Hives-project/HivePages/pkg/page"
	"github.com/gorilla/mux"
)

func CreatePageHandler(pageService page.PageService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var body page.GetPage

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			handlers.RenderErrorResponse(w, "Invalid request payload", r.URL.Path, err)
			return
		}

		pageService.CreatePage(body)
	}
}

func GetPageHandler(pageService page.PageService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := mux.Vars(r)["uuid"]
		if len(uuid) == 0 {
			err := errors.New("query parameters are invalid")
			handlers.RenderErrorResponse(w, err.Error(), r.URL.Path, err)
			return
		}

		page, err := pageService.GetPages()
		if err != nil {
			handlers.RenderErrorResponse(w, "internal server error", r.URL.Path, err)
		}

		handlers.RenderResponse(w, http.StatusOK, page)
	}
}

func GetPageByUuidHandler(pageService page.PageService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := mux.Vars(r)["uuid"]
		if len(uuid) == 0 {
			handlers.RenderErrorResponse(w, "query parameters are invalid", r.URL.Path, errors.New("query parameters are invalid"))
			return
		}

		page, err := pageService.GetPageByUuid(uuid)
		if err != nil {
			handlers.RenderErrorResponse(w, "internal server error", r.URL.Path, err)
			return
		}

		handlers.RenderResponse(w, http.StatusOK, page)
	}
}

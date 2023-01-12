package handlers

import (
	"net/http"

	"github.com/Hives-project/HivePages/pkg/page"
	"github.com/gorilla/mux"
)

func GetPageById(pageSvc page.PageService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := mux.Vars(r)["uuid"]
		page, err := pageSvc.GetPageById(r.Context(), uuid)
		if err != nil {
			RenderErrorResponse(w, "internal server error", r.URL.Path, err)
			return
		}

		RenderResponse(w, http.StatusOK, page)
	}
}

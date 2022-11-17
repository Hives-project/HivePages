package page

type GetPage struct {
	Uuid        string `json:"uuid"`
	PageName    string `json:"pagename"`
	Description string `json:"description"`
}

type CreatePage struct {
	Uuid        string `json:"uuid,omitempty"`
	PageName    string `json:"pagename"`
	Description string `json:"description"`
}

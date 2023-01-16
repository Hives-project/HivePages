package page

type Page struct {
	Uuid        string `json:"uuid,omitempty"`
	PageName    string `json:"pagename"`
	Description string `json:"description"`
	UserName    string `json:"username"`
}

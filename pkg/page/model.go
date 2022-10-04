package page

type GetPage struct {
	Firstname string `json:"name"`
	Lastname  string `json:"lastname"`
}

type CreatePage struct {
	Uuid      string `json:"_id"`
	Firstname string `json:"name"`
	Lastname  string `json:"lastname"`
}

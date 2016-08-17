package models

type Person struct {
	PrimaryEmail string   `json:"primary_email"`
	Name         string   `json:"name"`
	Mobile       string   `json:"mobile"`
	Class        []string `json:"class"`
	ProgLanguage []string `json:"prog_language"`
}

type Class struct {
	Classid      string   `json:"classid"`
	Name         string   `json:"name"`
	Coach        string   `json:"coach"`
	Descript     string   `json:"descript"`
	Date         []string `json:"date"`
	Time         []string `json:"time"`
	Prerequisite string   `json:"prerequisite"`
	Memo         string   `json:"memo"`
}

type RespMsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Email struct {
	Email string `json:"email"`
}

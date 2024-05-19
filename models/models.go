package models

type Subscriber struct {
	Email string `db:"email" json:"email"`
}

type EmailDataObject struct {
	Name    string
	Subject string
	Message string
}

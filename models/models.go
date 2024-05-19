package models

type Subscriber struct {
	Email string `db:"email" json:"email"`
}

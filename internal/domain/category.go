package domain

type Category struct {
	ID   string `db:"id" json:"id" validate:"uuid"`
	Name string `db:"name" json:"name"`
}

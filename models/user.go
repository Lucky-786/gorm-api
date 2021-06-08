package models

type User struct {
	Id        int     `gorm:"primary_key" json:"id"`
	CreatedAt int64   `json:"-"`
	UpdatedAt int64   `json:"-"`
	FirstName string  `json:"first_name" validate:"alpha"`
	LastName  string  `json:"last_name" validate:"alpha"`
	City      string  `json:"city"  validate:"alpha"`
	Phone     int     `json:"phone"`
	Height    float32 `json:"height"`
	Gender    string  `json:"gender"  validate:"alpha"`
	Password  string  `json:"password"`
	Married   string  `json:"married"  validate:"alpha"`
}

type Result struct {
	Ids []int `json:"ids"`
}

func (User) TableName() string {
	return "user"
}

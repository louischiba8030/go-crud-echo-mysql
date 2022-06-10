package dto

type MemberResponse struct {
	ID		uint		`json:"id" gorm:"auto_increment;primary_key"`
	Name	string	`json:"name" gorm:"type:varchar(255);not null"`
	Age		uint		`json:"age" gorm:"not null;default:0"`
	Bloodtype	string	`json:"bloodtype" gorm:"type:varchar(255);not null"`
	Origin	string	`json:"origin" gorm:"type:varchar(255);not null"`
}

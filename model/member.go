package model

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
	"fmt"
)

type UUIDBaseModel struct {
	ID uuid.UUID `gorm:"primaryKey;unique;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `"gorm:index"`
}

// This functions are called before creating Member
func (uuidbasemodel *UUIDBaseModel) BeforeCreate (tx *gorm.DB) error {
	if uuid, err := uuid.NewRandom(); err != nil {
		fmt.Println(err)
		return err
	} else {
		tx.Statement.SetColumn("ID", uuid.String())
	}

	return nil
}

type Member struct {
	UUIDBaseModel
	//	ID		uint		`json:"id" gorm:"auto_increment;primary_key"`
	Name	string	`json:"name" gorm:"type:varchar(255);not null"`
	Age		uint		`json:"age" gorm:"not null;default:0"`
	Bloodtype	string	`json:"bloodtype" gorm:"type:varchar(255);not null"`
	Origin	string	`json:"origin" gorm:"type:varchar(255);not null"`
}

func (p *Member) FirstById(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).First(&p)
}

func (p *Member) Create() (tx *gorm.DB) {
	return DB.Create(&p)
}

// all column update
func (p *Member) Save() (tx *gorm.DB) {
	return DB.Save(&p)
}

func (p *Member) Updates() (tx *gorm.DB) {
	return DB.Model(&p).Updates(p)
}

func (p *Member) Delete() (tx *gorm.DB) {
	return DB.Delete(&p)
}

func (p *Member) DeleteById (id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).Delete(&p)
}

func (p *Member) DropTable() (tx *gorm.DB) {
	DB.Migrator().DropTable(&p)

	return nil
}

func (p *Member) CreateTable() (tx *gorm.DB) {
	DB.Migrator().CreateTable(&p)

	return nil
}

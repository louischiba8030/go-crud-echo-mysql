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

func (p *Member) FirstById(id uuid.UUID) (tx *gorm.DB) {
	return DB.Where("ID = ?", id).First(&p)
}

func (p *Member) Create() (tx *gorm.DB) {
	return DB.Create(&p)
}

// all column update
func (p *Member) Save() (tx *gorm.DB) {
	return DB.Save(&p)
}

func (p *Member) Updates(id uuid.UUID) (tx *gorm.DB) {
	// fmt.Printf("id %v\n", id) // not p.ID ?
	// fmt.Printf("p.Name %v\n", p.Name)
	return DB.Model(&p).Where("ID = ?", id).Updates(
		Member{
			Name: p.Name,
			Age: p.Age,
			Bloodtype:p.Bloodtype,
			Origin: p.Origin,
		})
}

func (p *Member) Delete() (tx *gorm.DB) {
	return DB.Delete(&p)
}

func (p *Member) DeleteById (id uuid.UUID) (tx *gorm.DB) {
	return DB.Where("ID = ?", id).Delete(&p)
}

func (p *Member) DropTable() (tx *gorm.DB) {
	DB.Migrator().DropTable(&p)

	return nil
}

func (p *Member) CreateTable() (tx *gorm.DB) {
	DB.Migrator().CreateTable(&p)

	return nil
}

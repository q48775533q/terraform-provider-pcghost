package pet

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Create creates a pet in the database
func Create(db *gorm.DB, pet *Pet) (string, error) {
	err := db.Create(pet).Error
	if err != nil {
		return "", err
	}
	return pet.ID, nil
}

//FindById returns a pet with a given id, or nil if not found
func FindById(db *gorm.DB, id string) (*Pet, error) {
	var pet Pet
	err := db.Find(&pet, &Pet{ID: id}).Error
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

//FindByName returns a pet with a given name, or nil if not found
func FindByName(db *gorm.DB, name string) (*Pet, error) {
	var pet Pet
	err := db.Find(&pet, &Pet{Name: name}).Error
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

//List returns all Pets in database, with a given limit
func List(db *gorm.DB, limit uint) (*[]Pet, error) {
	var pets []Pet
	err := db.Find(&pets).Limit(limit).Error
	if err != nil {
		return nil, err
	}
	return &pets, nil
}

//Update updates a pet in the database
func Update(db *gorm.DB, pet *Pet) error {
	err := db.Save(pet).Error
	return err
}

//Delete deletes a pet in the database
func Delete(db *gorm.DB, id string) error {
	pet, err := FindById(db, id)
	if err != nil {
		fmt.Printf("1:%v", err)
		return err
	}
	err = db.Delete(pet).Error
	fmt.Printf("2:%v", err)
	return err
}
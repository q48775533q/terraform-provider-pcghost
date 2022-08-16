package pets

import (
	"github.com/jinzhu/gorm"
	"github.com/TyunTech/terraform-petstore/model/pet"
)

//GetPetRequest request struct
type GetPetRequest struct {
	ID string
}

//GetPet returns a pet from database
func GetPet(db *gorm.DB, req *GetPetRequest) (*pet.Pet, error) {
	p, err := pet.FindById(db, req.ID)
	res := p
	return res, err
}

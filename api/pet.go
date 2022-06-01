package api

import (
	"encoding/json"
	"errors"
	mydb "github.com/StepanShevelev/fixedtask/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func apiCreatePet(w http.ResponseWriter, r *http.Request) {
	if !isMethodPOST(w, r) {
		return
	}
	mydb.CreatePet(r)
}

func apiUpdatePet(w http.ResponseWriter, r *http.Request) {
	if !isMethodPUT(w, r) {
		return
	}

	id, okId := ParseId(w, r)
	if !okId {
		return
	}

	pet, okUsr := getPetById(id, w)
	if !okUsr {
		return
	}

	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		return
	}

	mydb.Database.Db.Save(&pet)

}
func apiGetAllPets(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}
	pets, okPets := getPets(w)
	if !okPets {
		return
	}
	SendData(pets, w)
}

func getPets(w http.ResponseWriter) ([]mydb.Pet, bool) {

	pets, err := mydb.FindPets()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	} else if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	return pets, true
}

func apiGetPet(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}
	petId, okId := ParseId(w, r)
	if !okId {
		return
	}

	pet, okPet := getPetById(petId, w)
	if !okPet {
		return
	}
	SendData(pet, w)
}

func getPetById(petId int, w http.ResponseWriter) (*mydb.Pet, bool) {

	pet, err := mydb.FindPetById(petId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	} else if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	return pet, true
}

func apiDeletePet(w http.ResponseWriter, r *http.Request) {
	if !isMethodDELETE(w, r) {
		return
	}
	id, okId := ParseId(w, r)
	if !okId {
		return
	}

	pet, okPet := getPetById(id, w)
	if !okPet {
		return
	}

	mydb.Database.Db.Unscoped().Delete(&pet)
	w.WriteHeader(http.StatusOK)
}

//func ShowSkill(name string) {
//	var pet mydb.Pet
//	mydb.Database.Db.Where("name = ?", name).First(&pet)
//	fmt.Printf("%d counter\n", pet.Counter)
//	fmt.Printf("%s показывает что умеет\n", name)
//	pet.Counter += 1
//	fmt.Printf("%d counter\n", pet.Counter)
//	mydb.Database.Db.Save(&pet)
//}

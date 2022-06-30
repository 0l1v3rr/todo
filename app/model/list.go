package model

import (
	"fmt"

	"github.com/0l1v3rr/todo/app/util"
)

type List struct {
	Id      int    `json:"id" gorm:"primaryKey" example:"1"`
	OwnerId int    `json:"ownerId" gorm:"not null;column:owner_id" example:"1"`
	ImageId int    `json:"imageId" gorm:"not null;column:image_id" example:"1"`
	Name    string `json:"name" gorm:"not null" example:"List"`
	Url     string `json:"url" gorm:"unique" example:"list-1"`
}

func (list List) Validate() (bool, string) {
	// if the name is less than 3 characters
	if len(list.Name) < 3 {
		return false, "The name has to be at least 3 characters long."
	}

	// if the name is too long
	if len(list.Name) > 32 {
		return false, "The name can be maximum 32 characters long."
	}

	return true, ""
}

func GetLists(ownerId int) ([]List, error) {
	var lists []List

	// getting the lists from the db where
	// the result-set should be ordered in descending order by id
	tx := DB.Where("owner_id = ?", ownerId).Order("id DESC").Find(&lists)
	if tx.Error != nil {
		return []List{}, tx.Error
	}

	return lists, nil
}

func GetListByUrl(url string) (List, error) {
	var list List

	// getting the list from the db by the specified url
	tx := DB.Where("url = ?", url).First(&list)
	if tx.Error != nil {
		return List{}, tx.Error
	}

	return list, nil
}

func GetListOwnerId(listId int) int {
	// getting the list from the db
	var list List
	tx := DB.Where("id = ?", listId).First(&list)

	// if there is an error, the list doesn't exist, so we return -1
	if tx.Error != nil {
		return -1
	}

	return list.OwnerId
}

func ListExists(id int) (List, bool) {
	// getting the list from the db
	var list List
	tx := DB.Where("id = ?", id).First(&list)

	// if there is an error, the list doesn't exist
	if tx.Error != nil {
		return List{}, false
	}

	// if the id is 0, the list doesn't exist
	if list.Id == 0 {
		return List{}, false
	}

	// the list exists
	return list, true
}

func CreateList(list List) (List, error) {
	// overriding the url
	list.Url = fmt.Sprintf("%s-%s", util.CreateUrlByTitle(list.Name), util.GenerateHash(8))
	list.ImageId = 1

	// creating the list in the db
	tx := DB.Create(&list)
	return list, tx.Error
}

func EditList(list List) (List, error) {
	// saving the edited list in the db
	tx := DB.Save(&list)
	return list, tx.Error
}

func DeleteList(id int) error {
	// deleting the list from the db
	tx := DB.Unscoped().Delete(&List{}, id)
	return tx.Error
}

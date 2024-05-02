package loket

import (
	"antrian/utils"
	"time"
)

var loketlist []LoketTable

func GetListLoket(filter FilterGetList) utils.ServiceReturn {
	db := utils.InitConnection().DB
	if len(filter.NamaLoket) > 0 {
		db = db.Where("nama_loket LIKE  ?", "%"+filter.NamaLoket+"%")
	}

	err := db.Where("status_loket", filter.Status).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order("created_at DESC").
		Find(&loketlist)

	if err.Error != nil {
		return utils.ServiceReturn{
			Status:  false,
			Message: err.Error.Error(),
			Data:    nil,
		}
	}

	return utils.ServiceReturn{
		Status:  true,
		Message: "",
		Data:    loketlist,
	}
}

func (create Create) InsertNewLoket() {
	db := utils.InitConnection().DB
	testInsert := LoketTable{
		NamaLoket:   create.NamaLoket,
		StatusLoket: "1",
		CreatedAt:   time.Now(),
	}

	db.Create(&testInsert)
}

func (updateDataStatus UpdateDataStatus) UpdateColStatus() bool {
	var LoketTable LoketTable
	db := utils.InitConnection().DB
	err := db.Where("id_loket = ?", updateDataStatus.IDLoket).First(&LoketTable).Error
	if err != nil {
		if err.Error() == "record not found" {
			return false
		}
		return false
	}

	LoketTable.StatusLoket = updateDataStatus.StatusLoket
	db.Updates(LoketTable)
	return true
}

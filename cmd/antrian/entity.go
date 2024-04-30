package antrian

import "time"

type AntrianTable struct {
	IDLoket        int       `json:"id_loket" gorm:"primary_key"`
	IDLoketAntrian string    `json:"id_loket_antrian"`
	NoAntrian      string    `json:"no_antrian"`
	StatusAntrian  string    `json:"status_antrian"`
	CreatedAt      time.Time `json:"created_at"`
}

func (a *AntrianTable) TableName() string {
	return "antrian_loket" // This overrides the default table name to "antrian" during migration
}

package loket

import (
	"time"
)

type LoketTable struct {
	IDLoket     int       `json:"id_loket" gorm:"primaryKey"`
	NamaLoket   string    `json:"nama_loket"`
	StatusLoket string    `json:"status_loket"`
	CreatedAt   time.Time `json:"created_at"`
}

func (l *LoketTable) TableName() string {
	return "loket"
}

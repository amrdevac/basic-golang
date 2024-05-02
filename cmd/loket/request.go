package loket

import "time"

type Create struct {
	NamaLoket   string    ` validate:"required" json:"nama_loket"`
	StatusLoket int       ` json:"status_loket"`
	CreatedAt   time.Time ` json:"created_at"`
}

type UpdateDataStatus struct {
	IDLoket     int    `validate:"required"  json:"id_loket"`
	StatusLoket string `validate:"required"  json:"status_loket"`
}

type FilterGetList struct {
	NamaLoket string `json:"nama_loket"`
	Status    string `json:"status"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
}

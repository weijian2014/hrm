package db

import (
	"time"
)

type Post struct {
	Id        uint64    `json:"id" description:"岗位ID"`
	Name      string    `json:"name" description:"岗位名称"`
	UpdatedAt time.Time `json:"updated_at" description:"更新时间"`
}

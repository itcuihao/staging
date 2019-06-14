package handle

import "github.com/itcuihao/staging/s1/storage"

type Handle struct {
	DB *storage.DB
}

func NewHandle(db *storage.DB) *Handle {
	return &Handle{DB: db}
}

package handle

import (
	"fmt"
	"net/http"

	"github.com/itcuihao/staging/s1/storage"
)

type UserHandle struct {
	*Handle
}

func NewUserHandle(db *storage.DB) *UserHandle {
	return &UserHandle{
		Handle: NewHandle(db),
	}
}

func (h *UserHandle) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"][0]
	data, err := h.DB.GetUserById(id)
	if err != nil {
		return
	}
	fmt.Fprint(w, data)
}

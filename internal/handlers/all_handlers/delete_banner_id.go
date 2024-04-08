package allHandlers

import "net/http"

func (h *Handlers) DeleteBannerId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteBannerId"))
}

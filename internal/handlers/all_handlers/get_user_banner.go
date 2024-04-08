package allHandlers

import "net/http"

func (h *Handlers) GetUserBanner(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUserBanner"))
}

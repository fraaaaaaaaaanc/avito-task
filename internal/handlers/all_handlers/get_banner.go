package allHandlers

import "net/http"

func (h *Handlers) GetBanner(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetBanner"))
}

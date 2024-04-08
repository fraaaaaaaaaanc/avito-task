package allHandlers

import "net/http"

func (h *Handlers) PostBanner(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PostBanner"))
}

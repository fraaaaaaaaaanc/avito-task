package allHandlers

import "net/http"

func (h *Handlers) PatchBannerId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PatchBannerId"))
}

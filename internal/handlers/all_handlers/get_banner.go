package allHandlers

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/model/hanlders_model"
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
)

func (h *Handlers) GetBanner(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logger.Error("Error parse query")
	}

	filter := new(hlModel.GetUserBannerModel)
	if err := schema.NewDecoder().Decode(filter, r.Form); err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
	}

	// Do something with filter
	fmt.Printf("%+v", filter)
}

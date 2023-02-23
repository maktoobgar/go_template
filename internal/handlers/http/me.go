package httpHandlers

import (
	"encoding/json"
	"net/http"

	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/models"
)

func me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	userBytes, _ := json.Marshal(user.UserCore)
	w.Write(userBytes)
}

var Me = g.Handler{
	Handler: me,
}

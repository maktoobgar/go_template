package httpHandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	g "github.com/maktoobgar/go_template/internal/global"
)

type hiResponse struct {
	Message string `json:"message"`
}

func hi(w http.ResponseWriter, r *http.Request) {
	splits := strings.Split(r.URL.Path, "/")
	name := splits[2]
	msg := fmt.Sprintf("✋ %s", name)
	res := hiResponse{
		Message: msg,
	}
	resBytes, _ := json.Marshal(res)
	w.Write(resBytes)
}

var Hi = g.Handler{
	Handler: hi,
}

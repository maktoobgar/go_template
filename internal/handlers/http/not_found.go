package httpHandlers

import (
	"net/http"

	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/translator"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	panic(errors.New(errors.NotFoundStatus, errors.DoNothing, translate("PageNotFound")))
}

var NotFound = g.Handler{
	Handler: notFound,
}

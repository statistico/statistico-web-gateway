package rest

import (
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	"net/http"
)

type ResultHandler struct {
	composer composer.ResultComposer
}

func(t *TeamHandler) Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

package web

import (
    "net/http"

    "github.com/amichelins/goexpert_labs_cloudrun/internal/dto"
    "github.com/amichelins/goexpert_labs_cloudrun/internal/infra/request"
    "github.com/amichelins/goexpert_labs_cloudrun/internal/presenters"
)

func TempCepHandler(w http.ResponseWriter, r *http.Request) {
    var sCep = r.FormValue("cep")
    var sKey string

    sKey, ok := r.Context().Value("key").(string)

    if !ok {
        http.Error(w, presenters.ToJson(dto.GeneralResponseError{Msg: "Missing Weather Api Key"}), http.StatusInternalServerError)
        return
    }

    Request := request.NewRequest(presenters.SoDigitos(sCep), sKey)

    if !Request.Valida() {
        http.Error(w, presenters.ToJson(dto.GeneralResponseError{Msg: "invalid zipcode"}), http.StatusUnprocessableEntity)
        return
    }

    err := Request.ViaCep()

    if err != nil && err != request.ErrNoCep {
        http.Error(w, presenters.ToJson(dto.GeneralResponseError{Msg: "An error has occurred" + err.Error()}), http.StatusInternalServerError)
        return
    }

    if err == request.ErrNoCep {
        http.Error(w, presenters.ToJson(dto.GeneralResponseError{Msg: "can not find zipcode"}), http.StatusNotFound)
        return
    }

    err = Request.GetTemperatura()

    if err != nil {
        http.Error(w, presenters.ToJson(dto.GeneralResponseError{Msg: "An error has occurred" + err.Error()}), http.StatusInternalServerError)
        return
    }

    _, _ = w.Write([]byte(presenters.ToJson(dto.GeneralResponse{TempC: Request.GetTempC(), TempF: Request.GetTempF(), TempK: Request.GetTempK()})))
}

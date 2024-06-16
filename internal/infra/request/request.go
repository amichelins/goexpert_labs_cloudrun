package request

import (
    "crypto/tls"
    "encoding/json"
    "errors"
    "io"
    "net/http"
    "net/url"
    "strings"

    "github.com/amichelins/goexpert_labs_cloudrun/internal/dto"
    "github.com/valyala/fastjson"
)

var ErrNoCep = errors.New("Cep não foi achado")

type Request struct {
    cep         string
    cidade      string
    key         string
    temperatura float64
}

func NewRequest(sCep string, sKey string) *Request {
    return &Request{cep: sCep, key: sKey}
}

// ViaCep Recebe um CEP e consulta VIACEP para saber os dados do cep
//
// PARAMETERS
//
//     sCep string Cep para obter os dados
//
// RETURN
//
//     *dto.ViaCepInput Dados do CEP
//
//     error Erro ocorrido ou nil
//
func (r *Request) ViaCep() error {
    var CepBrasil dto.ViaCep
    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
    request, err := http.Get("https://viacep.com.br/ws/" + r.cep + "/json/")

    if err != nil {
        return err
    }

    // Pegamos as resposta
    data, err := io.ReadAll(request.Body)

    if err != nil {
        return err
    }
    defer request.Body.Close()

    if strings.Contains(string(data), `"erro"`) {
        return ErrNoCep
    }

    err = json.Unmarshal(data, &CepBrasil)

    if err != nil {
        return err
    }

    r.cidade = CepBrasil.Localidade

    return nil

}

func (r *Request) Valida() bool {

    if len(r.cep) != 8 || len(r.key) == 0 {
        return false
    }
    return true
}

func (r *Request) GetTemperatura() error {
    var p fastjson.Parser

    req, err := http.NewRequest("GET", "http://api.weatherapi.com/v1/forecast.json?key="+r.key+"&q="+url.QueryEscape(r.cidade)+"&days=1&aqi=no&alerts=no", nil)

    if err != nil {
        return err
    }

    req.Header.Add("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        return err
    }

    // Pegamos as resposta
    data, err := io.ReadAll(resp.Body)

    if err != nil {
        return err
    }
    defer resp.Body.Close()

    Values, err := p.ParseBytes(data)

    if err != nil {
        return err
    }

    Temp, err := Values.GetObject("current").Get("temp_c").Float64()

    if err != nil {
        return err
    }

    r.temperatura = Temp

    return nil
}

func (r *Request) GetTempC() float64 {
    return r.temperatura
}

func (r *Request) GetTempF() float64 {
    return (r.temperatura * 1.8) + 32
}

func (r *Request) GetTempK() float64 {
    return r.temperatura + 273
}

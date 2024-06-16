package main

import (
    "net/http"

    "github.com/amichelins/goexpert_labs_cloudrun/configs"
    "github.com/amichelins/goexpert_labs_cloudrun/internal/infra/web"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    // Carregamos as configurações
    configs, err := configs.LoadConfig(".")

    if err != nil {
        panic("Application can't iniciate. Configuration error")
    }

    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // Enviamos a WeatherApiKey pelo Chi
    r.Use(middleware.WithValue("key", configs.WeatherApiKey))

    r.Get("/temp_cep", web.TempCepHandler)
    _ = http.ListenAndServe(configs.WebServerPort, r)
}

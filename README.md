# goexpert_labs_cloudrun
Laboratorio de Cloud Run

-- Configurações
Na pasta /cmd  tem o arquivo  .env
Tem duas configurações:
WEB_SERVER_PORT Porta que o http server vai utilizar. Padrão 8000
WEATHER_API_KEY Chave para autenticar na API api.weatherapi.com


--- Rodar o Sistema
Ir na pasta /cmd
go run main.go

Na pasta api tem dois arquivos
cep.http  para testar o sistema
teste.http para testar a API diretamente


-- DOCKER
- Foi adicionado ao http configuração para que dentro do Docker não desse erro ao chamar a ViaCep

docker build -t tempcep:latest -f cmd/Dockerfile.prod ./
docker run --rm -p 8080:8080 tempcep:latest

-- Google Cloud Run

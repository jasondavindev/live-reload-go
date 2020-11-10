# Hacktoberfest 2020

## Objetivo

Esse repositório tem como objetivo participar do evento [Hacktoberfest](https://hacktoberfest.digitalocean.com/) da DigitalOcean, construindo uma aplicação capaz de fazer hot reload de códigos para ambiente de desenvolvimento. Você ver como exemplo [nodemon](https://github.com/remy/nodemon), uma biblioteca JavaScript.

## Funcionalidades
- Escuta eventos de alteração (permissão, nome, conteúdo) em diretórios e arquivos
- Capaz de excluir arquivos do watcher
- Executa uma lista de comandos quando emitido um evento de alteração
- Arquivo de configuração
- Flags de configuração

## Buildando

Para buildar a aplicação, é necessário ter a versão 1.15 do [Golang](https://golang.org/dl/).

```bash
go build main.go -o live_reload
```

Com o binário construído, é necessário apenas executá-lo via command line.
```bash
./live_reload
./live_reload -config-file=/tmp/config.yml # passando o arquivo de configuração
```
Obs: por padrão, a aplicação tentará ler o arquivo `config.yml` na pasta local.

# Marys

[![Go Report Card](https://goreportcard.com/badge/github.com/indrasaputra/marys)](https://goreportcard.com/report/github.com/indrasaputra/marys)
[![Workflow](https://github.com/indrasaputra/marys/workflows/Deploy/badge.svg)](https://github.com/indrasaputra/marys/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/indrasaputra/marys.svg)](https://pkg.go.dev/github.com/indrasaputra/marys)

Marys is a webhook to send message or notification to [Telegram](https://telegram.org/).
Marys is derived from one of [One Piece](https://en.wikipedia.org/wiki/One_Piece) character: [The Marys](https://onepiece.fandom.com/wiki/Beasts_Pirates#Marys).
Its main purpose is currently as a notifier service for my personal projects.

## Owner

[Indra Saputra](https://github.com/indrasaputra)

## Usage

Send a POST request to the Marys endpoint. For example, the endpoint is `http://localhost:8080/notifications`, then send this JSON to the endpoint:

```json
{
    "sender": "Indra",
    "message": "Service X has been down for the last 5 minutes"
}
```

`sender` and `message` are required.

## How to Run

Since this project depends on Telegram, so we need to get ChatID (it can be channel, group, or personal message).

This repository also provides `main.go` that can be run in development.

```
$ TELEGRAM_RECIPIENT_ID=<chat-id> \
  TELEGRAM_TOKEN=<telegram-bot-token> \
  TELEGRAM_URL=https://api.telegram.org/bot \
  PORT=8080 \
  go run cmd/api/main.go
```

## Deployment

Currently, this project is deployed in [Google Cloud Functions](https://cloud.google.com/functions).
The deployment process definiton is stated and ruled in [Github Actions](.github/workflows/deploy.yml).
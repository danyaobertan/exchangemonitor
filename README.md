# Exchange Monitor

## Introduction
Exchange Monitor is a robust web application designed to monitor the exchange rate of USD to UAH by using the NBU API and to alert subscribers daily about the exchange rate on their email. 
The application is built using Go, PostgreSQL, and Docker. Mailtrap is used to send emails to subscribers.

## Features
- Monitor exchange rate of USD to UAH
- Alert subscribers daily about the exchange rate on their email (new subscribers will also be notified each day, if the application is running)
- Graceful shutdown
- Beautiful and responsive email template

## Requirements
- Docker
- Go (if running without Docker)

## Installation
Begin by cloning the repository to your local machine:

```bash
git clone https://github.com/danyaobertan/exchangemonitor.git
```

```bash
cd exchangemonitor
```

## Makefile
Take advantage of the Makefile to streamline operations:

```bash
make help
```

## Setup
Setup project with docker-compose

```bash
docker-compose up -d
```

Setup only database with docker-compose

```bash
docker-compose -f docker-compose-db-only.yaml up -d
```

## Test
```bash
go test ./... -v
```

## Database
- Actually, only one table is used by the application. The table is called `subscribers` it contains subscribers email addresses.
- As the future development, the table `rates` could be used to store the historical data of the exchange rate.
- Also `email_notifications` table could be used to store the information about all the notification events.
![Database](./docs/img.png)

## HTML Email Template
### Phone view
![Phone view](./docs/img_1.png)
### Tablet view
![Tablet view](./docs/img_2.png)
### Desktop view
![Desktop view](./docs/img_3.png)
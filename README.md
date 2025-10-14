# Budget Control Application

Backend for the Budget Control Application written in Go

## Stack

- Golang
- PostgreSQL
- [Chi](https://go-chi.io/#/pages/getting_started)

## Installation

Add the following to your `.env` file:

```env
PORT=${SERVER_PORT}
DATABASE_URL=postgresql://${USER}:${PASSWORD}@${HOST}:${PORT}/${DATABASE}?schema=public
AUTH_SERVER=${BETTER_AUTH_SERVER}
```

## Installation

```bash
go mod tidy
make run
```


## Tools

- [Neovim](https://neovim.io/)
- [Tmux](https://github.com/tmux/tmux/wiki)
- [Postgres](https://www.postgresql.org/)

## Developed by

- [Andres Court](https://linkedin.com/in/alcb1310)


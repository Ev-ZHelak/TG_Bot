# TG_Bot
The project is under development . . .
#### Useful links:
- [Telebot documentation](https://godoc.org/gopkg.in/telebot.v3)
- [Viper reading from JSON, TOML, YAML, HCL, envfile and Java properties config files](https://github.com/spf13/viper/tree/master)

#### Current project tree

```
TG_Bot
├─ README.md
├─ cmd
│  └─ app
│     └─ main.go
├─ config
│  └─ config.go
├─ go.mod
├─ go.sum
├─ internal
│  ├─ bot
│  │  ├─ games
│  │  │  ├─ config
│  │  │  ├─ millionaire
│  │  │  │  ├─ engine
│  │  │  │  │  ├─ questions.go
│  │  │  │  │  └─ session.go
│  │  │  │  ├─ handlers
│  │  │  │  └─ keyboards
│  │  │  └─ registry.go
│  │  ├─ handlers
│  │  │  ├─ callbacks.go
│  │  │  ├─ commands.go
│  │  │  └─ messages.go
│  │  ├─ initbot.go
│  │  ├─ keyboards
│  │  │  ├─ inline.go
│  │  │  └─ reply.go
│  │  └─ menu
│  │     └─ create.go
│  ├─ models
│  │  └─ user.go
│  ├─ repository
│  │  ├─ millionaire
│  │  │  └─ questions.json
│  │  └─ postgres
│  ├─ services
│  │  ├─ auth
│  │  └─ notification
│  └─ utils
│     └─ logger
└─ scripts
   └─ psql
      └─ install_psql.sh

```
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
│  │  ├─ bot.go
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
│  │  ├─ keyboards
│  │  │  ├─ inline.go
│  │  │  └─ reply.go
│  │  └─ menu
│  │     └─ create.go
│  ├─ db
│  │  ├─ pool.go
│  │  └─ stats.go
│  ├─ models
│  │  └─ user.go
│  ├─ repository
│  │  └─ millionaire
│  │     └─ questions.json
│  ├─ services
│  │  ├─ auth
│  │  └─ notification
│  ├─ sql
│  │  ├─ init
│  │  │  └─ create_user_tables.sql
│  │  └─ queries
│  └─ utils
│     └─ logger
└─ scripts
   └─ psql
      └─ install_psql.sh

```
package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/chrispaul1/blog/internal/commands"
	"github.com/chrispaul1/blog/internal/config"
	"github.com/chrispaul1/blog/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	var newState commands.State
	configStruct := config.ReadConfigFile()
	db, err := sql.Open("postgres", configStruct.URL)
	dbQueries := database.New(db)
	newState.C = &configStruct
	newState.DB = dbQueries
	cmds := &commands.Commands{
		HandlerFuncs: make(map[string]func(*commands.State, commands.Command) error),
	}
	cmds.Register("login", commands.HandleLogin)
	cmds.Register("register", commands.HandleRegister)
	cmds.Register("reset", commands.HandleReset)
	cmds.Register("users", commands.Users)
	cmds.Register("agg", commands.Agg)
	cmds.Register("addfeed", commands.MiddlewareLoggedIn(commands.HandleAddFeed))
	cmds.Register("feeds", commands.HandleFeeds)
	cmds.Register("follow", commands.MiddlewareLoggedIn(commands.HandleFollow))
	cmds.Register("following", commands.MiddlewareLoggedIn(commands.HandleFollowing))
	cmds.Register("unfollow", commands.MiddlewareLoggedIn(commands.HandleUnfollow))
	cmds.Register("browse", commands.MiddlewareLoggedIn(commands.HandleBrowse))

	if len(os.Args) < 2 {
		fmt.Println("Error, Invalid number of arguments")
		os.Exit(1)
	}

	commandName := os.Args[1]
	argsSlice := os.Args[2:]

	err = cmds.Execute(&newState, commandName, argsSlice)

	if err != nil {
		fmt.Printf("\n%s", err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}

	//config.SetName("Chris", &configStruct)
	//newConfigStruct := config.ReadConfigFile()
	//fmt.Println(newConfigStruct)
}

//goose -dir ./sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator" down

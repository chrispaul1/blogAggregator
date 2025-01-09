package commands

import (
	"errors"
	"fmt"

	"github.com/chrispaul1/blog/internal/config"
	"github.com/chrispaul1/blog/internal/database"
)

type State struct {
	DB      *database.Queries
	C       *config.Config
	FeedUrl string
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	HandlerFuncs map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.HandlerFuncs[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	if returnedCommand, exists := c.HandlerFuncs[cmd.Name]; exists {
		err := returnedCommand(s, cmd)
		if err != nil {
			return err
		}
	} else {
		errorStr := fmt.Sprintf("Unknown Command : %s", cmd.Name)
		return errors.New(errorStr)
	}

	return nil
}

func (c *Commands) Execute(s *State, commandName string, args []string) error {
	newCommand := Command{
		Name: commandName,
		Args: args,
	}
	return c.Run(s, newCommand)
}

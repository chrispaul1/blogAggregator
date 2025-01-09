package commands

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/chrispaul1/blog/internal/config"
)

func HandleLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("login handler expects a username")
	}
	name := cmd.Args[0]
	_, err := s.DB.GetUser(context.Background(), name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: user '%v' not found\n", name)
		os.Exit(1)
	}

	err = config.SetName(name, s.C)
	if err != nil {
		return err
	}
	return nil
}

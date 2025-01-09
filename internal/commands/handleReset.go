package commands

import (
	"context"
	"fmt"
	"os"
)

func HandleReset(s *State, cmd Command) error {
	err := s.DB.DeleteUsers(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: users could not be deleted '%s", err)
		os.Exit(1)
	}
	return nil
}

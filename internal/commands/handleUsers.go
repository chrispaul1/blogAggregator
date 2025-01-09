package commands

import (
	"context"
	"fmt"
	"os"
)

func Users(s *State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if len(users) == 0 {
		fmt.Println("No users were found")
		os.Exit(0)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: users could not be retreieved '%s", err)
		os.Exit(1)
	}
	for _, user := range users {
		if user.Name == s.C.User {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}

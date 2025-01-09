package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/chrispaul1/blog/internal/database"
)

func MiddlewareLoggedIn(handleFunc func(s *State, cmd Command, user database.User) error) func(s *State, cmd Command) error {
	return func(s *State, cmd Command) error {
		currentUser, err := s.DB.GetUser(context.Background(), s.C.User)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error, could not get user struct '%v'\n", err)
			return err
		}
		return handleFunc(s, cmd, currentUser)
	}
}

package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/chrispaul1/blog/internal/config"
	"github.com/chrispaul1/blog/internal/database"
	"github.com/google/uuid"
)

func HandleRegister(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("Register expects a name")
	}
	username := cmd.Args[0]
	_, err := s.DB.GetUser(context.Background(), username)
	if err == nil {
		fmt.Println("This user already exists")
		os.Exit(1)
	}
	userStruct := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	newUser, err := s.DB.CreateUser(context.Background(), userStruct)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: user '%v' was not created\n", username)
		os.Exit(1)
	}
	err = config.SetName(username, s.C)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: user '%v' could not be set\n", username)
		os.Exit(1)
	}

	fmt.Println("User was created :\n", newUser)
	return nil
}

package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	err := s.db.ClearUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error reseting table")
	}
	fmt.Println("Table reset was successful")
	return nil
}

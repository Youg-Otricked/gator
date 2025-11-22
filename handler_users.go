package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error selecting users from table")
	}
	if users == nil {
		fmt.Println("No users registered")
		return nil
	}
	for _, user := range users {
		cur := s.cfg.CurrentUserName
		if user.Name == cur {
			fmt.Printf("%s (current)", user.Name)
		} else {
			fmt.Print(user)
		}
	}
	return nil
}

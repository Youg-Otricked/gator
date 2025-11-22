package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error selecting feeds from table")
	}
	if feeds == nil {
		fmt.Println("No feeds registered")
		return nil
	}
	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		user, err2 := s.db.GetUserF(context.Background(), feed.UserID)
		if err2 != nil {
			return err2
		}
		fmt.Println(user.Name)
	}
	return nil
}

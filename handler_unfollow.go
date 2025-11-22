package main

import (
	"context"
	"fmt"

	"github.com/Youg-Otricked/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	feedid, err2 := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err2 != nil {
		return fmt.Errorf("error getting feed: %w", err2)
	}
	err := s.db.UnfollowFeedForUser(context.Background(), database.UnfollowFeedForUserParams{
		UserID: user.ID,
		FeedID: feedid.ID,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}
	fmt.Printf("Feed %s was successfuly unfolowed\n", feedid.Name)
	return nil
}

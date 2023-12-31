package services

import (
	"context"

	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/repos"
)

type FeedService struct {
	feedRepo repos.FeedRepo
}

func NewFeedService(
	feedRepo repos.FeedRepo,
) *FeedService {
	return &FeedService{
		feedRepo: feedRepo,
	}
}

func (f *FeedService) GetFeed(ctx context.Context, userId int64, page, limit int) ([]*entities.User, error) {
	return f.feedRepo.GetFeed(ctx, userId, page, limit)
}

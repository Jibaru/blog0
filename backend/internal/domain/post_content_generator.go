package domain

import "context"

type PostContentGenerator interface {
	GenerateSummary(ctx context.Context, rawPostContent string) (string, error)
	GenerateTags(ctx context.Context, rawPostContent string) ([]string, error)
}

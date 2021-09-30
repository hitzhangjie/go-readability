package linter

import (
	"context"

	"github.com/hitzhangjie/go-readability/pkg/result"
)

type Linter interface {
	Run(ctx context.Context, lintCtx *Context) ([]result.Issue, error)
	Name() string
	Desc() string
}

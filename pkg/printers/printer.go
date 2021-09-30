package printers

import (
	"context"

	"github.com/hitzhangjie/go-readability/pkg/result"
)

type Printer interface {
	Print(ctx context.Context, issues []result.Issue) error
}

package processors

import (
	"github.com/hitzhangjie/go-readability/pkg/result"
)

type Processor interface {
	Process(issues []result.Issue) ([]result.Issue, error)
	Name() string
	Finish()
}

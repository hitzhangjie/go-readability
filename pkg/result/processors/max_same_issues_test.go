package processors

import (
	"testing"

	"github.com/hitzhangjie/go-readability/pkg/config"
	"github.com/hitzhangjie/go-readability/pkg/logutils"
	"github.com/hitzhangjie/go-readability/pkg/result"
)

func TestMaxSameIssues(t *testing.T) {
	p := NewMaxSameIssues(1, logutils.NewStderrLog(""), &config.Config{})
	i1 := result.Issue{
		Text: "1",
	}
	i2 := result.Issue{
		Text: "2",
	}

	processAssertSame(t, p, i1)  // ok
	processAssertSame(t, p, i2)  // ok: another
	processAssertEmpty(t, p, i1) // skip
}

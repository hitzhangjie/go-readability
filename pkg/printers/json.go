package printers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hitzhangjie/go-readability/pkg/logutils"
	"github.com/hitzhangjie/go-readability/pkg/report"
	"github.com/hitzhangjie/go-readability/pkg/result"
)

type JSON struct {
	rd *report.Data
}

func NewJSON(rd *report.Data) *JSON {
	return &JSON{
		rd: rd,
	}
}

type JSONResult struct {
	Issues []result.Issue
	Report *report.Data
}

func (p JSON) Print(ctx context.Context, issues []result.Issue) error {
	res := JSONResult{
		Issues: issues,
		Report: p.rd,
	}

	outputJSON, err := json.Marshal(res)
	if err != nil {
		return err
	}

	fmt.Fprint(logutils.StdOut, string(outputJSON))
	return nil
}

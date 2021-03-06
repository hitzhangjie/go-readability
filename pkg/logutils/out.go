package logutils

import (
	"github.com/fatih/color"
	colorable "github.com/mattn/go-colorable"
)

var StdOut = color.Output // https://github.com/hitzhangjie/go-readability/issues/14
var StdErr = colorable.NewColorableStderr()

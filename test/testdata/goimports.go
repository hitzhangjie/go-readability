//args: -Egoimports
package testdata

import (
	"fmt" // ERROR "File is not `goimports`-ed"
	"github.com/hitzhangjie/go-readability/pkg/config"
)

func Bar() {
	fmt.Print("x")
	_ = config.Config{}
}

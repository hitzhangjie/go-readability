//args: -Egoimports
//config: linters-settings.goimports.local-prefixes=github.com/hitzhangjie/go-readability
package goimports

import (
	"fmt"

	"github.com/hitzhangjie/go-readability/pkg/config"
	"github.com/pkg/errors"
)

func GoimportsLocalTest() {
	fmt.Print("x")
	_ = config.Config{}
	_ = errors.New("")
}

//args: -Egci
//config: linters-settings.gci.local-prefixes=github.com/hitzhangjie/go-readability
package gci

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

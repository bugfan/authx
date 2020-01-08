package authx

import (
	"authx/settings"
	"testing"
)

func TestAuthx(t *testing.T) {
	Run(settings.Get("authx_host") + ":" + settings.Get("authx_port"))
}

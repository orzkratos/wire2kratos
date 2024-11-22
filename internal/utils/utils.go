package utils

import (
	"strings"

	"github.com/yyle88/done"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
)

func PWD() string {
	root := strings.TrimSpace(string(done.VAE(osexec.NewCMC().Exec("pwd")).Nice()))
	must.Nice(root)
	osmustexist.MustRoot(root)
	return root
}

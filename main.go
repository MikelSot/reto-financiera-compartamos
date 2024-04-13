package main

import (
	_ "embed"

	"github.com/MikelSot/reto-financiera-compartamos/bootstrap"
)

//go:embed boot.yaml
var boot []byte

func main() {
	bootstrap.Run(boot)
}

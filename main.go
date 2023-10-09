package main

import (
	"github.com/T-sugiyama-dev/consider_buy/route"
)

func main() {
	r := route.Router()
	r.Run(":8080")
}

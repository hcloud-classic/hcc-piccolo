package main

import (
	"hcc/piccolo/action/graphql"
	hccGatewayEnd "hcc/piccolo/end"
	hccGatewayInit "hcc/piccolo/init"
)

func init() {
	err := hccGatewayInit.MainInit()
	if err != nil {
		panic(err)
	}
}

func main() {
	defer func() {
		hccGatewayEnd.MainEnd()
	}()

	graphql.Init()
}

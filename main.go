package main

import (
	"utahw/gine"
	"utahw/mymysql"
)

func main() {

	mymysql.My_init()
	gine.GinInit()

}

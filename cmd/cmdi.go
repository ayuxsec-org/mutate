package main

import "github.com/ayuxsec/mutate"

func NewCmdi() *Cmdi {
	return &Cmdi{}
}

type Cmdi struct {
	URLsPath   string
	Identifier string
	Opts       mutate.Options
}

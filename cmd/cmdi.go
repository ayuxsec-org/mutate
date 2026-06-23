package main

import "github.com/ayuxsec-org/mutate"

func NewCmdi() *Cmdi {
	return &Cmdi{}
}

type Cmdi struct {
	URLsPath   string
	Identifier string
	Opts       mutate.Options
}

package main

import "github.com/ayuxsec-org/log"

func main() {
	rootCmd := NewCmdi().NewRootCmd()
	rootCmd.SilenceErrors = true
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}

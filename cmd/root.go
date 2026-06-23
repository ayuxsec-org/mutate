package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ayuxsec/mutate"
	"github.com/spf13/cobra"
)

func (cmdi *Cmdi) NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "mutate",
		Short: "mutate url paths and key params",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return cmdi.ValidateRoot()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmdi.RunRoot()
		},
	}
	rootCmd.PersistentFlags().StringVarP(&cmdi.URLsPath, "urls", "l", "", "Path to the input URL list")
	rootCmd.PersistentFlags().BoolVarP(&cmdi.Opts.MutatePaths, "path", "p", false, "Replace URL Paths")
	rootCmd.PersistentFlags().BoolVarP(&cmdi.Opts.MutateKeys, "keys", "k", false, "Replace URL query params")
	rootCmd.PersistentFlags().StringVarP(&cmdi.Identifier, "id", "i", mutate.DefaultIdentifier, "identifer to use")

	return rootCmd
}

func (cmdi *Cmdi) RunRoot() error {
	var rawURLs []string
	if isPipedStdin() {
		rawURLs = fileToSlice(os.Stdin)
	} else {
		rawURLs = fileToSlice(must(os.Open(cmdi.URLsPath)))
	}
	mutator := mutate.New(rawURLs, cmdi.Identifier,
		mutate.NewOpts(cmdi.Opts.MutatePaths, cmdi.Opts.MutateKeys),
	)
	mutatedURLs := mutator.MutateURLs()
	for _, u := range mutatedURLs {
		fmt.Println(u)
	}
	return nil
}

func (cmdi *Cmdi) ValidateRoot() error {
	if cmdi.URLsPath == "" && !isPipedStdin() {
		return errors.New("no urls provided")
	}
	if !cmdi.Opts.MutateKeys && !cmdi.Opts.MutatePaths {
		return errors.New("atleast one of the path mutation or keys mutation argument is required")
	}
	return nil
}

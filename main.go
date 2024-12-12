package main

import (
	"os"

	"github.com/maxwelbm/mdz/pkg/cmd/root"
	"github.com/maxwelbm/mdz/pkg/environment"
	"github.com/maxwelbm/mdz/pkg/factory"
	"github.com/maxwelbm/mdz/pkg/output"
)

func main() {
	env := environment.New()

	f := factory.NewFactory(env)
	cmd := root.NewCmdRoot(f)

	if err := cmd.Execute(); err != nil {
		printErr := output.Errorf(f.IOStreams.Err, err)
		if printErr != nil {
			output.Printf(os.Stderr, "Failed to print error output: "+printErr.Error())

			os.Exit(1)
		}

		os.Exit(1)
	}
}

/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

type GenHelpDocsCmd struct {
	cobraCommand *cobra.Command
	OutDir       string
}

var genHelpDocsCmd = GenHelpDocsCmd{
	cobraCommand: &cobra.Command{
		Use:    "genhelpdocs",
		Short:  "Generate CLI help docs",
		Hidden: true,
	},
}

func init() {
	cmd := genHelpDocsCmd.cobraCommand
	rootCommand.cobraCommand.AddCommand(cmd)

	cmd.Run = func(cmd *cobra.Command, args []string) {
		err := genHelpDocsCmd.Run()
		if err != nil {
			exitWithError(err)
		}
	}

	cmd.Flags().StringVar(&genHelpDocsCmd.OutDir, "out", "", "path to write out to.")
}

func (c *GenHelpDocsCmd) Run() error {
	rootCommand.cobraCommand.DisableAutoGenTag = true
	rootCommand.RegistryPath = ""
	err := doc.GenMarkdownTree(rootCommand.cobraCommand, c.OutDir)

	return err
}

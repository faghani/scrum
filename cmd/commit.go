// Copyright © 2017 Mester
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:     "commit",
	Args:    cobra.MinimumNArgs(1),
	Short:   `Do git commit`,
	RunE:    Commit,
	Example: `  scrum commit "added server.go file"`,
}

func init() {
	rootCmd.AddCommand(commitCmd)
}

func Commit(cmd *cobra.Command, args []string) error {
	var file = []string{"."}

	if args != nil && len(args) > 1 {
		file = args[0:]
	}
	defer git("commit", "-m", args[0])

	for _, a := range file {
		err := git("add", a)
		if err != nil {
			return err
		}
	}
	return nil
}

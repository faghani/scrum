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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	addArgs = []*cobra.Command{
		&cobra.Command{
			Use:     "file",
			Args:    cobra.MinimumNArgs(1),
			Aliases: []string{"f"},
			Short:   "Add file to your project",
			RunE: func(cmd *cobra.Command, args []string) error {
				file, err := os.OpenFile(args[0], os.O_CREATE, 0600)
				if err != nil {
					return err
				}
				file.Close()

				git("add", args[0])
				if len(args) > 1 {
					git("commit", "-m", args[1])
				} else {
					git("commit", "-m", fmt.Sprintf("Added file %s", args[0]))
				}
				if !pushFlg {
					git("push", "-u", upstream, branch)
				}
				return nil
			},
		},
		&cobra.Command{
			Use:     "task",
			Args:    cobra.MinimumNArgs(1),
			Aliases: []string{"t"},
			Short:   "Add task to your project (by default in ice-box)",
			RunE: func(cmd *cobra.Command, args []string) (err error) {
				var name = fmt.Sprintf("%s%cice-box", dbdir, os.PathSeparator)
				if len(args) > 1 {
					if isValidArg(args[1]) {
						name = fmt.Sprintf("%s%c%s", dbdir, os.PathSeparator, args[1])
					}
				}
				if _, err = os.Stat(name); err == nil {
					err = Write(args[0], name, -1)
					if err == nil {
						err = Commit(cmd, []string{"Added new task"})
						if err == nil {
							err = Push(cmd, nil)
						}
					}
				}
				return
			},
		},
		&cobra.Command{
			Use:     "folder",
			Args:    cobra.MinimumNArgs(1),
			Aliases: []string{"d"},
			Short:   "Add folder to your project",
			RunE: func(cmd *cobra.Command, args []string) error {
				err := os.Mkdir(args[0], 0700)
				if err != nil {
					return err
				}
				git("add", args[0])
				if len(args) > 1 {
					Commit(cmd, []string{args[1]})
				} else {
					Commit(cmd, []string{fmt.Sprintf("Created folder %s", args[0])})
				}
				if !pushFlg {
					Push(nil, nil)
				}
				return err
			},
		},
	}
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.MinimumNArgs(1),
	Short: "Add a task, file or folder to your project",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	for _, add := range addArgs {
		addCmd.AddCommand(add)
	}
	rootCmd.AddCommand(addCmd)
}

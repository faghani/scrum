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
	"strconv"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var (
	rmArgs = []*cobra.Command{
		&cobra.Command{
			Use:     "file",
			Args:    cobra.MinimumNArgs(1),
			Aliases: []string{"f"},
			Short:   "Delete file from your project",
			RunE: func(cmd *cobra.Command, args []string) error {
				err := os.Remove(args[0])
				if err != nil {
					return err
				}

				err = Commit(cmd,
					[]string{
						fmt.Sprintf("Deleted '%s' file", args[0]),
					},
				)
				if err == nil && !pushFlg {
					err = Push(nil, nil)
				}
				return err
			},
		},
		&cobra.Command{
			Use:     "task",
			Args:    cobra.MinimumNArgs(2),
			Aliases: []string{"t"},
			Example: `	scrum rm task ice-box 1`,
			Short: "Delete task from your project",
			RunE: func(cmd *cobra.Command, args []string) error {
				p, err := strconv.Atoi(args[1])
				if err == nil {
					name := fmt.Sprintf("%s%c%s", dbdir, os.PathSeparator, args[0])
					_, err = Peek(name, p)
				}

				err = Commit(cmd,
					[]string{
						fmt.Sprintf("Deleted task"),
					},
				)
				if err == nil && !pushFlg {
					err = Push(nil, nil)
				}
				return err
			},
		},
		&cobra.Command{
			Use:     "folder",
			Args:    cobra.MinimumNArgs(1),
			Aliases: []string{"d"},
			Short:   "Delete directory from your project",
			RunE: func(cmd *cobra.Command, args []string) error {
				err := os.RemoveAll(args[0])
				if err != nil {
					return err
				}
				err = Commit(cmd,
					[]string{
						fmt.Sprintf("Deleted '%s' directory", args[0]),
					},
				)
				if err == nil && !pushFlg {
					err = Push(nil, nil)
				}
				return err
			},
		},
	}
	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Removes file, folder or task from your project",
		Example: `	scrum rm task ice-box 1
	scrum rm file something.py`,
		Aliases: []string{"rm", "rmv"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)

func init() {
	for _, r := range rmArgs {
		removeCmd.AddCommand(r)
	}
	rootCmd.AddCommand(removeCmd)
}

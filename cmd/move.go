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

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move task from one column to another",
	Example: `	scrum move ice-box 1 emergency
	scrum move ice-box 3 emergency 4`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Move(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
}

func Move(cmd *cobra.Command, args []string) error {
	if !isValidArg(args[0]) {
		return fmt.Errorf("%s is invalid arg", args[0])
	}
	if !isValidArg(args[2]) {
		return fmt.Errorf("%s is invalid arg", args[2])
	}

	files := []string{
		fmt.Sprintf("%s%c%s", dbdir, os.PathSeparator, args[0]),
		fmt.Sprintf("%s%c%s", dbdir, os.PathSeparator, args[2]),
	}

	// source index
	s, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	// destination index
	p := -1
	if len(args) > 3 {
		p, err = strconv.Atoi(args[3])
		if err != nil {
			return err
		}
	}

	wr, err := Peek(files[0], s)
	if err == nil {
		err = Write(wr, files[1], p)
		err = Commit(cmd, []string{"Moved task"})
		if err == nil {
			err = Push(cmd, nil)
		}
	}

	return err
}

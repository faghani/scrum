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

// swapCmd represents the swap command
var swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Swaps two tasks",
	Example: `	scrum ice-box 1 testing 2`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !isValidArg(args[0]) {
			return fmt.Errorf("%s is invalid arg", args[0])
		}
		if !isValidArg(args[2]) {
			return fmt.Errorf("%s is invalid arg", args[2])
		}

		from := fmt.Sprintf("%s%c%s", dbdir, os.PathSeparator, args[0])
		to := fmt.Sprintf("%s%c%s", dbdir, os.PathSeparator, args[2])

		s, err := strconv.Atoi(args[1])
		if err == nil {
			d, err := strconv.Atoi(args[3])
			if err == nil {
				// source value
				sv, err := Peek(from, s)
				if err != nil {
					return err
				}
				// destination value
				sd, err := Peek(to, d)
				if err != nil {
					return err
				}
				err = Write(sv, to, d)
				if err == nil {
					err = Write(sd, from, s)
					if err == nil {
						err = Commit(cmd, []string{"Swapped two tasks"})
					}
				}
			}
		}
		if err == nil && !pushFlg {
			err = Push(cmd, nil)
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(swapCmd)
}

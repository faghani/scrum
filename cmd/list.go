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
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all task of project",
	RunE:  List,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func List(cmd *cobra.Command, args []string) (err error) {
	var files []string

	if args[0] != "all" {
		if !isValidArg(args[0]) {
			return fmt.Errorf("%s is invalid arg", args[0])
		}
		files = append(files, args[0])
	} else {
		fs, err := ioutil.ReadDir(dbdir)
		if err != nil {
			return err
		}

		for _, f := range fs {
			files = append(files, f.Name())
		}
		fs = nil
	}

	tbl := tablewriter.NewWriter(os.Stdout)
	tbl.SetHeader(files)
	tbl.SetRowSeparator("-")
	tbl.SetColumnSeparator("|")
	tbl.SetBorders(tablewriter.Border{
		Left: true, Top: true,
		Right: true, Bottom: false,
	})
	tbl.SetAlignment(tablewriter.ALIGN_LEFT)

	var x int = 0
	var content = make([][]string, 0, 0)
	for _, name := range files {
		file, err := os.Open(fmt.Sprintf(
			"%s%c%s", dbdir, os.PathSeparator, name))
		if err != nil {
			continue
		}
		reader := bufio.NewReader(file)

		for y := 0; ; y++ {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			content = append(content, make([]string, len(files)))
			content[y][x] = fmt.Sprintf("%d. %s", y, line)
		}

		file.Close()
		x++
	}
	tbl.AppendBulk(content)
	tbl.Render()
	content = nil
	return
}

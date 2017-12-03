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
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	pushFlg  bool
	branch   string
	upstream string
	dbdir    = "./.databases"
	rootCmd  = &cobra.Command{
		Use:   "scrum",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
to quickly create a Cobra application.`,
	}
	validArgs = []string{
		"ice-box", "emergency",
		"progress", "testing",
		"complete",
	}
)

func isValidArg(arg string) bool {
	for _, ar := range validArgs {
		if ar == arg {
			return true
		}
	}
	return false
}

func MakeDir() error {
	return os.Mkdir(dbdir, 0700)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&pushFlg, "no-push", false, "If push flag is activated scrum won't do a push to git")
	rootCmd.PersistentFlags().StringVar(&branch, "branch", "master", "Branch to your push")
	rootCmd.PersistentFlags().StringVar(&upstream, "upstream", "origin", "Upstream to your push")
}

func git(args ...string) (err error) {
	if args == nil {
		return
	}

	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Peek(file string, id int) (string, error) {
	var sel string
	content := make([][]byte, 0)
	c, err := ioutil.ReadFile(file)
	if err == nil {
		cnt := bytes.Split(c, []byte("\n"))
		if id < 0 {
			id = len(cnt)
		}
		if len(cnt) < id {
			id = len(cnt) - 1
		}

		for i := range cnt {
			// void value
			if len(cnt[i]) <= 1 {
				id++
				continue
			}
			if id == i {
				sel = string(cnt[i])
			} else {
				content = append(content, cnt[i])
			}
		}
		err = ioutil.WriteFile(file, bytes.Join(content, []byte("\n")), 0600)
	}

	return sel, err
}

func Write(line, file string, id int) error {
	content := make([][]byte, 0)

	c, err := ioutil.ReadFile(file)
	if err == nil {
		cnt := bytes.Split(c, []byte("\n"))
		if id < 0 {
			id = len(cnt) - 1
		}
		if len(cnt) < id {
			id = len(cnt) - 1
		}

		for i := range cnt {
			if len(cnt[i]) > 0 {
				content = append(content, cnt[i])
			}
			if id == i {
				content = append(content, []byte(line))
			}
		}
		err = ioutil.WriteFile(file, bytes.Join(content, []byte("\n")), 0600)
	}

	return err
}

/*
 * cmd.go
 *
 * Copyright 2018 Bill Zissimopoulos
 */
/*
 * This file is part of golib.
 *
 * It is licensed under the MIT license. The full license text can be found
 * in the License.txt file at the root of this project.
 */

// Package cmd provides (sub-)command functionality for command-line programs.
// This package works closely with the standard library flag package.
package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

// CmdMap encapsulates a (sub-)command map.
type CmdMap struct {
	cmdmap map[string]*Cmd
	cmdlst []string
	mux    sync.Mutex
}

// Cmd encapsulates a (sub-)command.
type Cmd struct {
	// Flag contains the command flag set.
	Flag *flag.FlagSet

	// Main is the function to run when the command is selected.
	Main func(cmd *Cmd, args []string)

	// Use contains the command usage string.
	Use string

	// Desc contains the command description.
	Desc string
}

// Add adds a new command in the command map.
//
// The name parameter is the command name. However if this parameter contains
// a space or newline it is interpreted as described below. Consider:
//
//     NAME ARGUMENTS
//     DESCRIPTION
//
// Then the command name becomes "NAME", the command Use field becomes
// "NAME ARGUMENTS" and the command Desc field becomes "DESCRIPTION".
func (self *CmdMap) Add(name string, main func(*Cmd, []string)) (cmd *Cmd) {
	lines := strings.SplitN(name, "\n", 2)
	use := lines[0]
	desc := ""
	if 2 == len(lines) {
		desc = lines[1]
	}
	name = strings.SplitN(use, " ", 2)[0]
	cmd = &Cmd{Flag: flag.NewFlagSet(name, flag.ExitOnError), Main: main, Use: use, Desc: desc}

	self.mux.Lock()
	defer self.mux.Unlock()
	self.cmdmap[name] = cmd
	self.cmdlst = append(self.cmdlst, name)

	return
}

// Get gets a command by name.
func (self *CmdMap) Get(name string) *Cmd {
	self.mux.Lock()
	defer self.mux.Unlock()
	return self.cmdmap[name]
}

// GetNames gets all command names.
func (self *CmdMap) GetNames() []string {
	self.mux.Lock()
	defer self.mux.Unlock()
	cmdlst := make([]string, len(self.cmdlst))
	copy(cmdlst, self.cmdlst)
	return cmdlst
}

// PrintCmds prints help text for all commands to stderr.
func (self *CmdMap) PrintCmds() {
	for _, name := range self.GetNames() {
		cmd := self.Get(name)
		if nil == cmd {
			continue
		}
		fmt.Fprintln(os.Stderr, "  "+name)
		if "" != cmd.Desc {
			fmt.Fprintln(os.Stderr, "    \t"+cmd.Desc)
		}
	}
}

// NewCmdMap creates a new command map.
func NewCmdMap() *CmdMap {
	return &CmdMap{
		cmdmap: map[string]*Cmd{},
	}
}

// GetFlag gets the value of the named flag.
func (self *Cmd) GetFlag(name string) interface{} {
	if f := self.Flag.Lookup(name); nil != f {
		if g, ok := f.Value.(flag.Getter); ok {
			return g.Get()
		}
	}
	return nil
}

// DefaultCmdMap is the default command map.
var DefaultCmdMap = NewCmdMap()

// Add adds a new command in the default command map.
//
// The name parameter is the command name. However if this parameter contains
// a space or newline it is interpreted as described below. Consider:
//
//     NAME ARGUMENTS
//     DESCRIPTION
//
// Then the command name becomes "NAME", the command Use field becomes
// "NAME ARGUMENTS" and the command Desc field becomes "DESCRIPTION".
func Add(name string, main func(*Cmd, []string)) *Cmd {
	return DefaultCmdMap.Add(name, main)
}

// PrintCmds prints help text for all commands in the default command map
// to stderr.
func PrintCmds() {
	DefaultCmdMap.PrintCmds()
}

// Run parses the command line and executes the specified (sub-)command.
func Run() {
	if !flag.Parsed() {
		flag.Parse()
	}

	arg := flag.Arg(0)
	cmd := DefaultCmdMap.Get(arg)

	if nil == cmd {
		flag.Usage()
		os.Exit(2)
	}

	cmd.Main(cmd, flag.Args()[1:])
}

func help(cmd *Cmd, args []string) {
	if 0 == len(args) {
		flag.Usage()
	} else {
		for _, name := range args {
			cmd := DefaultCmdMap.Get(name)
			if nil == cmd {
				continue
			}
			cmd.Flag.Usage()
		}
	}
	os.Exit(2)
}

func init() {
	Add("help", help)
}

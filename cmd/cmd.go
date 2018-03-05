/*
 * cmd.go
 *
 * Copyright 2018 Bill Zissimopoulos
 */

package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

type CmdMap struct {
	cmdmap map[string]*Cmd
	cmdlst []string
	mux    sync.Mutex
}

type Cmd struct {
	Flag *flag.FlagSet
	Main func(cmd *Cmd, args []string)
	Use  string
	Desc string
}

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

func (self *CmdMap) Get(name string) *Cmd {
	self.mux.Lock()
	defer self.mux.Unlock()
	return self.cmdmap[name]
}

func (self *CmdMap) GetNames() []string {
	self.mux.Lock()
	defer self.mux.Unlock()
	cmdlst := make([]string, len(self.cmdlst))
	copy(cmdlst, self.cmdlst)
	return cmdlst
}

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

func NewCmdMap() *CmdMap {
	return &CmdMap{
		cmdmap: map[string]*Cmd{},
	}
}

func (self *Cmd) GetFlag(name string) interface{} {
	if f := self.Flag.Lookup(name); nil != f {
		if g, ok := f.Value.(flag.Getter); ok {
			return g.Get()
		}
	}
	return nil
}

var DefaultCmdMap = NewCmdMap()

func Add(name string, main func(*Cmd, []string)) *Cmd {
	return DefaultCmdMap.Add(name, main)
}

func PrintCmds() {
	DefaultCmdMap.PrintCmds()
}

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

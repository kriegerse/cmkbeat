package cmd

import cmd "github.com/elastic/beats/libbeat/cmd"
import "github.com/kriegerse/cmkbeat/beater"

// Name of this beat
var Name = "cmkbeat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmd(Name, "", beater.New)

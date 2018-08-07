package main

import (
        "os"
        "github.com/kriegerse/cmkbeat/cmd"
)

func main() {
//	beat.Run("cmkbeat", "", beater.New)
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Tamnme/tasklist-cli/cmd"
	"github.com/Tamnme/tasklist-cli/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "task.db")
	handle(db.Init(dbPath))
	handle(cmd.RootCmd.Execute())
}

func handle(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

package main

import "github.com/AHAOAHA/Annal/binaries/internal/todo"

func init() {
	rootCmd.AddCommand(todo.Cmd)
}

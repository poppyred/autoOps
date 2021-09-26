package main

import "autoOps/cmd"

//go:generate echo go:generate staring
//go:generate make build_web
//go:generate echo go:generate end
func main() {   cmd.Execute() }
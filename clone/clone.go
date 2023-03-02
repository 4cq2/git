package main

import (
   "github.com/go-git/go-git/v5"
   "os"
)

func main() {
   is_bare := false
   git.PlainClone("filter", is_bare, &git.CloneOptions{
      Progress: os.Stdout,
      URL: "https://github.com/robpike/filter",
   })
}

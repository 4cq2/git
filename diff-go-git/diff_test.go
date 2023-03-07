package diff

import (
   "fmt"
   "github.com/go-git/go-git/v5"
   "testing"
)

var (
   _ diff.Chunk = chunk{}
   _ diff.File = file{}
   _ diff.FilePatch = file_patch{}
   _ diff.Patch = patch{}
)

func Test_Diff(t *testing.T) {
   repo, err := git.PlainOpen(`D:\GitHub\git`)
   if err != nil {
      t.Fatal(err)
   }
   tree, err := repo.Worktree()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(tree)
}

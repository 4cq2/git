package git

import (
   "fmt"
   "testing"
)

func Test_Diff(t *testing.T) {
   repo, err := PlainOpen(`D:\GitHub\git`)
   if err != nil {
      t.Fatal(err)
   }
   tree, err := repo.Worktree()
   if err != nil {
      t.Fatal(err)
   }
   change, err := tree.diffStagingWithWorktree(false)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(change)
}

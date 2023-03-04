package git

import (
   "fmt"
   "testing"
)

/*
utils\diff\diff.go

FIND THIS
20:func Do(src, dst string) (diffs []diffmatchpatch.Diff) {

42:func Dst(diffs []diffmatchpatch.Diff) string {
53:func Src(diffs []diffmatchpatch.Diff) string {
*/

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

package diff

import (
   "bytes"
   "errors"
   "fmt"
   "github.com/go-git/go-git/v5"
   "github.com/go-git/go-git/v5/plumbing"
   "github.com/go-git/go-git/v5/plumbing/filemode"
   "github.com/go-git/go-git/v5/plumbing/object"
   "github.com/go-git/go-git/v5/utils/merkletrie"
   "github.com/go-git/go-git/v5/utils/merkletrie/filesystem"
   "github.com/go-git/go-git/v5/utils/merkletrie/index"
   "github.com/go-git/go-git/v5/utils/merkletrie/noder"
)

func transformChildren(t *Tree) ([]noder.Noder, error) {
   var err error
   var e TreeEntry
   ret := make([]noder.Noder, 0, len(t.Entries))
   walker := NewTreeWalker(t, false, nil) // don't recurse
   for {
      _, e, err = walker.Next()
      if err == io.EOF {
         break
      }
      if err != nil {
         walker.Close()
         return nil, err
      }
      ret = append(ret, &treeNoder{
         parent: t,
         name:   e.Name,
         mode:   e.Mode,
         hash:   e.Hash,
      })
   }
   walker.Close()
   return ret, nil
}

func (t *treeNoder) Children() ([]noder.Noder, error) {
   if t.mode != filemode.Dir {
      return noder.NoChildren, nil
   }
   if t.children != nil {
      return t.children, nil
   }
   parent := t.parent
   if !t.isRoot() {
      var err error
      if parent, err = t.parent.Tree(t.name); err != nil {
         return nil, err
      }
   }
   return transformChildren(parent)
}

type treeNoder struct {
   parent   *object.Tree  // the root node is its own parent
   name     string // empty string for the root node
   mode     filemode.FileMode
   hash     plumbing.Hash
   children []noder.Noder // memoized
}

func newChangeEntry(p noder.Path) (object.ChangeEntry, error) {
   if p == nil {
      return empty, nil
   }
   asTreeNoder, ok := p.Last().(*treeNoder)
   if !ok {
      return object.ChangeEntry{}, errors.New("cannot transform non-TreeNoders")
   }
   return object.ChangeEntry{
      Name: p.String(),
      Tree: asTreeNoder.parent,
      TreeEntry: object.TreeEntry{
         Name: asTreeNoder.name,
         Mode: asTreeNoder.mode,
         Hash: asTreeNoder.hash,
      },
   }, nil
}

func newChange(c merkletrie.Change) (*object.Change, error) {
   ret := &object.Change{}
   var err error
   if ret.From, err = newChangeEntry(c.From); err != nil {
      return nil, fmt.Errorf("from field: %s", err)
   }
   if ret.To, err = newChangeEntry(c.To); err != nil {
      return nil, fmt.Errorf("to field: %s", err)
   }
   return ret, nil
}

func newChanges(src merkletrie.Changes) (object.Changes, error) {
   ret := make(object.Changes, len(src))
   var err error
   for i, e := range src {
      ret[i], err = newChange(e)
      if err != nil {
         return nil, fmt.Errorf("change #%d: %s", i, err)
      }
   }
   return ret, nil
}

var emptyNoderHash = make([]byte, 24)

func diffTreeIsEquals(a, b noder.Hasher) bool {
   hashA := a.Hash()
   hashB := b.Hash()
   if bytes.Equal(hashA, emptyNoderHash) || bytes.Equal(hashB, emptyNoderHash) {
      return false
   }
   return bytes.Equal(hashA, hashB)
}

func diffStagingWithWorktree(r *git.Repository) (merkletrie.Changes, error) {
   w, err := r.Worktree()
   if err != nil {
      return nil, err
   }
   idx, err := r.Storer.Index()
   if err != nil {
      return nil, err
   }
   from := index.NewRootNode(idx)
   to := filesystem.NewRootNode(w.Filesystem, nil)
   return merkletrie.DiffTree(from, to, diffTreeIsEquals)
}

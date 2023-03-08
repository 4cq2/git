package diff

import (
   "github.com/go-git/go-git/v5/plumbing"
   "github.com/go-git/go-git/v5/plumbing/filemode"
   "github.com/go-git/go-git/v5/plumbing/format/diff"
)

func (w *Worktree) diffStagingWithWorktree(reverse bool) (merkletrie.Changes, error) {
	idx, err := w.r.Storer.Index()
	if err != nil {
		return nil, err
	}

	from := mindex.NewRootNode(idx)
	submodules, err := w.getSubmodulesStatus()
	if err != nil {
		return nil, err
	}

	to := filesystem.NewRootNode(w.Filesystem, submodules)

	var c merkletrie.Changes
	if reverse {
		c, err = merkletrie.DiffTree(to, from, diffTreeIsEquals)
	} else {
		c, err = merkletrie.DiffTree(from, to, diffTreeIsEquals)
	}

	if err != nil {
		return nil, err
	}

	return w.excludeIgnoredChanges(c), nil
}

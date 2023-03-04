# diff-go-git

Diff Staging with Worktree

I want to Diff Staging with Worktree. I found this:

https://github.com/go-git/go-git/blob/d525a514057f97bc2b183e2c67f542dd6f0ac0aa/worktree_status.go#L116

but it returns a [merkletrie.Changes][1]. To make use of that, it would need to
be converted to [object.Changes][2]. I found function `newChanges`:

https://github.com/go-git/go-git/blob/d525a514057f97bc2b183e2c67f542dd6f0ac0aa/plumbing/object/change_adaptor.go#L50

but it is not public. I tried just copying the function into my code, but that
does work as it calls another private function `newChange`:

https://github.com/go-git/go-git/blob/d525a514057f97bc2b183e2c67f542dd6f0ac0aa/plumbing/object/change_adaptor.go#L14

which calls another private function `newChangeEntry`:

https://github.com/go-git/go-git/blob/d525a514057f97bc2b183e2c67f542dd6f0ac0aa/plumbing/object/change_adaptor.go#L29

which calls a private struct `treeNoder`:

https://github.com/go-git/go-git/blob/d525a514057f97bc2b183e2c67f542dd6f0ac0aa/plumbing/object/treenoder.go#L19

after that I gave up.

[1]://pkg.go.dev/github.com/go-git/go-git/v5/utils/merkletrie#Changes
[2]://pkg.go.dev/github.com/go-git/go-git/v5/plumbing/object#Changes

https://github.com/go-git/go-git/issues/700

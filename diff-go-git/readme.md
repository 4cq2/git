# diff-go-git

https://github.com/go-git/go-git/issues/700

looks like what I want is the `Encode` method:

https://godocs.io/github.com/go-git/go-git/v5/plumbing/format/diff#UnifiedEncoder.Encode

which requires the `Patch` interface:

https://godocs.io/github.com/go-git/go-git/v5/plumbing/format/diff#Patch

which is implemented by `object.Patch`:

https://godocs.io/github.com/go-git/go-git/v5/plumbing/object#Patch

assuming the following functions were exported, we could call
`diffStagingWithWorktree`:

<https://github.com/go-git/go-git/blob/7e345bb5/worktree_status.go#L116>

then call `newChanges`:

<https://github.com/go-git/go-git/blob/7e345bb5/plumbing/object/change_adaptor.go#L50>

then call `Patch`:

https://github.com/go-git/go-git/blob/7e345bb5/plumbing/object/change.go#L149

until those functions are exported, the `Patch` interface has to be manually
implemented. I found some such code here:

https://github.com/fhs/gig/blob/dd59dc92/cli/diff.go#L181-L187

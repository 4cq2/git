# diff-go-git

## 2023-03-07

no one implements this:

~~~
type Chunk interface {
    Content() string
    Type() Operation
}
~~~

no one implements this:

~~~
type File interface {
    Hash() plumbing.Hash
    Mode() filemode.FileMode
    Path() string
}
~~~

no one implements this:

~~~
type FilePatch interface {
    IsBinary() bool
    Files() (from, to File)
    Chunks() []Chunk
}
~~~

this:

~~~
type Patch interface {
    FilePatches() []FilePatch
    Message() string
}
~~~

is implemented by:

https://godocs.io/github.com/go-git/go-git/v5/plumbing/object#Patch

these all work with Trees, so cannot be used with the index:

~~~
func (c *Change) Patch() (*Patch, error)
func (c Changes) Patch() (*Patch, error)
func (t *Tree) Patch(to *Tree) (*Patch, error)
~~~

not dealing with a commit, so this fails:

~~~
func (c *Commit) Patch(to *Commit) (*Patch, error)
~~~

we can get the index contents like this:

https://godocs.io/github.com/go-git/go-git/v5/plumbing/object#File.Contents

we can get a File like this:

https://pkg.go.dev/github.com/go-git/go-git/v5/plumbing/object#NewFile

## 2023-03-05

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

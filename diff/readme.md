# diff

https://github.com/sergi/go-diff

## driusan/dgit

https://github.com/driusan/dgit/issues/286

module is not pure Go

the current description says:

> A Pure Go Git Implementation 

this is not true. If I run a diff, everything looks normal:

~~~
> dgit diff
diff --git a/README.md b/README.md
--- a/filter/README.md
+++ b/filter/README.md
@@ -4,3 +4,5 @@

 You shouldn't use it either.

+asdf
~~~

but then I discovered that it only works because its calling `diff` from my
path. this is evidenced by it breaking when you clear the path:

~~~
> $env:path = ''
> dgit diff
diff --git a/README.md b/README.md
~~~

so not only is it not pure Go, but its calling an external executable. and to
top it off, it fails without error.

https://github.com/driusan/dgit/blob/f39f0c15/git/hashdiff.go#L81

## fhs/gig

color diff

Does Gig have an option for color?

https://github.com/fhs/gig

## go-git/go-git

https://github.com/go-git/go-git/issues/700

Diff Staging with Worktree

I want to Diff Staging with Worktree. I found `diffStagingWithWorktree`:

<https://github.com/go-git/go-git/blob/d525a514/worktree_status.go#L116>

but it returns a [merkletrie.Changes][1]. To make use of that, it would need to
be converted to [object.Changes][2]. I found function `newChanges`:

<https://github.com/go-git/go-git/blob/d525a514/plumbing/object/change_adaptor.go#L50>

but it is not public. I tried just copying the function into my code, but that
does work as it calls another private function `newChange`:

<https://github.com/go-git/go-git/blob/d525a514/plumbing/object/change_adaptor.go#L14>

which calls another private function `newChangeEntry`:

<https://github.com/go-git/go-git/blob/d525a514/plumbing/object/change_adaptor.go#L29>

which calls a private struct `treeNoder`:

https://github.com/go-git/go-git/blob/d525a514/plumbing/object/treenoder.go#L19

after that I gave up.

[1]://pkg.go.dev/github.com/go-git/go-git/v5/utils/merkletrie#Changes
[2]://pkg.go.dev/github.com/go-git/go-git/v5/plumbing/object#Changes

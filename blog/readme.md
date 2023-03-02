# Blog

## Byron/gitoxide

unrecognized subcommand 'clone'

the README:

https://github.com/Byron/gitoxide/blob/main/README.md

shows this command:

~~~
gix clone https://github.com/Byron/gitoxide gitoxide-clones-itself
~~~

but if I try that command, it fails:

~~~
error: unrecognized subcommand 'clone'

Usage: gix.exe [OPTIONS] <COMMAND>

For more information, try '--help'.
~~~

I built `gix` just now using the `main` branch.

https://github.com/Byron/gitoxide/issues/745

## Nivl/git-go

git clone

is it supported?

https://github.com/Nivl/git-go/issues/203

## amirkhaki/kit

git clone

is it supported?

https://github.com/amirkhaki/kit/issues/1

## blurfx/minigit

git clone

is it supported?

https://github.com/blurfx/minigit/issues/1

## chrisdickinson/git-rs

git clone

Is it possible to `git clone` with this package?

https://github.com/chrisdickinson/git-rs/issues/12

## driusan/dgit

color diff

If I do `dgit diff`, it outputs in black and white. If I do `git diff`, it
outputs in color. Does dGit have an option for color?

https://github.com/driusan/dgit/issues/284

## fhs/gig

color diff

Does Gig have an option for color?

https://github.com/fhs/gig/issues/31

## git-for-windows/git

git push fails without credential.helper

If I download this:

https://github.com/git-for-windows/git/releases/download/v2.39.2.windows.1/MinGit-2.39.2-busybox-64-bit.zip

and remove these from `etc/gitconfig`:

~~~init
[credential]
   helper = manager
~~~

any `git push` after that fails:

~~~
bash: can't create /dev/tty: nonexistent directory
error: failed to execute prompt script (exit code 1)
fatal: could not read Username for 'https://github.com': No such file or
directory
~~~

according to the docs, Git should have a fallback when a credential helper is
not defined:

<https://git-scm.com/docs/gitcredentials#_requesting_credentials>

https://github.com/git-for-windows/git/issues/4301

## go-git/go-git

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

https://github.com/go-git/go-git/issues/686

## gogs/git-module

clarify project purpose

I think it should be more clear that this package is just a wrapper for the
command line Git tool:

https://github.com/gogs/git-module/blob/7f9509a611f7bd3bba64aacbb78b9c00137bb144/repo.go#L127-L142

the current description "Git access through shell commands", someone could
think that the "shell commands" are provided by this module itself

https://github.com/gogs/git-module/issues/91

## izhujiang/gogit

finish remote functions

https://github.com/izhujiang/gogit/blob/f0cb8bc16f433481f54a382dd1505ad44d6571b5/api/remote.go#L5-L20

https://github.com/izhujiang/gogit/issues/1

## jelmer/dulwich

KeyError: b'HEAD'

I can create a new repo with Git, `git add` a change, then add another change,
and `git diff` the index with worktree:

~~~
> git init
> 'hello world' > hello.txt
> git add hello.txt
> 'world hello' >> hello.txt
> git diff
diff --git a/hello.txt b/hello.txt
index f35d3e6..d87f0a5 100644
--- a/hello.txt
+++ b/hello.txt
@@ -1 +1,2 @@
 hello world
+world hello
~~~

but Dulwich fails on the diff:

~~~
> dulwich diff
Traceback (most recent call last):
  File "runpy.py", line 196, in _run_module_as_main
  File "runpy.py", line 86, in _run_code
  File "C:\python\Scripts\dulwich.exe\__main__.py", line 7, in <module>
  File "C:\python\lib\site-packages\dulwich\cli.py", line 811, in main
    return cmd_kls().run(argv[1:])
  File "C:\python\lib\site-packages\dulwich\cli.py", line 180, in run
    commit = parse_commit(r, commit_id)
  File "C:\python\lib\site-packages\dulwich\objectspec.py", line 239, in parse_commit
    raise KeyError(committish)
KeyError: b'HEAD'
~~~

https://github.com/jelmer/dulwich/issues/1147

## samrat/rug

archive repo

can you please archive repo, to signify it is no longer being worked on

https://github.com/samrat/rug/issues/5

## ssrathi/gogit

git clone

is it supported?

https://github.com/ssrathi/gogit/issues/3

## sunshine69/gogit

git diff not working

no output:

~~~
> gitg status
Modified files not staged
M readme.md
> gitg diff
~~~

https://github.com/sunshine69/gogit/issues/1

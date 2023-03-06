# Blog

~~~
git client language:go NOT github

git client pushed:>2022-03-05 stars:>0 NOT github
~~~

## DonutLaser/gitgud

git clone

is it supported?

https://github.com/DonutLaser/gitgud/issues/1

## Nivl/git-go

no git clone

https://github.com/Nivl/git-go

## OLUWAMUYIWA/got

Windows

If I try to build, I get this:

~~~
got-main\cli> go build
go: downloading github.com/hexops/gotextdiff v1.0.3
go: downloading golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e
# github.com/OLUWAMUYIWA/got/pkg
..\pkg\index.go:153:17: undefined: unix.Stat_t
..\pkg\index.go:154:13: undefined: unix.Stat
..\pkg\index.go:169:17: undefined: unix.Stat_t
..\pkg\index.go:170:13: undefined: unix.Stat
..\pkg\index.go:177:32: undefined: unix.Stat_t
~~~

https://github.com/OLUWAMUYIWA/got/issues/1

## amirkhaki/kit

no git clone

https://github.com/amirkhaki/kit

## as/git

git clone

is it supported?

https://github.com/as/git/issues/1

## blurfx/minigit

git clone

is it supported?

https://github.com/blurfx/minigit/issues/2

## chrisdickinson/git-rs

git clone

Is it possible to `git clone` with this package?

https://github.com/chrisdickinson/git-rs/issues/13

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

https://github.com/git-for-windows/git

## gogs/git-module

just a wrapper for the command line Git tool:

https://github.com/gogs/git-module/blob/7f9509a6/repo.go#L127-L142

## izhujiang/gogit

finish remote functions

https://github.com/izhujiang/gogit/blob/f0cb8bc1/api/remote.go#L5-L20

https://github.com/izhujiang/gogit/issues/2

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

https://github.com/jelmer/dulwich/issues/1153

## samrat/rug

archive repo

can you please archive repo, to signify it is no longer being worked on

https://github.com/samrat/rug/issues/6

## sba1/simplegit

release

It would be helpful if builds were available for a couple of platforms

https://github.com/sba1/simplegit/issues/52

## ssrathi/gogit

git clone

is it supported?

https://github.com/ssrathi/gogit/issues/4

## sunshine69/gogit

git diff not working

no output:

~~~
> gitg status
Modified files not staged
M readme.md
> gitg diff
~~~

https://github.com/sunshine69/gogit

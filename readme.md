# Git

Git implementation

1. [x] git clone
2. [ ] git diff
3. [ ] git add
4. [ ] git branch
5. [ ] git checkout
6. [ ] git clean
7. [ ] git commit
8. [ ] git config
9. [ ] git log
10. [ ] git push
11. [ ] git reset
12. [ ] git status
13. [ ] git tag

# Git diff

- https://github.com/fhs/gig/issues/31
- https://github.com/sergi/go-diff

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

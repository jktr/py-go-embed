# py-go-embed

This is purely an an experiment.

I wanted to try out embedding a python interpreter in a go project using CGo.
Turns out: it works.

```
$ git clone github.com/jktr/py-go-embed
$ cd py-go-embed
$ nix run
> :l main.py
hello from python
> foo = "bar"
> print(foo)
bar
> :l main.py
hello from python
foo is bar
> ^D
```

# gvd is vim dictionary completion for golang

For example, we want to create a vim dictionary only for 
Golang standard library and store dictionary in file `$HOME/gostd.dict`

Add next line in `vimrc`:

```vim
:set iskeyword+=.                 " add point for dictionary work
:set dictionary+=$HOME/gostd.dict " location of dictionary file
```

Create a dictionary by running:

```
$ go run main.go > $HOME/gostd.dict
```

```
go run main.go > ~/dotfiles/gostd.txt
```

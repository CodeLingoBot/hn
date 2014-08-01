##About

A Hacker News ncurses CLI reader written in Go. Currently known to work on Linux, and a few people have gotten it to work on OSX.

Right now it's able to view articles, view comments, and open a page in your default browser, all done directly from the site using goquery (jquery-like library for Go), goncurses, and xdg-open for opening pages.

![Story view](https://raw.github.com/andrewstuart/hn/master/readme/stories.png)

![Comment view](https://raw.github.com/andrewstuart/hn/master/readme/comments.png)

##Installation

Assuming you have your GOPATH and PATH set appropriately:

Unfortunately, you'll also need mercurial installed, for packages hosted at code.google.com.  If you're working with Go, you've probably already done that.

```bash
go get github.com/andrewstuart/hn
```

##Usage

```bash
$ hn
```

###Story view
- n) Go to next page
- p) Go to previous page
- (num)c) View comments for story (num)
- (num)o) Open story (num) in default browser
- q) Quit hn

###Comments view
- d) Go down 30 lines
- u) Go up 30 lines
- j) Go down 1 line
- k) Go up 1 line
- n) Go down 1 page
- p) Go up 1 page
- q) Go back to story view

##API (unfinished/deprecated/idk)

This basically only works for page 1 in its current state, IIRC.

```bash
$ hn -s -p 3000 & 

$ curl localhost:3000
```

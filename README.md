# git-generate

A tool to generate random commits in a git repository.

## Installation

```bash
> go get github.com/lukahartwig/git-generate
```

## Usage

```bash
> git init
> git-generate -c 5
> git log --oneline
```

Will result in the following output:

```txt
e0063c2 (HEAD -> master) Your commit is writing checks your merge can't cash.
9693c03 Push poorly written test can down the road another ten years
35a8d49 bumping poms
fcbf894 changed things...
cbba695 Well the book was obviously wrong.
```

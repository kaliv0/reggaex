<p align="center">
  <img src="https://github.com/kaliv0/reggaex/blob/main/gopher-rasta.png?raw=true" width="400" alt="Reggaex">
</p>

# Reggaex
## buggy regex engine in Go

![Golang 1.2x](https://img.shields.io/badge/go-1.23-blue?style=flat-square&logo=Go&logoColor=blue)
[![tests](https://img.shields.io/github/actions/workflow/status/kaliv0/reggaex/ci.yml)](https://github.com/kaliv0/reggaex/actions/workflows/ci.yml)

## Example

```go
expr := `^https://\w+@[a-z0-9]+.(com|net|org)$`
str := `https://qwerty123@heythere42.com`
res, err := rgx.Match(expr, str)
if err != nil {
    fmt.Println(err)
}
fmt.Println(res.Matched)
fmt.Println(res.MatchStr)
```
output:
```console
$ true 
$ https://qwerty123@heythere42.com
```

<br>WIth invalid input:
```go
expr := `^\d{x}$`
str := `1234`
_, err := rgx.Match(expr, str)
if err != nil {
    fmt.Print(err)
}
```
```console
$ supplied value 'x' is not a number
```
# Data parser for *not-so-popular* anime game

## Status
Out of total 394 files:
- 344 files have 1 correct parser
- 3 files have multiple parsers (have fun guessing which is correct)
- 39 files have no match
- 8 unity containers

Was tested with Global version v6.3

## Using
Tool has 2 main features: 
- decoding binary files (if you want to explore them yourself), and
- parsing binary files into structured json.
> Note that you don't *need* to decode files - it will be done automatically during parsing process.

To simplify usage, 2 folders are used:
- `testdata` as input
- `result` as output

Put your files into `testdata` folder, run app and pick "Parse all". Then just check `result` folder for output.

You can also pass absolute path to file if you really need to.

## How to obtain
Check last release or github actions for precompiled executable file

## Building manually
You need to have installed GO (install from [go.dev](https://go.dev)) version 1.19 or higher

Clone repo:
```sh
git clone <url>
```
and open it in terminal:
```sh
cd <path-to-cloned-repo>
```
While inside folder, run 
```sh
go build ./cmd/app
```
Lauch created `app.exe`

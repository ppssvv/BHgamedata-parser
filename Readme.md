# Data parser for *not-so-popular* anime game

> WIP. Not all files are implemented yet

## Using
1. Download latest release
1. Run app.exe
1. Follow instructions

When parsing one file - you can use absolute path or relative to `/templates` folder.
If you want to parse several files together - put them inside `/template`. For example, `/templates/...`

## Building manually
> You need to have installed GO (get from [go.dev](https://go.dev)) version 1.19 or higher

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

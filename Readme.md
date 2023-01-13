# Data parser for *not-so-popular* anime game

> WIP. Not all files are implemented yet

## Using
1. Download latest release
1. Run app.exe
1. Follow instructions

When parsing one file - you can use absolute path or relative to `/templates` folder.
If you want to parse several files together - put them in respective folder inside `/template`. For example, `/templates/DialogueData/...`

## Building manually
> You need to have installed GO [go.dev](https://go.dev) version 1.19 or higher

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

## Updating kaitai structures
> If you change any .ksy file - you'll have to run 
```sh
go run ./cmd/genKaitai/
``` 
to re-generate code and then re-build application, following steps from [Building manually](#Building_Manually)
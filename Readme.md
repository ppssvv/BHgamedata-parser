# Data parser for *not-so-popular* anime game

## Status
For the list of unsupported files check [Missing.md](Missing.md).
Files with non-number name (e.g. "settings_*.unity3d") are not excel files, use Studio for them.

## Versioning
Starting from v6.6.0 version will correspond to the game version. 
Keep in mind that format of files changes often, so newer version of tool might not work with files from older version. 

## Using
Just lauch `app.exe` and it will ask you to provide files.

Alternatively, you can use CLI mode
```
app.exe <input_path> <output_path> [decode-only]

Required parameters:
    <input_path>
        file or folder with files to parse
    <output_path>
        output folder where to store results

Optional parameters:
    [decode-only]
        instead of parsing will just decode files (if you want to investigate)
        **Make sure to set different _output_ folder to avoid overriding your data as it will use same filenames!**

    
```

## How to obtain
Check [releases](https://github.com/funplay133/gamedata-parser/releases/latest) 
or get latest build from [github actions](https://github.com/funplay133/gamedata-parser/actions/workflows/build.yml)

By the way, your antivirus might treat tool as a virus, but its a [common issue](https://go.dev/doc/faq#virus) with Go binaries 

## Building manually
You need to have installed GO (install from [go.dev](https://go.dev)) version 1.19 or higher

In your terminal run:
```sh
git clone https://github.com/funplay133/gamedata-parser.git
cd gamedata-parser
go build ./cmd/app
```

Now you can just launch created `app.exe` or use it from command line

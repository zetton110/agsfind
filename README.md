# agsf

## Description
This command line tool is used to find for information on song and programs used in anime, games, and special effects(SFX).

## Usage

### Find song and program information

Find the song title (of anime, game and SFX) according to conditions.

#### COMMANDS:

#### `agsf s,song [-n <SONG-NAME>][-x <PROGRAM-NAME>][-s <SINGER-NAME>]`
Find the song title (of anime, game and SFX) according to conditions.
This command allows you to search by song title for information on the programs in which the song is used.

| OPTION | DESCRIPTION |
| ---- | ---- |
| `-n, --name` | Find information about songs by part of its name. (default) |
| `-x, --xlookup-by-program` | Find information about theme song by part of the program name. |
| `-s, --singer` | Find songs by singer name. |

#### `agsf p,program [-n <PROGRAM-NAME>][-x <SONG-NAME>][-s <SINGER-NAME>]`
Find the program title (of anime, game and SFX) according to conditions.
This command allows you to search for information on a program's theme song by its name.

| OPTION | DESCRIPTION |
| ---- | ---- |
| `-n, --name` | Find information about programs by part of those name. (default) |
| `-x, --xlookup-by-theme-song` | Find information about the programs whose song name is the theme song. |
| `-s, --singer-of-theme-song` | Find programs by artist name. |

#### Common Options

| OPTION | DESCRIPTION |
| ---- | ---- |
| `-v, --verbose` | Find information about the programs with details. |
| `-w, --word-regexp` | Find information about programs only if they exactly match the search word. |
| `-o, --output-with-format <FORMAT>` | Specify output format. `-o CSV` `-o JSON` |
| `-c, --category <CATEGORY>` | Specify the category to find. `-c anime` `-c game` `-c sf` |
| `-from <YYYY-MM-DD>` <br> `-to <YYYY-MM-DD>` | Specify the period of program started to search. `-from 2024-07-01 -to 2024-09-30` |

### Manipulate local database

Create and Update local database.

#### COMMANDS:

#### `agsf updatedb`
Update local db to latest.

## How to develop

### Build app

```
make build
```
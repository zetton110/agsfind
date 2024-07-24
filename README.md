# agsfind

## Description
This command line tool is used to find for information on song and programs used in anime, games, and special effects(SFX).

## Usage

### Find song and program information

Find the song title (of anime, game and SFX) according to conditions.

#### COMMANDS:

#### `agsfind s,song [-t TITLE][-p PROGRAM-TITLE][-a ARTIST-NAME]`
Find the song title (of anime, game and SFX) according to conditions.
This command allows you to search by song title for information on the programs in which the song is used.

| OPTION | DESCRIPTION |
| ---- | ---- |
| `-t, --title` | Find information about songs by part of its title. (default) |
| `-p, --program-title` | Find information about theme song by part of the program title. |
| `-a, --artist` | Find songs by artist name. |

#### `agsfind p,program [-t TITLE][-s SONG-TITLE][-a ARTIST-NAME]`
Find the program title (of anime, game and SFX) according to conditions.
This command allows you to search for information on a program's theme song by its name.

| OPTION | DESCRIPTION |
| ---- | ---- |
| `-t, --title` | Find information about programs by part of those title. (default) |
| `-s, --find-by-song-title` | Find information about the programs whose song title is the theme song. |
| `-a, --artist` | Find programs by artist name. |

#### Common Options

| OPTION | DESCRIPTION |
| ---- | ---- |
| `-v, --verbose` | Find information about the programs with details. |
| `-w, --word-regexp` | Find information about programs only if they exactly match the search word. |
| `-o, --output-format` FORMAT | Specify output format. `-o CSV` `-o JSON` |
| `-c, --category` CATEGORY| Specify the category to find. `-c anime` `-c game` `-c sf` |
| `-from` YYYY-MM-DD <br> `-to` YYYY-MM-DD | Specify the period of program started to search. `-from 2024-07-01 -to 2024-09-30` |

### Manipulate local database

Create and Update local database.

#### COMMANDS:

#### `agsfind updatedb`
Update local db to latest.

## How to develop

### Build app

```
make build
```
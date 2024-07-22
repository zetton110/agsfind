# agsfind

## Description
This command line tool is used to find for information on song and programs used in anime, games, and special effects(SFX).

## Usage

### Find song and program information

Find the song title (of anime, game and SFX) according to conditions.

#### COMMANDS:

#### `agsfind song [-t TITLE][-p PROGRAM-TITLE]`
Find the song title (of anime, game and SFX) according to conditions.
This command allows you to search by song title for information on the programs in which the song is used.

| OPTION | DESCRIPTION |
| ---- | ---- |
| `-t, --find-by-song-title` | Find information about songs by part of its title. (default) |
| `-p, --find-by-program-title` | Find information about theme song by part of the program title. |

#### `agsfind prg [-t TITLE][-s SONG-TITLE]`
Find the program title (of anime, game and SFX) according to conditions.
This command allows you to search for information on a program's theme song by its name.

| OPTION | DESCRIPTION |
| ---- | ---- |
| `-t, --find-by-program-title` | Find information about programs by part of those title. (default) |
| `-p, --find-by-song-title` | Find information about the programs whose song title is the theme song. |

#### Common Options

| OPTION | DESCRIPTION |
| ---- | ---- |
| `-v, --verbose` | Find information about the programs with details. |
| `-w, --word-regexp` | Find information about programs only if they exactly match the search word. |
| `-o, --output-format` FORMAT | Specify output format. `-o CSV` `-o JSON` |
| `-c, --category` CATEGORY| Specify the category to find. `-c anime` `-c game` `-c sf` |
| `-from` YYYY-MM-DD <br> `-to` YYYY-MM-DD | Specify the period to search. `-from 2024-07-01 -to 2024-09-30` |

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
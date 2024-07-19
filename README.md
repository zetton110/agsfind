# cmkish-cli

## Description
This command line tool is used to find for information on song and programs used in anime, games, and special effects(SFX).

## Usage

### Find song and program information

Find the song title (of anime, game and SFX) according to conditions.

#### COMMANDS:

#### `cmkish song [-t TITLE][-p PROGRAM-TITLE]`
Find the song title (of anime, game and SFX) according to conditions.
This command allows you to search by song title for information on the programs in which the song is used.

| OPTION | DESCRIPTION |
| ---- | ---- |
| `-t, --find-by-song-title` | Find information about songs by part of its title. (default) |
| `-p, --find-by-program-title` | Find information about theme song by part of the program title. |
| `-v, --verbose` | Find information about theme song with details. |
| `-w, --word-regexp` | Find information about theme songs only if they exactly match the search word. |

#### `cmkish program [-t TITLE][-s SONG-TITLE]`
Find the program title (of anime, game and SFX) according to conditions.
This command allows you to search for information on a program's theme song by its name.

| OPTION | DESCRIPTION |
| ---- | ---- |
| `-t, --find-by-program-title` | Find information about programs by part of those title. (default) |
| `-p, --find-by-song-title` | Find information about the programs whose song title is the theme song. |
| `-v, --verbose` | Find information about the programs with details. |
| `-w, --word-regexp` | Find information about programs only if they exactly match the search word. |

### Manipulate local database

Create and Update local database.

#### COMMANDS:

#### `cmkish makedb`
Create local db.

#### `cmkish updatedb`
Update local db to latest.

## How to develop

### Build app

```
make build
```
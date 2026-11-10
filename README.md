# Scaligo
A tui guitar scale visualizer in go, using tcell.

This project is a wip and is NOT ai assisted. In other words, this is not slop, and therefor it may take a long time to develop further.


# Usage
Download a version from releases for your OS

or

clone and build the project using the `build.sh` script. 

Run the binary from your shell and press 'h' for help. 

You can also run `./scaligo -h` to learn about the flags
 
Use 's' in interactive mode to learn the shortcuts

# Flags
Use `-s` to start with a scale name or alias, such as `./scaligo -s am`. Default is A Minor Pentatonic.

Use `-l` to list scales, `-a` to list aliases, and `-h` for commandline help.

Use `-c` to customize colors and markers. Colors accept names or hex values.

- `-style` : scale notes above the fretboard
- `-scalename` : scale name
- `-stringstyle` : strings
- `-fretstyle` : frets and junctions
- `-matchstyle` : natural notes
- `-matchsharpstyle` : sharp notes
- `-inlaystyle` : inlays
- `-noterune` : natural note marker
- `-sharprune` : sharp note marker
- `-inlayrune` : inlay marker

Marker flags use the first character you pass. Use `-e` to use the built-in emoji markers, this overrides custom markers.

Examples:

Change the colors:
```sh
./scaligo -c -style "#EBDBB2" -scalename "#FABD2F" -stringstyle "#665C54" -fretstyle "#D79921" -matchstyle "#B8BB26" -matchsharpstyle "#FB4934" -inlaystyle "#D3869B"
```

Change just the markers:
```
./scaligo -c -noterune=o -sharprune=x
```

A fully customized look:
```
 ./scaligo -c -style "#EBDBB2" -scalename "#FABD2F" -stringstyle "#665C54" -fretstyle "#D79921" -matchstyle "#B8BB26" -matchsharpstyle "#FB4934" -inlaystyle "#D3869B" -noterune o -sharprune x -inlayrune + -s fsharpminor

```

# Customization
All the characters and runes used to draw the fretboard and the scales are customizable using the parameters that follow the `-c` flag.

Some examples would be:
```
./scaligo -c -style "#cdd6f4" -scalename "#f5c2e7" -fretstyle "#585b70" -stringstyle "#313244" -matchstyle "#89b4fa" -matchsharpstyle "#f5c2e7" -inlaystyle "#cba6f7" -noterune ⬤ -sharprune ⬤ -inlayrune ✦
```

Dracula — double-line board, pink/green on purple-gray:
```
./scaligo -c -style "#f8f8f2" -scalename "#ff79c6" -fretstyle "#6272a4" -stringstyle "#44475a" -matchstyle "#50fa7b" -matchsharpstyle "#ff79c6" -inlaystyle "#bd93f9" -fretrune ║ -stringrune ═ -junctionrune ╬ -trune ╦ -revtrune ╩ -noterune ● -inlayrune ◆
```

Nord — cold, quiet, everything ice-blue:
```
./scaligo -c -style "#eceff4" -scalename "#88c0d0" -fretstyle "#4c566a" -stringstyle "#434c5e" -matchstyle "#81a1c1" -matchsharpstyle "#ebcb8b" -inlaystyle "#5e81ac" -noterune ⬤ -inlayrune ◇
````

Gruvbox — warm earth tones, yellow inlays like real mother-of-toilet-seat:
```
./scaligo -c -style "#ebdbb2" -scalename "#fabd2f" -fretstyle "#504945" -stringstyle "#3c3836" -matchstyle "#b8bb26" -matchsharpstyle "#fb4934" -inlaystyle "#d79921" -noterune ● -inlayrune ◆
```

Neon cyberpunk — frets glow magenta, notes are cyan:
```
./scaligo -c -style "#e0e0e0" -scalename "#00ffd5" -fretstyle "#ff2bd6" -stringstyle "#4a4a6a" -matchstyle "#00ffd5" -matchsharpstyle "#ff2bd6" -inlaystyle "#ffe66d" -noterune ⏺ -inlayrune ◆
```

If you wanna be goofy:
```
./scaligo -c   -style "#e0e0e0" -scalename "#00ffd5" -fretstyle "#ff2bd6" -stringstyle "#4a4a6a" -matchstyle "#00ffd5" -matchsharpstyle "#ff2bd6" -e -noteemoji="😀" -sharpnoteemoji="🫠" -inlayemoji="🤣" -s fsharpminor
```
  
# The Name
Somewhat inspired by skeligrow potion in HP, since you know, we are sort of growing a skeleton for our fretboard (don't laugh)

# Looks (using uxterm)
<img width="1157" height="884" alt="image" src="https://github.com/user-attachments/assets/8ca2e212-519b-4a27-b45d-5be1ee358d3e" />
<img width="1484" height="859" alt="image" src="https://github.com/user-attachments/assets/9d9be510-9969-41c8-9095-f03ebabec54c" />
<img width="800" height="450" alt="image" src="https://github.com/user-attachments/assets/20ceff32-fe0c-4bbc-9a9b-0e3e5c2a8ae1" />

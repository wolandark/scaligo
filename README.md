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
<img width="1097" height="483" alt="image" src="https://github.com/user-attachments/assets/8e7d9efc-ce4a-4494-bf03-228ac375a3b3" />


Change just the markers:
```
./scaligo -c -noterune=o -sharprune=x
```
<img width="1097" height="483" alt="image" src="https://github.com/user-attachments/assets/6a3dba33-16cf-4380-b39d-e022300b1842" />

A fully customized look:
```
 ./scaligo -c -style "#EBDBB2" -scalename "#FABD2F" -stringstyle "#665C54" -fretstyle "#D79921" -matchstyle "#B8BB26" -matchsharpstyle "#FB4934" -inlaystyle "#D3869B" -noterune o -sharprune x -inlayrune + -s fsharpminor

```
<img width="1097" height="483" alt="image" src="https://github.com/user-attachments/assets/9e55e43a-a0bd-4dd1-9fca-a2af53a975ac" />

# Tunings
Use `-T` to get a list of available tunings. Once you pick a desired tuning, you can pass it with `-t`

Example:
```
./scaligo -t dropd -s dm
```

# Customization
All the characters and runes used to draw the fretboard and the scales are customizable using the parameters that follow the `-c` flag.

Some examples would be:
```
./scaligo -c -style "#cdd6f4" -scalename "#f5c2e7" -fretstyle "#585b70" -stringstyle "#313244" -matchstyle "#89b4fa" -matchsharpstyle "#f5c2e7" -inlaystyle "#cba6f7" -noterune ⬤ -sharprune ⬤ -inlayrune ✦
```
<img width="1097" height="483" alt="image" src="https://github.com/user-attachments/assets/87900f6b-8310-4316-81c7-b37a5fbdc33f" />

Dracula — double-line board, pink/green on purple-gray:
```
./scaligo -c -style "#f8f8f2" -scalename "#ff79c6" -fretstyle "#6272a4" -stringstyle "#44475a" -matchstyle "#50fa7b" -matchsharpstyle "#ff79c6" -inlaystyle "#bd93f9" -fretrune ║ -stringrune ═ -junctionrune ╬ -trune ╦ -revtrune ╩ -noterune ● -inlayrune ◆
```
<img width="1097" height="483" alt="image" src="https://github.com/user-attachments/assets/06b1317c-b594-4179-9fab-f436137511f0" />

Nord — cold, quiet, everything ice-blue:
```
./scaligo -c -style "#eceff4" -scalename "#88c0d0" -fretstyle "#4c566a" -stringstyle "#434c5e" -matchstyle "#81a1c1" -matchsharpstyle "#ebcb8b" -inlaystyle "#5e81ac" -noterune ⬤ -inlayrune ◇
````
<img width="1097" height="483" alt="image" src="https://github.com/user-attachments/assets/667ce5ae-4db4-4cd7-a1bf-fa491ea1f438" />


Gruvbox — warm earth tones, yellow inlays like real mother-of-toilet-seat:
```
./scaligo -c -style "#ebdbb2" -scalename "#fabd2f" -fretstyle "#504945" -stringstyle "#3c3836" -matchstyle "#b8bb26" -matchsharpstyle "#fb4934" -inlaystyle "#d79921" -noterune ● -inlayrune ◆
```
<img width="1097" height="483" alt="image" src="https://github.com/user-attachments/assets/dbcb7414-20c4-44e8-a111-cf2bb43aaf92" />

Neon cyberpunk — frets glow magenta, notes are cyan:
```
./scaligo -c -style "#e0e0e0" -scalename "#00ffd5" -fretstyle "#ff2bd6" -stringstyle "#4a4a6a" -matchstyle "#00ffd5" -matchsharpstyle "#ff2bd6" -inlaystyle "#ffe66d" -noterune ⏺ -inlayrune ◆
```
<img width="1097" height="483" alt="image" src="https://github.com/user-attachments/assets/aab83bec-d31b-4ce5-82a1-a9827c8be5ed" />

Default emojis with `-e`:
```
./scaligo -e -s em
```
<img width="1097" height="483" alt="image" src="https://github.com/user-attachments/assets/9d5490c8-c8ae-47d9-be94-b997ec051c95" />


If you wanna be goofy:
```
./scaligo -c   -style "#e0e0e0" -scalename "#00ffd5" -fretstyle "#ff2bd6" -stringstyle "#4a4a6a" -matchstyle "#00ffd5" -matchsharpstyle "#ff2bd6" -e -noteemoji="😀" -sharpnoteemoji="🫠" -inlayemoji="🤣" -s fsharpminor
```
<img width="1097" height="483" alt="image" src="https://github.com/user-attachments/assets/87182f70-7fae-435a-8525-3f3622fffc61" />

# The Name
Somewhat inspired by skeligrow potion in HP, since you know, we are sort of growing a skeleton for our fretboard (don't laugh)

# Default Look (using uxterm)
<img width="1157" height="884" alt="image" src="https://github.com/user-attachments/assets/8ca2e212-519b-4a27-b45d-5be1ee358d3e" />

Press `h` for the help overlay
<img width="1350" height="520" alt="image" src="https://github.com/user-attachments/assets/20f5d6aa-da4c-4f42-8020-c4e082b259f2" />

Press `s` to get the shortcuts overlay
<img width="1869" height="997" alt="image" src="https://github.com/user-attachments/assets/b7e4cb0d-14f2-4876-83c4-dfc46045094d" />

# That is it for now
<img width="800" height="450" alt="image" src="https://github.com/user-attachments/assets/20ceff32-fe0c-4bbc-9a9b-0e3e5c2a8ae1" />

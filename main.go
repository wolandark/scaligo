package main

import (
	"flag"
	"fmt"
	"gs/scales"
	"os"
	"slices"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func GetScale(input string) ([]string, bool) {
	key := strings.ToLower(strings.TrimSpace(input))
	if alias, ok := scales.ScaleAliases[key]; ok {
		key = alias
	} else {
		key = strings.ReplaceAll(key, " ", "_")
	}
	scale, ok := scales.Scales[key]
	return scale, ok
}

func main() {
	//======================================================================================
	// flag section
	//======================================================================================
	s_flag := flag.String("s", "", "scale")
	flag.Parse()

	provided := false

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "s" {
			provided = true
		}
	})

	if !provided {
		fmt.Println("NO SCALE PROVIDED!")
		return
	}

	scale, exists := GetScale(*s_flag)

	if !exists {
		fmt.Println("Unknown scale:", *s_flag)
		os.Exit(1)
		return
	}

	//======================================================================================
	// tcell screen
	//======================================================================================
	s, _ := tcell.NewScreen()
	s.Init()
	defer s.Fini()

	//======================================================================================
	// Vars
	//======================================================================================
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	stringStyle := tcell.StyleDefault.Foreground(tcell.ColorSilver)
	fretStyle := tcell.StyleDefault.Foreground(tcell.ColorGold)
	matchStyle := tcell.StyleDefault.Foreground(tcell.ColorDarkCyan)
	matchSharpStyle := tcell.StyleDefault.Foreground(tcell.ColorRed)
	inlayStyle := tcell.StyleDefault.Foreground(tcell.ColorWhiteSmoke)

	noteRune := '〇'
	sharpRune := 'Δ'
	inlayRune := '◆'
	fretRune := '│'
	stringRune := '─'
	junctionRune := '┼'
	tRune := '┬'
	revTRune := '┴'

	termWidth, termHeight := s.Size() // get terminal sizes from tcell not x/terminal

	fretboard_e := [][]string{
		{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
		{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
		{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
		{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
		{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
		{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
	}

	const (
		fretSpacing   = 8
		stringSpacing = 2
	)

	frets := len(fretboard_e[0]) // cause fretboard is a 2d slice already the size
	strings := len(fretboard_e)  // of a fretboard

	boardWidth := (frets - 1) * fretSpacing
	boardHeight := (strings-1)*stringSpacing + 1

	x := (termWidth - boardWidth) / 2
	y := (termHeight - boardHeight) / 2

	//======================================================================================
	// draw loops
	//======================================================================================
	// draw inlays
	//======================================================================================
	inlays := []int{1, 3, 5, 7, 9, 12}
	for fretIndex := 0; fretIndex < frets; fretIndex++ {
		if slices.Contains(inlays, fretIndex) {
			for j := 4; j < 6; j++ {
				for i := 1; i < boardHeight-1; i++ {
					// s.SetContent(x+(fretIndex-1)*fretSpacing+j, y+i, '🔷', nil, style)
					s.SetContent(x+(fretIndex-1)*fretSpacing+fretSpacing/2, y+boardHeight/2, inlayRune, nil, inlayStyle)
				}
			}
		}
	}
	//======================================================================================
	// draw frets only
	//======================================================================================
	for i := 0; i < frets; i++ {
		for j := 0; j < boardHeight; j++ {
			s.SetContent(x+i*fretSpacing, y+j, fretRune, nil, fretStyle)
		}
	}

	//======================================================================================
	// draw strings only
	//======================================================================================
	for i := 0; i < strings; i++ {
		for j := 0; j < boardWidth; j++ {
			s.SetContent(x+j, y+i*stringSpacing, stringRune, nil, stringStyle)
		}
	}

	//======================================================================================
	// junctions
	//======================================================================================
	for i := 0; i < frets; i++ {
		for j := 0; j < strings; j++ {
			nx := x + i*fretSpacing
			ny := y + j*stringSpacing

			r := junctionRune

			if j == 0 {
				r = tRune
			} else if j == strings-1 {
				r = revTRune
			}

			s.SetContent(nx, ny, r, nil, fretStyle)
		}
	}

	//======================================================================================
	// draw scale notes in abcd
	//======================================================================================
	offset := 0
	for i := 0; i < len(scale); i++ {
		for j := 0; j < len(scale[i]); j++ {
			a := rune(scale[i][j])
			s.SetContent(x+offset, y-3, a, nil, style)
			offset++
		}
		offset++
	}

	//======================================================================================
	// draw notes on the fretboard
	//======================================================================================
	for stringIndex := 0; stringIndex < len(fretboard_e); stringIndex++ {
		notes := fretboard_e[stringIndex]

		for fretIndex := 0; fretIndex < len(notes); fretIndex++ {
			note := notes[fretIndex]

			if !slices.Contains(scale, note) {
				continue
			}

			nx := x + fretIndex*fretSpacing
			ny := y + stringIndex*stringSpacing

			if fretIndex > 0 {
				nx -= fretSpacing / 2
			}

			runes := []rune(note)
			r := noteRune
			var runeColor = matchStyle

			if slices.Contains(runes, '#') {
				r = sharpRune
				runeColor = matchSharpStyle
			}
			if len(runes) > 1 {
				r = sharpRune
				runes = runes[:1]
			}

			for i := 0; i < len(runes); i++ {
				s.SetContent(nx+i, ny, runes[i], nil, runeColor)
				s.SetContent(nx+i, ny, r, nil, runeColor)
			}
		}
	}

	//======================================================================================
	// main loop || game loop
	//======================================================================================
	for {
		s.Show()
		key, ok := s.PollEvent().(*tcell.EventKey)
		if !ok {
			continue
		}
		if key.Key() == tcell.KeyEscape || key.Rune() == 'q' {
			return
		}
	}
}

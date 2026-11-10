package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"scaligo/help"
	"scaligo/keys"
	"scaligo/opts"
	"slices"
	"strings"
)

// ======================================================================================
// Constants
// ======================================================================================
const (
	fretSpacing   = 8
	stringSpacing = 2
)

// ======================================================================================
// Styles moved to opts package
// ======================================================================================
// Fretboard data
// ======================================================================================
var fretboard_e = [][]string{
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
	{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
}

func main() {
	//======================================================================================
	// flags
	//======================================================================================
	scale, title, cfg := opts.ParseOpts()

	//======================================================================================
	// tcell screen
	//======================================================================================
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := s.Init(); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Fini()

	//======================================================================================
	// main loop || game loop
	//======================================================================================
	for {
		s.Clear()

		//======================================================================================
		// Dimensions and positioning
		//======================================================================================
		termWidth, termHeight := s.Size()

		frets := len(fretboard_e[0])
		str := len(fretboard_e)

		boardWidth := (frets - 1) * fretSpacing
		boardHeight := (str-1)*stringSpacing + 1

		x := (termWidth - boardWidth) / 2
		y := (termHeight - boardHeight) / 2

		//======================================================================================
		// reject small terminal
		//======================================================================================
		if termWidth < 105 && termHeight < 21 {
			fmt.Println("term too small, go fuck yourself!")

			for i, r := range "Hello, World!" {
				s.SetContent(x, y, r, nil, cfg.Style)
				i++
			}
			s.SetContent(x, y, '!', nil, tcell.StyleDefault)

			s.Clear()
			defer s.Fini()
			return
		}

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
						// s.SetContent(x+(fretIndex-1)*fretSpacing+j, y+i, '🔷', nil, cfg.Style)
						s.SetContent(x+(fretIndex-1)*fretSpacing+fretSpacing/2, y+boardHeight/2, cfg.InlayRune, nil, cfg.InlayStyle)
					}
				}
			}
		}

		//======================================================================================
		// draw frets only
		//======================================================================================
		for i := 0; i < frets; i++ {
			for j := 0; j < boardHeight; j++ {
				s.SetContent(x+i*fretSpacing, y+j, cfg.FretRune, nil, cfg.FretStyle)
			}
		}

		//======================================================================================
		// draw strings only
		//======================================================================================
		for i := 0; i < str; i++ {
			for j := 0; j < boardWidth; j++ {
				s.SetContent(x+j, y+i*stringSpacing, cfg.StringRune, nil, cfg.StringStyle)
			}
		}

		//======================================================================================
		// junctions
		//======================================================================================
		for i := 0; i < frets; i++ {
			for j := 0; j < str; j++ {
				nx := x + i*fretSpacing
				ny := y + j*stringSpacing

				r := cfg.JunctionRune

				if j == 0 {
					r = cfg.TRune
				} else if j == str-1 {
					r = cfg.RevTRune
				}

				s.SetContent(nx, ny, r, nil, cfg.FretStyle)
			}
		}

		//======================================================================================
		// draw scale notes in abcd
		//======================================================================================
		offset := 0
		// title = strings.ReplaceAll(title, "maj", "Major")
		// title = strings.ReplaceAll(title, "min", "Minor")
		// title = strings.ReplaceAll(title, "sharp", "Sharp")
		// title = strings.ToUpper(title[:1]) + title[1:]
		title = strings.ReplaceAll(title, "_", " ")
		for i := 0; i < len(title); i++ {
			s.SetContent(x+i, y-5, rune(title[i]), nil, cfg.ScaleName)
		}

		for i := 0; i < len(scale); i++ {
			for j := 0; j < len(scale[i]); j++ {
				a := rune(scale[i][j])
				s.SetContent(x+offset, y-3, a, nil, cfg.Style)
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
				r := cfg.NoteRune
				var runeColor = cfg.MatchStyle

				if slices.Contains(runes, '#') {
					r = cfg.SharpRune
					runeColor = cfg.MatchSharpStyle
				}
				if len(runes) > 1 {
					r = cfg.SharpRune
					runes = runes[:1]
				}

				for i := 0; i < len(runes); i++ {
					s.SetContent(nx+i, ny, runes[i], nil, runeColor)
					s.SetContent(nx+i, ny, r, nil, runeColor)
				}
			}
		}

		help.ShowHelp(s, termWidth, termHeight)

		s.Show()

		// read keys
		newScale, newTitle, toggleOverlay, toggleShortcuts, running := keys.ReadKeys(s)
		if !running {
			break
		}

		if toggleOverlay {
			help.Toggle()
			toggleShortcuts = false
		}

		if toggleShortcuts {
			help.ToggleShortcuts()
			toggleOverlay = false
		}

		if newScale != nil {
			scale = newScale
			title = newTitle
		}
	}
}

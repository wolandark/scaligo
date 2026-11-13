package draw

import (
	"strings"

	"fmt"
	"github.com/gdamore/tcell/v2"
	"os"
	"scaligo/opts"
	"slices"
	"strconv"
)

const (
	fretSpacing   = 8
	stringSpacing = 2
)

func fretCount(cfg opts.Config) int { return len(cfg.Tuning[0]) }

func RejectSmallTerm(termWidth, termHeight int, s tcell.Screen) {
	if termWidth < 105 || termHeight < 21 {
		s.Fini()
		fmt.Fprintf(os.Stderr, "Terminal Too Small! Need 105x21, Got %dx%d\n", termWidth, termHeight)
		os.Exit(1)
	}
}

func DrawInlays(s tcell.Screen, x, y, boardHeight int, cfg opts.Config) {
	inlays := []int{1, 3, 5, 7, 9, 12}

	for fretIndex := 0; fretIndex < fretCount(cfg); fretIndex++ {
		if slices.Contains(inlays, fretIndex) {
			s.SetContent(x+(fretIndex-1)*fretSpacing+fretSpacing/2, y+boardHeight/2, cfg.InlayRune, nil, cfg.InlayStyle)
		}
	}
}

func DrawFrets(s tcell.Screen, x, y, boardHeight int, cfg opts.Config) {
	for i := 0; i < fretCount(cfg); i++ {
		for j := 0; j < boardHeight; j++ {
			s.SetContent(x+i*fretSpacing, y+j, cfg.FretRune, nil, cfg.FretStyle)
		}
	}
}

func DrawFretNumbers(s tcell.Screen, x, y, boardHeight int, cfg opts.Config) {
	for i := 1; i < fretCount(cfg); i++ {
		num := strconv.Itoa(i)
		for j := 0; j < len(num); j++ {
			s.SetContent(x+i*fretSpacing+j-4, y+boardHeight+1, rune(num[j]), nil, cfg.FretStyle)
		}
	}
}

func DrawStrings(s tcell.Screen, x, y, boardWidth, stringSpacing int, cfg opts.Config) {
	str := len(cfg.Tuning)
	for i := 0; i < str; i++ {
		for j := 0; j < boardWidth; j++ {
			s.SetContent(x+j, y+i*stringSpacing, cfg.StringRune, nil, cfg.StringStyle)
		}
	}
}

func DrawTabs(s tcell.Screen, x, y, boardWidth, stringSpacing int, cfg opts.Config) {
	str := len(cfg.Tuning)
	for i := 0; i < str; i++ {
		for j := 0; j < boardWidth; j++ {
			s.SetContent(x+j, y+i*stringSpacing+15, cfg.StringRune, nil, cfg.StringStyle)
		}
	}
}

func DrawJunctions(s tcell.Screen, x, y, fretSpacing, stringSpacing int, cfg opts.Config) {
	str := len(cfg.Tuning)
	for i := 0; i < fretCount(cfg); i++ {
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
}

func DrawScaleInAlphabet(s tcell.Screen, x, y int, scale []string, title string, cfg opts.Config) {
	offset := 0
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
}

func DrawNotesOnTheFretboard(s tcell.Screen, x, y int, scale []string, cfg opts.Config) {
	for stringIndex := 0; stringIndex < len(cfg.Tuning); stringIndex++ {
		notes := cfg.Tuning[stringIndex]

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
				// s.SetContent(nx+i, ny, runes[i], nil, runeColor)
				s.SetContent(nx+i, ny, r, nil, runeColor)
			}

			tab := strconv.Itoa(fretIndex)

			for i := 0; i < len(tab); i++ {
				s.SetContent(nx+i, ny+15, rune(tab[i]), nil, cfg.FretStyle)
			}
		}
	}
}

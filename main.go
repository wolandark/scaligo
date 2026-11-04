package main

import (
	"flag"
	"fmt"
	"gs/scales"
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
		return
	}

	s, _ := tcell.NewScreen()
	s.Init()
	defer s.Fini()

	style := tcell.StyleDefault.Foreground(tcell.ColorLavender)
	matchStyle := tcell.StyleDefault.Foreground(tcell.ColorGreen)

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

	// draw frets only
	for i := 0; i < frets; i++ {
		for j := 0; j < boardHeight; j++ {
			s.SetContent(x+i*fretSpacing, y+j, '│', nil, style)
		}
	}

	// draw strings only
	for i := 0; i < strings; i++ {
		for j := 0; j < boardWidth; j++ {
			s.SetContent(x+j, y+i*stringSpacing, '─', nil, style)
		}
	}

	// draw scale notes in abcd
	offset := 0
	for i := 0; i < len(scale); i++ {
		for j := 0; j < len(scale[i]); j++ {
			a := rune(scale[i][j])
			s.SetContent(x+offset, y-3, a, nil, style)
			offset++
		}
		offset++
	}

	// draw notes on the fretboard
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
			r := '●'
			if slices.Contains(runes, '#') {
				r = 'Δ'
			}
			if len(runes) > 1 {
				r = 'Δ'
				runes = runes[:1]
			}

			for i := 0; i < len(runes); i++ {
				s.SetContent(nx+i, ny, runes[i], nil, matchStyle)
				s.SetContent(nx+i, ny, r, nil, matchStyle)
			}
		}
	}

	// main loop || game loop
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

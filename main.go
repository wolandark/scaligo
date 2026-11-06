package main

import (
	"flag"
	"fmt"
	"gs/scales"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
)

func GetScale(input string) ([]string, bool, string) {
	key := strings.ToLower(strings.TrimSpace(input))
	if alias, ok := scales.ScaleAliases[key]; ok {
		key = alias
	} else {
		key = strings.ReplaceAll(key, " ", "_")
	}
	scale, ok := scales.Scales[key]
	title := key
	return scale, ok, title
}

// ======================================================================================
// Constants
// ======================================================================================
const (
	fretSpacing   = 8
	stringSpacing = 2

	noteRune     = '⬤'
	sharpRune    = '⬤'
	inlayRune    = '◆'
	fretRune     = '│'
	stringRune   = '─'
	junctionRune = '┼'
	tRune        = '┬'
	revTRune     = '┴'
)

// ======================================================================================
// Styles
// ======================================================================================
var (
	style           = tcell.StyleDefault.Foreground(tcell.ColorWhite)
	stringStyle     = tcell.StyleDefault.Foreground(tcell.ColorSilver)
	fretStyle       = tcell.StyleDefault.Foreground(tcell.ColorGold)
	matchStyle      = tcell.StyleDefault.Foreground(tcell.ColorDarkCyan)
	matchSharpStyle = tcell.StyleDefault.Foreground(tcell.ColorRed)
	inlayStyle      = tcell.StyleDefault.Foreground(tcell.ColorWhiteSmoke)
)

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

var input string
var timer *time.Timer

func main() {
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

	scale, exists, title := GetScale(*s_flag)

	if !exists {
		fmt.Println("Unknown scale:", *s_flag)
		os.Exit(1)
		return
	}

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
				s.SetContent(x, y, r, nil, style)
				i++
			}
			s.SetContent(x, y, '!', nil, tcell.StyleDefault)

			os.Exit(1)
			s.Clear()
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
		for i := 0; i < str; i++ {
			for j := 0; j < boardWidth; j++ {
				s.SetContent(x+j, y+i*stringSpacing, stringRune, nil, stringStyle)
			}
		}

		//======================================================================================
		// junctions
		//======================================================================================
		for i := 0; i < frets; i++ {
			for j := 0; j < str; j++ {
				nx := x + i*fretSpacing
				ny := y + j*stringSpacing

				r := junctionRune

				if j == 0 {
					r = tRune
				} else if j == str-1 {
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
			for i := 0; i < len(title); i++ {
				s.SetContent(i, 0, rune(title[i]), nil, style)
			}
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
		s.Show()

		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Rune() == 'q' {
				return
			}

			input += string(ev.Rune())

			if timer != nil {
				timer.Stop()
			}

			timer = time.AfterFunc(500*time.Millisecond, func() {
				s.PostEvent(tcell.NewEventInterrupt(nil))
			})

		case *tcell.EventInterrupt:
			switch input {

			// C
			case "C", "cM":
				scale, _, title = GetScale("c_maj")
			case "cm":
				scale, _, title = GetScale("c_min")
			case "C#":
				scale, _, title = GetScale("c_sharp_maj")
			case "c#m":
				scale, _, title = GetScale("c_sharp_min")

			// D
			case "D", "dM":
				scale, _, title = GetScale("d_maj")
			case "dm":
				scale, _, title = GetScale("d_min")
			case "D#":
				scale, _, title = GetScale("d_sharp_maj")
			case "d#m":
				scale, _, title = GetScale("d_sharp_min")

			// E
			case "E", "eM":
				scale, _, title = GetScale("e_maj")
			case "em":
				scale, _, title = GetScale("e_min")

			// F
			case "F", "fM":
				scale, _, title = GetScale("f_maj")
			case "fm":
				scale, _, title = GetScale("f_min")
			case "F#":
				scale, _, title = GetScale("f_sharp_maj")
			case "f#m":
				scale, _, title = GetScale("f_sharp_min")

			// G
			case "G", "fM":
				scale, _, title = GetScale("g_maj")
			case "gm":
				scale, _, title = GetScale("g_min")
			case "G#":
				scale, _, title = GetScale("g_sharp_maj")
			case "g#m":
				scale, _, title = GetScale("g_sharp_min")

			// A
			case "A", "fM":
				scale, _, title = GetScale("a_maj")
			case "am":
				scale, _, title = GetScale("a_min")
			case "A#":
				scale, _, title = GetScale("a_sharp_maj")
			case "a#m":
				scale, _, title = GetScale("a_sharp_min")

			// B
			case "B", "fM":
				scale, _, title = GetScale("b_maj")
			case "bm":
				scale, _, title = GetScale("b_min")






			input = ""

		case *tcell.EventResize:
			s.Sync()
		}
	}
}

package opts

import (
	"flag"
	"fmt"
	"maps"
	"os"
	"scaligo/scales"
	"slices"

	"github.com/gdamore/tcell/v2"
)

type Config struct {
	Scale           string
	Emoji           bool
	Style           tcell.Style
	ScaleName       tcell.Style
	MatchStyle      tcell.Style
	MatchSharpStyle tcell.Style
	FretStyle       tcell.Style
	StringStyle     tcell.Style
	InlayStyle      tcell.Style
	NoteRune        rune
	SharpRune       rune
	InlayRune       rune
	FretRune        rune
	StringRune      rune
	JunctionRune    rune
	TRune           rune
	RevTRune        rune
	Tuning          [][]string
}

var fretboard_e = [][]string{
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
	{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
}

var fretboard_d = [][]string{
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
	{"F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F"},
	{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
}

var fretboard_c = [][]string{
	{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#"},
	{"A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#"},
	{"F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F"},
	{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C"},
}

var fretboard_eb = [][]string{
	{"D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#"},
	{"A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#"},
	{"F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#"},
	{"C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#"},
	{"G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"},
	{"D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#"},
}

var fretboard_drop_d = [][]string{
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
	{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
}

var fretboard_drop_c = [][]string{
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
	{"F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F"},
	{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C"},
}

var fretboard_drop_c_sharp = [][]string{
	{"D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#"},
	{"A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#"},
	{"F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#"},
	{"C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#"},
	{"G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"},
	{"C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#"},
}

var fretboard_drop_b = [][]string{
	{"C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#"},
	{"G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"},
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
	{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	{"F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#"},
	{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
}

var fretboard_dadgad = [][]string{
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
}

var fretboard_open_d = [][]string{
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
	{"F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
}

var fretboard_open_g = [][]string{
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
}

var fretboard_open_e = [][]string{
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
	{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	{"G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"},
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
	{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
}

var fretboard_open_c = [][]string{
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
	{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C"},
}

var fretboard_double_drop_d = [][]string{
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
}

func ParseOpts() ([]string, string, Config) {

	var cfg Config

	cfg.Style = tcell.StyleDefault.Foreground(tcell.ColorWhite)
	cfg.ScaleName = tcell.StyleDefault.Foreground(tcell.ColorGreen)
	cfg.StringStyle = tcell.StyleDefault.Foreground(tcell.ColorSilver)
	cfg.FretStyle = tcell.StyleDefault.Foreground(tcell.ColorLightYellow)
	cfg.MatchStyle = tcell.StyleDefault.Foreground(tcell.ColorDarkCyan)
	cfg.MatchSharpStyle = tcell.StyleDefault.Foreground(tcell.ColorRed)
	cfg.InlayStyle = tcell.StyleDefault.Foreground(tcell.ColorWhiteSmoke)
	cfg.NoteRune = '⬤'
	cfg.SharpRune = '⬤'
	cfg.InlayRune = '◆'

	cfg.FretRune = '│'
	cfg.StringRune = '─'
	cfg.JunctionRune = '┼'
	cfg.TRune = '┬'
	cfg.RevTRune = '┴'

	t_flag := flag.String("t", "predefined tunings", "tuning")
	s_flag := flag.String("s", "A_Minor_Pentatonic", "scale") // default scale if no -s
	l_flag := flag.Bool("l", false, "list known scales")      // list known scales
	a_flag := flag.Bool("a", false, "list abbreviations")     // list scale aliases
	h_flag := flag.Bool("h", false, "commandline help")       // obviously prints help
	e_flag := flag.Bool("e", false, "use emojis")
	c_flag := flag.Bool("c", false, "customize colors")

	c_style := flag.String("style", "", "Fret color (hex)")
	c_scaleName := flag.String("scalename", "", "Fret color (hex)")
	c_fretStyle := flag.String("fretstyle", "", "Fret color (hex)")
	c_stringStyle := flag.String("stringstyle", "", "String color (hex)")
	c_matchStyle := flag.String("matchstyle", "", "Match color (hex)")
	c_matchSharpStyle := flag.String("matchsharpstyle", "", "Match color (hex)")
	c_inlayStyle := flag.String("inlaystyle", "", "Inlay color (hex)")

	c_noteRune := flag.String("noterune", "", "note rune")
	c_sharpRune := flag.String("sharprune", "", "sharp note rune")
	c_inlayRune := flag.String("inlayrune", "", "inlay rune")

	c_fretRune := flag.String("fretrune", "", "rune for the fret wire")
	c_stringRune := flag.String("stringrune", "", "rune for the strings")

	c_junctionRune := flag.String("junctionrune", "", "rune for the crosspoint of strings/fret wires")
	c_tRune := flag.String("trune", "", "t rune")
	c_revTRune := flag.String("revtrune", "", "revt rune")

	c_noteEmoji := flag.String("noteemoji", "", "note emoji")
	c_sharpNoteEmoji := flag.String("sharpnoteemoji", "", "sharp note emoji")
	c_inlayEmoji := flag.String("inlayemoji", "", "inlay emoji")
	cfg.Tuning = fretboard_e

	flag.Parse()

	if *t_flag != "" {
		switch *t_flag {
		case "d":
			cfg.Tuning = fretboard_d
		case "c":
			cfg.Tuning = fretboard_c
		case "eb":
			cfg.Tuning = fretboard_eb
		case "dropd":
			cfg.Tuning = fretboard_drop_d
		case "dropc":
			cfg.Tuning = fretboard_drop_c
		case "dropc#":
			cfg.Tuning = fretboard_drop_c_sharp
		case "dropb":
			cfg.Tuning = fretboard_drop_b
		case "dadgad":
			cfg.Tuning = fretboard_dadgad
		case "opend":
			cfg.Tuning = fretboard_open_d
		case "openg":
			cfg.Tuning = fretboard_open_g
		case "opene":
			cfg.Tuning = fretboard_open_e
		case "openc":
			cfg.Tuning = fretboard_open_c
		case "doubledropd":
			cfg.Tuning = fretboard_double_drop_d
		}
	}

	if *c_flag {
		if *c_style != "" {
			cfg.Style = tcell.StyleDefault.Foreground(tcell.GetColor(*c_style))
		}
		if *c_scaleName != "" {
			cfg.ScaleName = tcell.StyleDefault.Foreground(tcell.GetColor(*c_scaleName))
		}
		if *c_matchSharpStyle != "" {
			cfg.MatchSharpStyle = tcell.StyleDefault.Foreground(tcell.GetColor(*c_matchSharpStyle))
		}
		if *c_fretStyle != "" {
			cfg.FretStyle = tcell.StyleDefault.Foreground(tcell.GetColor(*c_fretStyle))
		}
		if *c_stringStyle != "" {
			cfg.StringStyle = tcell.StyleDefault.Foreground(tcell.GetColor(*c_stringStyle))
		}
		if *c_matchStyle != "" {
			cfg.MatchStyle = tcell.StyleDefault.Foreground(tcell.GetColor(*c_matchStyle))
		}
		if *c_inlayStyle != "" {
			cfg.InlayStyle = tcell.StyleDefault.Foreground(tcell.GetColor(*c_inlayStyle))
		}
		if *c_noteRune != "" {
			cfg.NoteRune = []rune(*c_noteRune)[0]
		}
		if *c_inlayRune != "" {
			cfg.InlayRune = []rune(*c_inlayRune)[0]
		}
		if *c_sharpRune != "" {
			cfg.SharpRune = []rune(*c_sharpRune)[0]
		}
		if *c_fretRune != "" {
			cfg.FretRune = []rune(*c_fretRune)[0]
		}
		if *c_stringRune != "" {
			cfg.StringRune = []rune(*c_stringRune)[0]
		}
		if *c_junctionRune != "" {
			cfg.JunctionRune = []rune(*c_junctionRune)[0]
		}
		if *c_tRune != "" {
			cfg.TRune = []rune(*c_tRune)[0]
		}
		if *c_revTRune != "" {
			cfg.RevTRune = []rune(*c_revTRune)[0]
		}
	}

	if *e_flag {
		cfg.NoteRune, cfg.SharpRune, cfg.InlayRune = '🟢', '🔴', '💠'

		if r := []rune(*c_noteEmoji); len(r) > 0 {
			cfg.NoteRune = r[0]
		}
		if r := []rune(*c_sharpNoteEmoji); len(r) > 0 {
			cfg.SharpRune = r[0]
		}
		if r := []rune(*c_inlayEmoji); len(r) > 0 {
			cfg.InlayRune = r[0]
		}
	}

	if *h_flag {
		flag.Usage()
		os.Exit(0)
	}

	if *a_flag {
		for i, k := range slices.Sorted(maps.Keys(scales.ScaleMap)) {
			v := scales.ScaleMap[k]
			fmt.Printf("\n%d - %s -%s", i, k, v)
			if (i+1)%4 == 0 {
				fmt.Println()
			}
		}
		fmt.Println()
		os.Exit(0)
	}

	if *l_flag {
		for i, name := range slices.Sorted(maps.Keys(scales.Scales)) {
			fmt.Printf("%-35s", name)
			if (i+1)%4 == 0 {
				fmt.Println()
			}
		}
		fmt.Println()
		os.Exit(0)
	}

	scale, exists, title := scales.GetScale(*s_flag)
	if !exists {
		fmt.Println("Unknown Scale:", *s_flag)
		os.Exit(1)
	}

	cfg.Scale = *s_flag
	cfg.Emoji = *e_flag

	return scale, title, cfg
}

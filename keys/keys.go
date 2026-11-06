package keys

import (
	"github.com/gdamore/tcell/v2"
	"gs/scales"
	"time"
)

var input string
var title string
var timer *time.Timer
var scale []string

func ReadKeys(s tcell.Screen) ([]string, string, bool) {
	switch ev := s.PollEvent().(type) {
	case *tcell.EventKey:
		if ev.Rune() == 'q' {
			return nil, "", false
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
			scale, _, title = scales.GetScale("C_Major")
			input = ""
		case "cm":
			scale, _, title = scales.GetScale("C_Minor")
			input = ""
		case "C#":
			scale, _, title = scales.GetScale("C_Sharp_Major")
			input = ""
		case "c#m":
			scale, _, title = scales.GetScale("C_Sharp_Minor")
			input = ""

			// D
		case "D", "dM":
			scale, _, title = scales.GetScale("D_Major")
			input = ""
		case "dm":
			scale, _, title = scales.GetScale("D_Minor")
			input = ""
		case "D#":
			scale, _, title = scales.GetScale("D_Sharp_Major")
			input = ""
		case "d#m":
			scale, _, title = scales.GetScale("D_Sharp_Minor")
			input = ""

			// E
		case "E", "eM":
			scale, _, title = scales.GetScale("E_Major")
			input = ""
		case "em":
			scale, _, title = scales.GetScale("E_Minor")
			input = ""

			// F
		case "F", "fM":
			scale, _, title = scales.GetScale("F_Major")
			input = ""
		case "fm":
			scale, _, title = scales.GetScale("F_Minor")
			input = ""
		case "F#":
			scale, _, title = scales.GetScale("F_Sharp_Major")
			input = ""
		case "f#m":
			scale, _, title = scales.GetScale("F_Sharp_Minor")
			input = ""

			// G
		case "G", "gM":
			scale, _, title = scales.GetScale("G_Major")
			input = ""
		case "gm":
			scale, _, title = scales.GetScale("G_Minor")
			input = ""
		case "G#":
			scale, _, title = scales.GetScale("G_Sharp_Major")
			input = ""
		case "g#m":
			scale, _, title = scales.GetScale("G_Sharp_Minor")
			input = ""

			// A
		case "A", "aM":
			scale, _, title = scales.GetScale("A_Major")
			input = ""
		case "am":
			scale, _, title = scales.GetScale("A_Minor")
			input = ""
		case "A#":
			scale, _, title = scales.GetScale("A_Sharp_Major")
			input = ""
		case "a#m":
			scale, _, title = scales.GetScale("A_Sharp_Minor")
			input = ""

			// B
		case "B", "bM":
			scale, _, title = scales.GetScale("B_Major")
			input = ""
		case "bm":
			scale, _, title = scales.GetScale("B_Minor")
			input = ""
		}

		if title != "" {
			return scale, title, true
		}

	case *tcell.EventResize:
		s.Sync()
	}
	return nil, "", true
}

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

		timer = time.AfterFunc(1000*time.Millisecond, func() {
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
		case "cmp":
			scale, _, title = scales.GetScale("C_Minor_Pentatonic")
			input = ""
		case "cMp":
			scale, _, title = scales.GetScale("C_Major_Pentatonic")
			input = ""
		case "cmpb":
			scale, _, title = scales.GetScale("C_Minor_Pentatonic_Blues")
			input = ""
		case "cMpb":
			scale, _, title = scales.GetScale("C_Major_Pentatonic_Blues")
			input = ""
		case "C#":
			scale, _, title = scales.GetScale("C_Sharp_Major")
			input = ""
		case "c#m":
			scale, _, title = scales.GetScale("C_Sharp_Minor")
			input = ""
		case "c#mp":
			scale, _, title = scales.GetScale("C_Sharp_Minor_Pentatonic")
			input = ""
		case "c#Mp":
			scale, _, title = scales.GetScale("C_Sharp_Major_Pentatonic")
			input = ""
		case "c#mpb":
			scale, _, title = scales.GetScale("C_Sharp_Minor_Pentatonic_Blues")
			input = ""
		case "c#Mpb":
			scale, _, title = scales.GetScale("C_Sharp_Major_Pentatonic_Blues")
			input = ""

			// D
		case "D", "dM":
			scale, _, title = scales.GetScale("D_Major")
			input = ""
		case "dm":
			scale, _, title = scales.GetScale("D_Minor")
			input = ""
		case "dmp":
			scale, _, title = scales.GetScale("D_Minor_Pentatonic")
			input = ""
		case "dMp":
			scale, _, title = scales.GetScale("D_Major_Pentatonic")
			input = ""
		case "dmpb":
			scale, _, title = scales.GetScale("D_Minor_Pentatonic_Blues")
			input = ""
		case "dMpb":
			scale, _, title = scales.GetScale("D_Major_Pentatonic_Blues")
			input = ""
		case "D#":
			scale, _, title = scales.GetScale("D_Sharp_Major")
			input = ""
		case "d#m":
			scale, _, title = scales.GetScale("D_Sharp_Minor")
			input = ""
		case "d#mp":
			scale, _, title = scales.GetScale("D_Sharp_Minor_Pentatonic")
			input = ""
		case "d#Mp":
			scale, _, title = scales.GetScale("D_Sharp_Major_Pentatonic")
			input = ""
		case "d#mpb":
			scale, _, title = scales.GetScale("D_Sharp_Minor_Pentatonic_Blues")
			input = ""
		case "d#Mpb":
			scale, _, title = scales.GetScale("D_Sharp_Major_Pentatonic_Blues")
			input = ""

			// E
		case "E", "eM":
			scale, _, title = scales.GetScale("E_Major")
			input = ""
		case "em":
			scale, _, title = scales.GetScale("E_Minor")
			input = ""
		case "emp":
			scale, _, title = scales.GetScale("E_Minor_Pentatonic")
			input = ""
		case "eMp":
			scale, _, title = scales.GetScale("E_Major_Pentatonic")
			input = ""
		case "empb":
			scale, _, title = scales.GetScale("E_Minor_Pentatonic_Blues")
			input = ""
		case "eMpb":
			scale, _, title = scales.GetScale("E_Major_Pentatonic_Blues")
			input = ""

			// F
		case "F", "fM":
			scale, _, title = scales.GetScale("F_Major")
			input = ""
		case "fm":
			scale, _, title = scales.GetScale("F_Minor")
			input = ""
		case "fmp":
			scale, _, title = scales.GetScale("F_Minor_Pentatonic")
			input = ""
		case "fMp":
			scale, _, title = scales.GetScale("F_Major_Pentatonic")
			input = ""
		case "fmpb":
			scale, _, title = scales.GetScale("F_Minor_Pentatonic_Blues")
			input = ""
		case "fMpb":
			scale, _, title = scales.GetScale("F_Major_Pentatonic_Blues")
			input = ""
		case "F#":
			scale, _, title = scales.GetScale("F_Sharp_Major")
			input = ""
		case "f#m":
			scale, _, title = scales.GetScale("F_Sharp_Minor")
			input = ""
		case "f#mp":
			scale, _, title = scales.GetScale("F_Sharp_Minor_Pentatonic")
			input = ""
		case "f#Mp":
			scale, _, title = scales.GetScale("F_Sharp_Major_Pentatonic")
			input = ""
		case "f#mpb":
			scale, _, title = scales.GetScale("F_Sharp_Minor_Pentatonic_Blues")
			input = ""
		case "f#Mpb":
			scale, _, title = scales.GetScale("F_Sharp_Major_Pentatonic_Blues")
			input = ""

			// G
		case "G", "gM":
			scale, _, title = scales.GetScale("G_Major")
			input = ""
		case "gm":
			scale, _, title = scales.GetScale("G_Minor")
			input = ""
		case "gmp":
			scale, _, title = scales.GetScale("G_Minor_Pentatonic")
			input = ""
		case "gMp":
			scale, _, title = scales.GetScale("G_Major_Pentatonic")
			input = ""
		case "gmpb":
			scale, _, title = scales.GetScale("G_Minor_Pentatonic_Blues")
			input = ""
		case "gMpb":
			scale, _, title = scales.GetScale("G_Major_Pentatonic_Blues")
			input = ""
		case "G#":
			scale, _, title = scales.GetScale("G_Sharp_Major")
			input = ""
		case "g#m":
			scale, _, title = scales.GetScale("G_Sharp_Minor")
			input = ""
		case "g#mp":
			scale, _, title = scales.GetScale("G_Sharp_Minor_Pentatonic")
			input = ""
		case "g#Mp":
			scale, _, title = scales.GetScale("G_Sharp_Major_Pentatonic")
			input = ""
		case "g#mpb":
			scale, _, title = scales.GetScale("G_Sharp_Minor_Pentatonic_Blues")
			input = ""
		case "g#Mpb":
			scale, _, title = scales.GetScale("G_Sharp_Major_Pentatonic_Blues")
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
		case "a#mp":
			scale, _, title = scales.GetScale("A_Sharp_Minor_Pentatonic")
			input = ""
		case "a#Mp":
			scale, _, title = scales.GetScale("A_Sharp_Major_Pentatonic")
			input = ""
		case "a#mpb":
			scale, _, title = scales.GetScale("A_Sharp_Minor_Pentatonic_Blues")
			input = ""
		case "a#Mpb":
			scale, _, title = scales.GetScale("A_Sharp_Major_Pentatonic_Blues")
			input = ""
		case "amp":
			scale, _, title = scales.GetScale("A_Minor_Pentatonic")
			input = ""
		case "aMp":
			scale, _, title = scales.GetScale("A_Major_Pentatonic")
			input = ""
		case "ampb":
			scale, _, title = scales.GetScale("A_Minor_Pentatonic_Blues")
			input = ""
		case "aMpb":
			scale, _, title = scales.GetScale("A_Major_Pentatonic_Blues")
			input = ""

			// B
		case "B", "bM":
			scale, _, title = scales.GetScale("B_Major")
			input = ""
		case "bm":
			scale, _, title = scales.GetScale("B_Minor")
			input = ""
		case "bmp":
			scale, _, title = scales.GetScale("B_Minor_Pentatonic")
			input = ""
		case "bMp":
			scale, _, title = scales.GetScale("B_Major_Pentatonic")
			input = ""
		case "bmpb":
			scale, _, title = scales.GetScale("B_Minor_Pentatonic_Blues")
			input = ""
		case "bMpb":
			scale, _, title = scales.GetScale("B_Major_Pentatonic_Blues")
			input = ""
		}

		input = ""

		if title != "" {
			return scale, title, true
		}

	case *tcell.EventResize:
		s.Sync()
	}
	return nil, "", true
}

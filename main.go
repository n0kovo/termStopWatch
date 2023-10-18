package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/getwe/figlet4go"
	"github.com/spf13/pflag"
)

const (
	ZeroTimeStrWithMillis    = "00:00:00.000"
	ZeroTimeStrWithoutMillis = "00:00:00"
	TickerInterval           = 10 * time.Millisecond
)

var (
	defaultFont  = pflag.StringP("font", "f", "Georgia11", "Font to use. Must be in fonts/ directory.")
	hideHelp     = pflag.BoolP("hide-help", "q", false, "Hide help text.")
	showDebug    = pflag.BoolP("debug", "d", false, "Show debug text.")
	pauseAtStart = pflag.BoolP("pause", "p", false, "Start paused.")
	hideMillis   = pflag.BoolP("noms", "n", false, "Hide milliseconds.")
)

func drawHelpText(screen tcell.Screen, helpText []string, currentFont string, debugString string, showDebug bool) {
	_, height := screen.Size()

	helpText = append(helpText, fmt.Sprintf("Font: %s", currentFont))
	if showDebug {
		helpText = append(helpText, fmt.Sprintf("Debug: %s", debugString))
	}
	startY := height - len(helpText) // Start Y-coordinate for help text

	for i, line := range helpText {
		for x, ch := range []rune(line) {
			screen.SetContent(x, startY+i, ch, nil, tcell.StyleDefault)
		}
	}
}

func handleKeyEvent(screen tcell.Screen, quit, pause, reset, toggleMillis, toggleHelp, nextFont, prevFont, debug chan struct{}) {
	go func() {
		keyActionMap := map[rune]chan struct{}{
			' ': pause,
			'r': reset,
			'm': toggleMillis,
			'h': toggleHelp,
			'd': debug,
		}

		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyCtrlC || ev.Key() == tcell.KeyEsc {
					close(quit)
					return
				}
				lowerRune := unicode.ToLower(ev.Rune())
				if actionChan, exists := keyActionMap[lowerRune]; exists {
					actionChan <- struct{}{}
				}
				switch ev.Key() {
				case tcell.KeyRight:
					nextFont <- struct{}{}
				case tcell.KeyLeft:
					prevFont <- struct{}{}
				}

			}
		}
	}()
}

func drawAscii(screen tcell.Screen, asciiStr string, startX, startY int) {
	for _, line := range strings.Split(asciiStr, "\n") {
		for i, ch := range []rune(line) {
			screen.SetContent(startX+i, startY, ch, nil, tcell.StyleDefault)
		}
		startY++
	}
}

func main() {
	pflag.Parse()
	DefaultFont := *defaultFont

	helpText := []string{
		"(SPACE) Pause",
		"(M) Milliseconds",
		"(R) Reset",
		"(H) Hide help",
		"(ESC) Exit",
	}

	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Println("Error creating new screen:", err)
		return
	}

	if err := screen.Init(); err != nil {
		fmt.Println("Error initializing screen:", err)
		return
	}
	defer screen.Fini()

	var asciiRender *figlet4go.AsciiRender
	var options *figlet4go.RenderOptions

	updateAsciiRenderer := func(currentFontName string, showMillis bool) int {
		asciiRender = figlet4go.NewAsciiRender()
		options = figlet4go.NewRenderOptions()
		options.FontName = currentFontName
		asciiRender.LoadFont("fonts/")

		zeroTimeWithMillis, _ := asciiRender.RenderOpts(ZeroTimeStrWithMillis, options)
		zeroTimeWithoutMillis, _ := asciiRender.RenderOpts(ZeroTimeStrWithoutMillis, options)

		zeroWidthWithMillis := len(strings.Split(zeroTimeWithMillis, "\n")[1])
		zeroWidthWithoutMillis := len(strings.Split(zeroTimeWithoutMillis, "\n")[1])

		if showMillis {
			return zeroWidthWithMillis
		}
		return zeroWidthWithoutMillis
	}

	currentZeroWidth := updateAsciiRenderer(DefaultFont, true)
	_ = currentZeroWidth // Placeholder to avoid compiler warning

	fontDir := "fonts/"
	files, err := os.ReadDir(fontDir)
	if err != nil {
		fmt.Println("Error reading font directory:", err)
		return
	}

	var fontNames []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fontNames = append(fontNames, strings.TrimSuffix(file.Name(), ".flf"))
	}
	sort.Strings(fontNames)
	// Find index of default font
	currentFontIndex := sort.SearchStrings(fontNames, DefaultFont)
	if currentFontIndex >= len(fontNames) || fontNames[currentFontIndex] != DefaultFont {
		fmt.Println("Error finding default font.")
		// Default to first font
		currentFontIndex = 0
	}
	options.FontName = fontNames[currentFontIndex]

	quit := make(chan struct{})
	pause := make(chan struct{})
	reset := make(chan struct{})
	toggleMillis := make(chan struct{})
	toggleHelp := make(chan struct{})
	toggleDebug := make(chan struct{})

	nextFont := make(chan struct{})
	prevFont := make(chan struct{})

	showMillis := !*hideMillis
	showHelp := !*hideHelp
	showDebug := *showDebug

	width, _ := screen.Size()

	handleKeyEvent(screen, quit, pause, reset, toggleMillis, toggleHelp, nextFont, prevFont, toggleDebug)

	ticker := time.NewTicker(TickerInterval)
	defer ticker.Stop()

	startTime := time.Now()
	elapsed := 0 * time.Second
	isPaused := *pauseAtStart
	fontChanged := false
	msViewChanged := true

	for {
		select {
		case <-quit:
			return
		case <-pause:
			isPaused = !isPaused
		case <-reset:
			elapsed = 0
			startTime = time.Now()
		case <-toggleHelp:
			showHelp = !showHelp
		case <-toggleDebug:
			showDebug = !showDebug
		case <-toggleMillis:
			showMillis = !showMillis
			msViewChanged = true
		case <-nextFont:
			currentFontIndex = (currentFontIndex + 1) % len(fontNames)
			fontChanged = true
		case <-prevFont:
			currentFontIndex = (currentFontIndex - 1 + len(fontNames)) % len(fontNames)
			fontChanged = true

		case <-ticker.C:
			if !isPaused {
				elapsed = time.Since(startTime)
			}
			var timeStr string
			h, m, s := int(elapsed.Hours()), int(elapsed.Minutes())%60, int(elapsed.Seconds())%60
			ms := int(elapsed.Milliseconds()) % 1000

			if showMillis {
				timeStr = fmt.Sprintf("%02d:%02d:%02d.%03d", h, m, s, ms)
			} else {
				timeStr = fmt.Sprintf("%02d:%02d:%02d", h, m, s)
			}

			if fontChanged || msViewChanged {
				currentZeroWidth = updateAsciiRenderer(fontNames[currentFontIndex], showMillis)
				fontChanged = false
				msViewChanged = false
			}
			screen.Clear()
			_, height := screen.Size()
			fixedStartX := width/2 - currentZeroWidth/2
			asciiTime, _ := asciiRender.RenderOpts(timeStr, options)

			lines := strings.Split(asciiTime, "\n")
			startY := height/2 - len(lines)/2

			if showHelp {
				debugString := fmt.Sprintf("startY: %d, currentZeroWidth: %d, len(lines): %d, fixedStartX: %d", startY, currentZeroWidth, len(lines), fixedStartX)
				drawHelpText(screen, helpText, options.FontName, debugString, showDebug)
			}
			drawAscii(screen, asciiTime, fixedStartX, startY)
			screen.Show()

		}
	}
}

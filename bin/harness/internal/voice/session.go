package voice

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"harness/internal/session"
)

// Mode defines the voice interaction mode.
type Mode int

const (
	ModeOff Mode = iota
	ModeListen        // Only listen (STT), response as text
	ModeSpeak         // Only speak (TTS), input as text
	ModeFull          // Full voice: STT + TTS
)

func (m Mode) String() string {
	switch m {
	case ModeListen:
		return "listen"
	case ModeSpeak:
		return "speak"
	case ModeFull:
		return "full"
	default:
		return "off"
	}
}

// Config holds voice session configuration.
type Config struct {
	Mode       Mode
	Duration   int  // seconds per recording
	Continuous bool // continuous listening loop
	WakeWord   string
}

// DefaultConfig returns the default voice configuration.
func DefaultConfig() Config {
	return Config{
		Mode:       ModeFull,
		Duration:   defaultDuration,
		Continuous: false,
		WakeWord:   "",
	}
}

// Session manages voice interaction for a CLI session.
type Session struct {
	Config
	root    string
	apiKey  string
	tempDir string
	logger  *SessionLogger
}

// NewSession creates a new voice session.
func NewSession(root string, cfg Config) *Session {
	apiKey := os.Getenv("GEMINI_API_KEY")
	return &Session{
		Config:  cfg,
		root:    root,
		apiKey:  apiKey,
		tempDir: EnsureTempDir(root),
		logger:  NewSessionLogger(),
	}
}

// SetTree associates a session tree (for metadata, not required).
func (vs *Session) SetTree(_ *session.Tree) {
	// placeholder for future tree logging
}

// Listen records audio and returns transcribed text.
func (vs *Session) Listen() (string, error) {
	recordPath := filepath.Join(vs.tempDir,
		fmt.Sprintf("recording_%d.wav", time.Now().UnixMilli()))

	vs.logger.Log("Recording for %d seconds...", vs.Duration)
	fmt.Fprintf(os.Stderr, "\033[1;33m🎤 [REC %ds] Speak now!\033[0m\n", vs.Duration)

	if err := RecordWAV(recordPath, vs.Duration); err != nil {
		return "", fmt.Errorf("record: %w", err)
	}

	// Check file size — if too small, probably no speech
	info, _ := os.Stat(recordPath)
	if info != nil && info.Size() < 1000 {
		return "", fmt.Errorf("audio too short (%d bytes)", info.Size())
	}

	vs.logger.Log("Transcribing via Gemini...")
	fmt.Fprintf(os.Stderr, "\033[90m⟲ Transcribing...\033[0m\n")

	text, err := TranscribeWAV(recordPath, vs.apiKey)
	if err != nil {
		return "", fmt.Errorf("transcribe: %w", err)
	}

	// Cleanup old recordings in background
	go Cleanup(vs.root, 30*time.Minute)

	return text, nil
}

// Speak synthesizes and plays the response through speakers.
func (vs *Session) Speak(text string) error {
	if text == "" || vs.Mode == ModeListen {
		return nil
	}

	// Truncate to speakable length
	spoken := truncateForSpeech(text)

	vs.logger.Log("Speaking (%d chars)...", len(spoken))
	fmt.Fprintf(os.Stderr, "\033[90m🔊 Speaking...\033[0m\n")

	return Speak(spoken, vs.apiKey)
}

// InteractiveLoop runs a continuous voice interaction loop.
// The callback receives transcribed text and returns the AI response.
func (vs *Session) InteractiveLoop(callback func(string) (string, error)) error {
	// Ensure temp dir
	os.MkdirAll(vs.tempDir, 0755)

	fmt.Fprintf(os.Stderr, "\033[1;36m")
	fmt.Fprintf(os.Stderr, "\n╔══════════════════════════════════════════════╗\n")
	fmt.Fprintf(os.Stderr, "║     🎤 Voice Mode: %-6s               ║\n", strings.ToUpper(vs.Mode.String()))
	fmt.Fprintf(os.Stderr, "║     Recording: %ds / Continuous: %-5t  ║\n", vs.Duration, vs.Continuous)
	fmt.Fprintf(os.Stderr, "║     Press Enter to record  |  Ctrl+C quit║\n")
	fmt.Fprintf(os.Stderr, "╚══════════════════════════════════════════════╝\n")
	fmt.Fprintf(os.Stderr, "\033[0m\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		var text string
		var err error

		if vs.apiKey != "" {
			fmt.Fprintf(os.Stderr, "\033[1;35m🎤 Press Enter to speak (type '/t' for text, '/q' to quit):\033[0m ")
			var promptConfirm string
			fmt.Scanln(&promptConfirm)
			promptConfirm = strings.TrimSpace(promptConfirm)

			if promptConfirm == "/q" || promptConfirm == "/quit" || promptConfirm == "/exit" {
				break
			}

			if promptConfirm == "/t" || promptConfirm == "/text" {
				fmt.Fprintf(os.Stderr, "\033[1;33m⌨️  Digite sua mensagem:\033[0m ")
				keyboardInput, _ := reader.ReadString('\n')
				text = strings.TrimSpace(keyboardInput)
			} else {
				// Record and transcribe
				text, err = vs.Listen()
				if err != nil {
					fmt.Fprintf(os.Stderr, "\033[38;5;208m⚠️  Voice STT error: %v\033[0m\n", err)
					fmt.Fprintf(os.Stderr, "\033[1;33m👉 Fallback (Keyboard): Digite sua mensagem:\033[0m ")
					keyboardInput, _ := reader.ReadString('\n')
					text = strings.TrimSpace(keyboardInput)
				}
			}
		} else {
			// No API key - use local/free keyboard resource
			fmt.Fprintf(os.Stderr, "\033[1;30m[STT Offline — GEMINI_API_KEY missing]\033[0m\n")
			fmt.Fprintf(os.Stderr, "\033[1;33m⌨️  Digite sua mensagem (or /q to quit):\033[0m ")
			keyboardInput, _ := reader.ReadString('\n')
			text = strings.TrimSpace(keyboardInput)
			if text == "/q" || text == "/quit" || text == "/exit" {
				break
			}
		}

		if text == "" {
			continue
		}

		fmt.Fprintf(os.Stderr, "\033[1;32m🎤 You: \033[0m%s\n", text)

		// Process through callback
		response, err := callback(text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[38;5;196mError: %v\033[0m\n", err)
			continue
		}

		// Print response
		fmt.Fprintf(os.Stderr, "\033[1;34m🤖 AI:\033[0m %s\n", response)

		// Speak response
		if vs.Mode == ModeFull || vs.Mode == ModeSpeak {
			if err := vs.Speak(response); err != nil {
				fmt.Fprintf(os.Stderr, "\033[38;5;208mTTS: %v\033[0m\n", err)
			}
		}

		// Short pause before next iteration
		time.Sleep(200 * time.Millisecond)
	}

	return nil
}

// OneShot records, transcribes, and returns the text.
func (vs *Session) OneShot() (string, error) {
	return vs.Listen()
}

// truncateForSpeech cuts text at a natural sentence boundary within max length.
func truncateForSpeech(text string) string {
	const maxSpeechLen = 500
	if len(text) <= maxSpeechLen {
		return text
	}

	// Look for last sentence boundary within limit
	trunc := text[:maxSpeechLen]
	cut := strings.LastIndexAny(trunc, ".!?")
	if cut > maxSpeechLen/2 {
		return text[:cut+1]
	}
	return text[:maxSpeechLen] + "..."
}

// ── Logger ──────────────────────────────────────────────────────────────

// SessionLogger provides structured logging for voice sessions.
type SessionLogger struct {
	entries []string
}

func NewSessionLogger() *SessionLogger {
	return &SessionLogger{}
}

func (l *SessionLogger) Log(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.entries = append(l.entries, msg)
}

func (l *SessionLogger) Entries() []string {
	return l.entries
}

// ── Dependencies Check ──────────────────────────────────────────────────

// CheckDeps verifies that required system tools for recording are available.
// It detects the OS at runtime and checks the appropriate tools.
// If recording is not possible, it exits with os.Exit(1).
// Returns a list of missing TTS tools (optional).
func CheckDeps() (missingTTS []string) {
	hasRecorder := false
	var recorderErr string

	switch runtime.GOOS {
	case "linux":
		// Check for arecord, rec, or ffmpeg
		for _, tool := range []string{"arecord", "rec", "ffmpeg"} {
			if _, err := exec.LookPath(tool); err == nil {
				hasRecorder = true
				break
			}
		}
		if !hasRecorder {
			recorderErr = "Required recording tool not found: arecord, rec, or ffmpeg.\n   Install ALSA utils: sudo apt install alsa-utils\n   Or SoX: sudo apt install sox\n   Or FFmpeg: sudo apt install ffmpeg"
		}

		// Check optional TTS tools
		for _, tool := range []string{"spd-say", "espeak-ng", "espeak"} {
			if _, err := exec.LookPath(tool); err != nil {
				missingTTS = append(missingTTS, tool)
			}
		}

	case "darwin":
		// Check for rec or ffmpeg
		for _, tool := range []string{"rec", "ffmpeg"} {
			if _, err := exec.LookPath(tool); err == nil {
				hasRecorder = true
				break
			}
		}
		if !hasRecorder {
			recorderErr = "Required recording tool not found: rec (SoX) or ffmpeg.\n   Install SoX: brew install sox\n   Or FFmpeg: brew install ffmpeg"
		}

		// macOS has native 'say' tool pre-installed, no missing TTS

	case "windows":
		// Windows has built-in C# MCI recording via powershell (no external tool needed)
		hasRecorder = true

		// Windows has native PowerShell SpeechSynthesizer, no missing TTS
	}

	if !hasRecorder {
		fmt.Fprintf(os.Stderr, "\033[38;5;196m❌ Recording dependencies missing on %s:\033[0m\n", runtime.GOOS)
		fmt.Fprintf(os.Stderr, "   %s\n", recorderErr)
		os.Exit(1)
	}

	return missingTTS
}

// PrintDepsWarning prints a warning about missing TTS tools.
func PrintDepsWarning(missing []string) {
	if len(missing) > 0 {
		fmt.Fprintf(os.Stderr, "\033[38;5;208m⚠️  TTS not available (%s).\033[0m\n", strings.Join(missing, ", "))
		fmt.Fprintf(os.Stderr, "\033[38;5;208m   Responses will be text-only. Install: sudo apt install espeak-ng\033[0m\n")
	}
}
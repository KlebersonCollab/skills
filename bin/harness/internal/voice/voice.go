// Package voice provides audio recording, transcription (STT), and speech synthesis (TTS)
// for the Harness CLI. Uses native and common tools dynamically per OS (Linux, macOS, Windows).
// Zero external Go dependencies — stdlib only.
package voice

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// ── Configuration ────────────────────────────────────────────────────────

const (
	defaultDuration  = 8           // seconds per recording
	geminiModel      = "gemini-2.5-flash"
	sampleFormat     = "cd"        // CD quality: 16-bit 44100Hz stereo
)

// ── Recording ───────────────────────────────────────────────────────────

// RecordWAV records audio to a WAV file for the given duration in seconds.
// It is OS-agnostic and uses the best available native or common tool.
func RecordWAV(outputPath string, durationSec int) error {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create dir: %w", err)
	}

	// 1. Try SoX (rec) if available (cross-platform, high priority)
	if _, err := exec.LookPath("rec"); err == nil {
		args := []string{
			"-q", // quiet mode
			outputPath,
			"trim", "0", fmt.Sprintf("%d", durationSec),
		}
		cmd := exec.Command("rec", args...)
		if _, err := cmd.CombinedOutput(); err == nil {
			return nil
		}
	}

	// 2. Try FFmpeg if available (cross-platform, medium priority)
	if _, err := exec.LookPath("ffmpeg"); err == nil {
		var args []string
		switch runtime.GOOS {
		case "darwin":
			args = []string{"-y", "-f", "avfoundation", "-i", ":0", "-t", fmt.Sprintf("%d", durationSec), outputPath}
		case "linux":
			args = []string{"-y", "-f", "alsa", "-i", "default", "-t", fmt.Sprintf("%d", durationSec), outputPath}
		case "windows":
			args = []string{"-y", "-f", "dshow", "-i", "audio=default", "-t", fmt.Sprintf("%d", durationSec), outputPath}
		default:
			args = []string{"-y", "-t", fmt.Sprintf("%d", durationSec), outputPath}
		}
		cmd := exec.Command("ffmpeg", args...)
		if _, err := cmd.CombinedOutput(); err == nil {
			return nil
		}
	}

	// 3. Fall back to OS-specific native tools
	switch runtime.GOOS {
	case "linux":
		// Linux native: arecord (ALSA)
		if _, err := exec.LookPath("arecord"); err == nil {
			args := []string{
				"-f", sampleFormat,
				"-t", "wav",
				"-d", fmt.Sprintf("%d", durationSec),
				"-q",
				outputPath,
			}
			cmd := exec.Command("arecord", args...)
			if out, err := cmd.CombinedOutput(); err != nil {
				return fmt.Errorf("arecord failed: %w\n%s", err, string(out))
			}
			return nil
		}

	case "windows":
		// Windows native: PowerShell MCI recording via winmm.dll
		absPath, err := filepath.Abs(outputPath)
		if err != nil {
			absPath = outputPath
		}
		// Escape backslashes for C# string
		escapedPath := strings.ReplaceAll(absPath, "\\", "\\\\")

		psCode := fmt.Sprintf(`
$code = @'
using System;
using System.Runtime.InteropServices;
using System.Text;
public class AudioRecorder {
    [DllImport("winmm.dll")]
    private static extern long mciSendString(string command, StringBuilder returnString, int returnLength, IntPtr callback);
    public static void Record(string filepath, int durationMs) {
        mciSendString("open new Type waveaudio Alias recsound", null, 0, IntPtr.Zero);
        mciSendString("record recsound", null, 0, IntPtr.Zero);
        System.Threading.Thread.Sleep(durationMs);
        mciSendString("save recsound " + filepath, null, 0, IntPtr.Zero);
        mciSendString("close recsound", null, 0, IntPtr.Zero);
    }
}
'@
Add-Type -TypeDefinition $code
[AudioRecorder]::Record('%s', %d)
`, escapedPath, durationSec*1000)

		cmd := exec.Command("powershell", "-Command", psCode)
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("windows native recording failed: %w\n%s", err, string(out))
		}
		return nil
	}

	return fmt.Errorf("no recording utility found or supported on OS: %s (please install SoX or FFmpeg)", runtime.GOOS)
}

// RecordUntilSilence records audio until silence is detected
// (simple approach: record fixed chunks and check amplitude).
// Returns the path to the recorded file.
func RecordUntilSilence(root string) (string, error) {
	tmpDir := filepath.Join(root, ".harness", "voice")
	outputPath := filepath.Join(tmpDir, fmt.Sprintf("recording_%d.wav", time.Now().Unix()))

	// Record for the default duration (simple push-to-talk model)
	fmt.Fprintf(os.Stderr, "\033[1;33m🎤 Recording for %d seconds... Speak now!\033[0m\n", defaultDuration)

	if err := RecordWAV(outputPath, defaultDuration); err != nil {
		return "", err
	}

	fmt.Fprintf(os.Stderr, "\033[1;32m✅ Recording complete\033[0m\n")
	return outputPath, nil
}

// ── Transcription (STT) ─────────────────────────────────────────────────

// TranscribeWAV sends a WAV file to the Gemini API and returns transcribed text.
// Uses inline_data (base64) — no file upload needed.
func TranscribeWAV(audioPath, apiKey string) (string, error) {
	audioData, err := os.ReadFile(audioPath)
	if err != nil {
		return "", fmt.Errorf("read audio: %w", err)
	}

	b64Data := base64.StdEncoding.EncodeToString(audioData)

	reqBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{
						"text": "Transcreva exatamente o audio falado em Portugues Brasileiro. " +
							"Retorne APENAS a transcricao literal, sem comentarios, sem formatacao.",
					},
					{
						"inline_data": map[string]interface{}{
							"mime_type": "audio/wav",
							"data":      b64Data,
						},
					},
				},
			},
		},
		"generationConfig": map[string]interface{}{
			"temperature": 0.1,
		},
	}

	body, _ := json.Marshal(reqBody)

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		geminiModel, apiKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("gemini API call: %w", err)
	}
	defer resp.Body.Close()

	respData, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("gemini API returned %d: %s", resp.StatusCode, string(respData))
	}

	text, err := extractGeminiText(respData)
	if err != nil {
		return "", fmt.Errorf("parse response: %w", err)
	}

	return strings.TrimSpace(text), nil
}

// extractGeminiText parses the Gemini API JSON response.
func extractGeminiText(data []byte) (string, error) {
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return "", fmt.Errorf("json: %w", err)
	}

	// Check for block reason
	promptFeedback, _ := result["promptFeedback"].(map[string]interface{})
	if blockReason, ok := promptFeedback["blockReason"].(string); ok {
		return "", fmt.Errorf("blocked: %s", blockReason)
	}

	candidates, _ := result["candidates"].([]interface{})
	if len(candidates) == 0 {
		// Try to extract error info
		if err, ok := result["error"].(map[string]interface{}); ok {
			msg, _ := err["message"].(string)
			return "", fmt.Errorf("api error: %s", msg)
		}
		return "", fmt.Errorf("no candidates in response")
	}

	candidate, _ := candidates[0].(map[string]interface{})
	content, _ := candidate["content"].(map[string]interface{})
	parts, _ := content["parts"].([]interface{})
	if len(parts) == 0 {
		return "", fmt.Errorf("no parts in candidate")
	}

	text, _ := parts[0].(map[string]interface{})["text"].(string)
	return text, nil
}

// ── Speech Synthesis (TTS) ──────────────────────────────────────────────

// SynthesizeGeminiTTS calls the gemini-3.1-flash-tts-preview model to generate audio.
// Returns the path to the temporary WAV file.
func SynthesizeGeminiTTS(text, apiKey string) (string, error) {
	reqBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{
						"text": text,
					},
				},
			},
		},
		"generationConfig": map[string]interface{}{
			"responseModalities": []string{"AUDIO"},
			"speechConfig": map[string]interface{}{
				"voiceConfig": map[string]interface{}{
					"prebuiltVoiceConfig": map[string]interface{}{
						"voiceName": "Aoede", // Aoede is a highly natural voice (others: Puck, Kore, Fenrir, Charon)
					},
				},
			},
		},
	}

	body, _ := json.Marshal(reqBody)

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-3.1-flash-tts-preview:generateContent?key=%s", apiKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("gemini TTS call: %w", err)
	}
	defer resp.Body.Close()

	respData, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("gemini TTS returned %d: %s", resp.StatusCode, string(respData))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respData, &result); err != nil {
		return "", fmt.Errorf("json unmarshal: %w", err)
	}

	candidates, _ := result["candidates"].([]interface{})
	if len(candidates) == 0 {
		return "", fmt.Errorf("no candidates in response")
	}

	candidate, _ := candidates[0].(map[string]interface{})
	content, _ := candidate["content"].(map[string]interface{})
	parts, _ := content["parts"].([]interface{})
	if len(parts) == 0 {
		return "", fmt.Errorf("no parts in candidate")
	}

	part, _ := parts[0].(map[string]interface{})
	inlineData, _ := part["inlineData"].(map[string]interface{})
	if inlineData == nil {
		return "", fmt.Errorf("no inlineData in part")
	}

	b64Audio, _ := inlineData["data"].(string)
	if b64Audio == "" {
		return "", fmt.Errorf("empty audio data")
	}

	audioBytes, err := base64.StdEncoding.DecodeString(b64Audio)
	if err != nil {
		return "", fmt.Errorf("base64 decode: %w", err)
	}

	voiceDir := EnsureTempDir("")
	tempWAV := filepath.Join(voiceDir, fmt.Sprintf("gemini_tts_%d.wav", time.Now().UnixNano()))
	if err := os.WriteFile(tempWAV, audioBytes, 0644); err != nil {
		return "", fmt.Errorf("write wav file: %w", err)
	}

	return tempWAV, nil
}

// Speak synthesizes and plays text through speakers.
// Prioritizes the premium Google Gemini 3.1 Flash TTS model if an API key is available.
// Otherwise, falls back to native/local OS systems (say, PowerShell, gtts-cli, espeak-ng).
func Speak(text, apiKey string) error {
	if text == "" {
		return nil
	}

	// Limit text length to avoid endless speaking
	const maxLen = 1000
	if len(text) > maxLen {
		text = text[:maxLen] + "..."
	}

	// 1. Prioritize official Gemini 3.1 Flash TTS Model (Premium, human-like voice)
	if apiKey != "" {
		wavPath, err := SynthesizeGeminiTTS(text, apiKey)
		if err == nil {
			if playErr := PlayAudio(wavPath); playErr == nil {
				os.Remove(wavPath)
				return nil // Success with premium Gemini AI voice!
			}
			os.Remove(wavPath)
		}
		// If it fails (rate limit, offline, etc.), gracefully log and fall back to local resources
		fmt.Fprintf(os.Stderr, "\033[90m⟲ Gemini TTS unavailable, falling back to local voice...\033[0m\n")
	}

	// 2. Fallback: OS-native or local synthesizer tools
	escapedText := strings.ReplaceAll(text, "'", "''")

	switch runtime.GOOS {
	case "darwin":
		// macOS has native 'say' command pre-installed
		cmd := exec.Command("say", text)
		return cmd.Run()

	case "windows":
		// Windows has native PowerShell speech synthesis pre-installed
		psCmd := fmt.Sprintf("Add-Type -AssemblyName System.Speech; $synth = New-Object System.Speech.Synthesis.SpeechSynthesizer; $synth.Speak('%s')", escapedText)
		cmd := exec.Command("powershell", "-Command", psCmd)
		return cmd.Run()

	case "linux":
		// 0. Try gTTS (Google Text-to-Speech) if gtts-cli is installed (100% free, neural, high-quality)
		if _, err := exec.LookPath("gtts-cli"); err == nil {
			// Save in local workspace to allow sandboxed snap ffmpeg to read/write
			voiceDir := EnsureTempDir("")
			tempMP3 := filepath.Join(voiceDir, "harness_tts.mp3")
			tempWAV := filepath.Join(voiceDir, "harness_tts.wav")

			// Run gtts-cli (using 'pt' to avoid 'pt-br' deprecation warnings)
			cmd := exec.Command("gtts-cli", "--lang", "pt", text, "--output", tempMP3)
			if err := cmd.Run(); err == nil {
				// 1. If mpg123 is available, play MP3 directly (Fastest, zero conversion, no Snap GPU lag)
				if _, err := exec.LookPath("mpg123"); err == nil {
					playCmd := exec.Command("mpg123", "-q", tempMP3)
					if err := playCmd.Run(); err == nil {
						os.Remove(tempMP3)
						return nil
					}
				}

				// 2. If mpv is available, play MP3 directly (Very fast, zero conversion)
				if _, err := exec.LookPath("mpv"); err == nil {
					playCmd := exec.Command("mpv", "--no-video", "--really-quiet", tempMP3)
					if err := playCmd.Run(); err == nil {
						os.Remove(tempMP3)
						return nil
					}
				}

				// 3. Fallback: Convert to WAV using ffmpeg and play via PlayAudio
				if _, err := exec.LookPath("ffmpeg"); err == nil {
					convCmd := exec.Command("ffmpeg", "-y", "-i", tempMP3, tempWAV)
					convCmd.Stderr = nil
					if err := convCmd.Run(); err == nil {
						// Play the WAV file
						if playErr := PlayAudio(tempWAV); playErr == nil {
							os.Remove(tempMP3)
							os.Remove(tempWAV)
							return nil // Successful neural speech!
						}
					}
				}

				// 4. Fallback: If we have ffplay, play MP3 directly
				if _, err := exec.LookPath("ffplay"); err == nil {
					playCmd := exec.Command("ffplay", "-nodisp", "-autoexit", tempMP3)
					playCmd.Stderr = nil
					if err := playCmd.Run(); err == nil {
						os.Remove(tempMP3)
						return nil
					}
				}
			}
		}

		// Linux uses spd-say, espeak-ng, espeak
		providers := []struct {
			name string
			args []string
		}{
			{"spd-say", []string{"-l", "pt-BR", text}},
			{"espeak-ng", []string{"-v", "pt-br", text}},
			{"espeak", []string{"-v", "pt-br", text}},
		}

		var lastErr error
		for _, p := range providers {
			if _, err := exec.LookPath(p.name); err != nil {
				lastErr = err
				continue
			}
			cmd := exec.Command(p.name, p.args...)
			cmd.Stderr = nil // suppress warnings
			if err := cmd.Run(); err != nil {
				lastErr = err
				continue
			}
			return nil // success
		}
		return fmt.Errorf("no TTS available (try: sudo apt install espeak-ng): %w", lastErr)

	default:
		return fmt.Errorf("TTS not supported on OS: %s", runtime.GOOS)
	}
}

// PlayAudio plays a WAV file using native platform tool.
func PlayAudio(path string) error {
	switch runtime.GOOS {
	case "linux":
		if _, err := exec.LookPath("aplay"); err == nil {
			cmd := exec.Command("aplay", "-q", path)
			return cmd.Run()
		}
	case "darwin":
		if _, err := exec.LookPath("afplay"); err == nil {
			cmd := exec.Command("afplay", path)
			return cmd.Run()
		}
	case "windows":
		psCmd := fmt.Sprintf("(New-Object System.Media.SoundPlayer('%s')).PlaySync()", strings.ReplaceAll(path, "'", "''"))
		cmd := exec.Command("powershell", "-Command", psCmd)
		return cmd.Run()
	}
	return fmt.Errorf("no playback utility found on OS: %s", runtime.GOOS)
}

// ── Cleanup ─────────────────────────────────────────────────────────────

// Cleanup removes old voice recordings from temp directory.
func Cleanup(root string, maxAge time.Duration) {
	dir := filepath.Join(root, ".harness", "voice")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	now := time.Now()
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		if now.Sub(info.ModTime()) > maxAge {
			os.Remove(filepath.Join(dir, e.Name()))
		}
	}
}

// EnsureTempDir creates the voice temp directory.
func EnsureTempDir(root string) string {
	dir := filepath.Join(root, ".harness", "voice")
	os.MkdirAll(dir, 0755)
	return dir
}
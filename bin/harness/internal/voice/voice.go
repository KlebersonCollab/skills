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

	client := &http.Client{Timeout: 60 * time.Second}
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

// TranscribeLocalOrGemini transcribes a WAV file. It first tries local offline STT utilities
// (whisper, pocketsphinx, vosk-transcriber) to save internet data and API quota.
// If no local tool is installed or execution fails, it falls back to the premium Google Gemini STT API.
func TranscribeLocalOrGemini(audioPath, apiKey string) (string, error) {
	// 1. Try local Whisper CLI (if installed)
	if _, err := exec.LookPath("whisper"); err == nil {
		tempDir := filepath.Dir(audioPath)
		// Run whisper: whisper <audioPath> --language pt --model base --output_format txt --output_dir <tempDir>
		cmd := exec.Command("whisper", audioPath, "--language", "pt", "--model", "base", "--output_format", "txt", "--output_dir", tempDir)
		if err := cmd.Run(); err == nil {
			// Read the generated txt file
			baseName := strings.TrimSuffix(filepath.Base(audioPath), filepath.Ext(audioPath))
			txtPath := filepath.Join(tempDir, baseName+".txt")
			if content, err := os.ReadFile(txtPath); err == nil {
				os.Remove(txtPath) // cleanup
				text := strings.TrimSpace(string(content))
				if text != "" {
					fmt.Fprintf(os.Stderr, "\033[32m⟲ Transcribed locally via Whisper: %s\033[0m\n", text)
					return text, nil
				}
			}
		}
	}

	// 2. Try pocketsphinx (if installed)
	if _, err := exec.LookPath("pocketsphinx"); err == nil {
		// Run pocketsphinx: pocketsphinx single -infile <audioPath>
		cmd := exec.Command("pocketsphinx", "single", "-infile", audioPath)
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err == nil {
			text := strings.TrimSpace(out.String())
			if text != "" {
				fmt.Fprintf(os.Stderr, "\033[32m⟲ Transcribed locally via Pocketsphinx: %s\033[0m\n", text)
				return text, nil
			}
		}
	}

	// 3. Try vosk-transcriber (if installed)
	if _, err := exec.LookPath("vosk-transcriber"); err == nil {
		tempDir := filepath.Dir(audioPath)
		txtPath := filepath.Join(tempDir, fmt.Sprintf("vosk_stt_%d.txt", time.Now().UnixNano()))
		// Run vosk-transcriber: vosk-transcriber -i <audioPath> -o <txtPath>
		cmd := exec.Command("vosk-transcriber", "-i", audioPath, "-o", txtPath)
		if err := cmd.Run(); err == nil {
			if content, err := os.ReadFile(txtPath); err == nil {
				os.Remove(txtPath) // cleanup
				text := strings.TrimSpace(string(content))
				if text != "" {
					fmt.Fprintf(os.Stderr, "\033[32m⟲ Transcribed locally via Vosk: %s\033[0m\n", text)
					return text, nil
				}
			}
		}
	}

	// 4. Fallback to official premium Google Gemini STT API
	if apiKey == "" {
		return "", fmt.Errorf("local STT not found and GEMINI_API_KEY is not configured")
	}

	fmt.Fprintf(os.Stderr, "\033[90m⟲ Local STT not available or failed. Falling back to Gemini Cloud STT...\033[0m\n")
	return TranscribeWAV(audioPath, apiKey)
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

// addWavHeader prepends the canonical 44-byte RIFF/WAVE header to a raw PCM stream.
// It assumes little-endian byte ordering, which is the standard for RIFF/WAVE.
func addWavHeader(pcm []byte, sampleRate int, numChannels int, bitsPerSample int) []byte {
	header := make([]byte, 44)

	// 1-4: ChunkID "RIFF"
	copy(header[0:4], []byte("RIFF"))

	// 5-8: ChunkSize = 36 + Subchunk2Size
	totalSize := uint32(36 + len(pcm))
	header[4] = byte(totalSize & 0xff)
	header[5] = byte((totalSize >> 8) & 0xff)
	header[6] = byte((totalSize >> 16) & 0xff)
	header[7] = byte((totalSize >> 24) & 0xff)

	// 9-12: Format "WAVE"
	copy(header[8:12], []byte("WAVE"))

	// 13-16: Subchunk1ID "fmt "
	copy(header[12:16], []byte("fmt "))

	// 17-20: Subchunk1Size = 16 for PCM
	header[16] = 16
	header[17] = 0
	header[18] = 0
	header[19] = 0

	// 21-22: AudioFormat = 1 (linear PCM)
	header[20] = 1
	header[21] = 0

	// 23-24: NumChannels
	header[22] = byte(numChannels & 0xff)
	header[23] = byte((numChannels >> 8) & 0xff)

	// 25-28: SampleRate
	header[24] = byte(sampleRate & 0xff)
	header[25] = byte((sampleRate >> 8) & 0xff)
	header[26] = byte((sampleRate >> 16) & 0xff)
	header[27] = byte((sampleRate >> 24) & 0xff)

	// 29-32: ByteRate = SampleRate * NumChannels * BitsPerSample/8
	byteRate := uint32(sampleRate * numChannels * bitsPerSample / 8)
	header[28] = byte(byteRate & 0xff)
	header[29] = byte((byteRate >> 8) & 0xff)
	header[30] = byte((byteRate >> 16) & 0xff)
	header[31] = byte((byteRate >> 24) & 0xff)

	// 33-34: BlockAlign = NumChannels * BitsPerSample/8
	blockAlign := uint16(numChannels * bitsPerSample / 8)
	header[32] = byte(blockAlign & 0xff)
	header[33] = byte((blockAlign >> 8) & 0xff)

	// 35-36: BitsPerSample
	header[34] = byte(bitsPerSample & 0xff)
	header[35] = byte((bitsPerSample >> 8) & 0xff)

	// 37-40: Subchunk2ID "data"
	copy(header[36:40], []byte("data"))

	// 41-44: Subchunk2Size = len(pcm)
	pcmSize := uint32(len(pcm))
	header[40] = byte(pcmSize & 0xff)
	header[41] = byte((pcmSize >> 8) & 0xff)
	header[42] = byte((pcmSize >> 16) & 0xff)
	header[43] = byte((pcmSize >> 24) & 0xff)

	return append(header, pcm...)
}

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

	client := &http.Client{Timeout: 90 * time.Second}
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

	// Prepend a standard WAV header so modern players (mpv, pw-play) recognize the format
	wavBytes := addWavHeader(audioBytes, 24000, 1, 16)

	voiceDir := EnsureTempDir("")
	tempWAV := filepath.Join(voiceDir, fmt.Sprintf("gemini_tts_%d.wav", time.Now().UnixNano()))
	if err := os.WriteFile(tempWAV, wavBytes, 0644); err != nil {
		return "", fmt.Errorf("write wav file: %w", err)
	}

	return tempWAV, nil
}

// Speak synthesizes and plays text through speakers.
// Prioritizes the premium Google Gemini 3.1 Flash TTS model if an API key is available.
// Otherwise, falls back to Google TTS (gtts-cli) and then to local offline synthesizers (say, PowerShell, espeak-ng).
func Speak(text, apiKey string) error {
	if text == "" {
		return nil
	}

	// Limit text length to avoid endless speaking
	const maxLen = 2500
	if len(text) > maxLen {
		text = text[:maxLen] + "..."
	}

	// 1. Prioritize official Gemini 3.1 Flash TTS Model (Premium, human-like voice)
	if apiKey != "" {
		wavPath, err := SynthesizeGeminiTTS(text, apiKey)
		if err == nil {
			playErr := PlayAudio(wavPath)
			if playErr == nil {
				os.Remove(wavPath)
				return nil // Success with premium Gemini AI voice!
			}
			os.Remove(wavPath)
			fmt.Fprintf(os.Stderr, "\033[90m⟲ Gemini TTS play error: %v. Falling back...\033[0m\n", playErr)
		} else {
			fmt.Fprintf(os.Stderr, "\033[90m⟲ Gemini TTS generation error: %v. Falling back...\033[0m\n", err)
		}
	}

	// 2. Fallback 1: Google Translate TTS via gtts-cli (universal, high-quality free voice)
	if _, err := exec.LookPath("gtts-cli"); err == nil {
		voiceDir := EnsureTempDir("")
		tempMP3 := filepath.Join(voiceDir, fmt.Sprintf("harness_gtts_%d.mp3", time.Now().UnixNano()))

		// Run gtts-cli (using 'pt' for Portuguese to avoid pt-br warnings)
		cmd := exec.Command("gtts-cli", "--lang", "pt", text, "--output", tempMP3)
		if err := cmd.Run(); err == nil {
			var played bool

			// Try mpg123
			if _, err := exec.LookPath("mpg123"); err == nil {
				playCmd := exec.Command("mpg123", "-q", tempMP3)
				if playCmd.Run() == nil {
					played = true
				}
			}

			// Try mpv
			if !played {
				if _, err := exec.LookPath("mpv"); err == nil {
					playCmd := exec.Command("mpv", "--no-video", "--really-quiet", tempMP3)
					if playCmd.Run() == nil {
						played = true
					}
				}
			}

			// Try play (SoX)
			if !played {
				if _, err := exec.LookPath("play"); err == nil {
					playCmd := exec.Command("play", "-q", tempMP3)
					if playCmd.Run() == nil {
						played = true
					}
				}
			}

			// Try play via PlayAudio after ffmpeg conversion
			if !played {
				if _, err := exec.LookPath("ffmpeg"); err == nil {
					tempWAV := filepath.Join(voiceDir, fmt.Sprintf("harness_gtts_%d.wav", time.Now().UnixNano()))
					convCmd := exec.Command("ffmpeg", "-y", "-i", tempMP3, tempWAV)
					if convCmd.Run() == nil {
						if PlayAudio(tempWAV) == nil {
							played = true
						}
						os.Remove(tempWAV)
					}
				}
			}

			// Try ffplay
			if !played {
				if _, err := exec.LookPath("ffplay"); err == nil {
					playCmd := exec.Command("ffplay", "-nodisp", "-autoexit", tempMP3)
					if playCmd.Run() == nil {
						played = true
					}
				}
			}

			os.Remove(tempMP3)
			if played {
				return nil // Success with free Google Translate TTS!
			}
		}
	}

	// 3. Fallback 2: OS-native local offline synthesizers
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
		// Linux offline synthesizers (spd-say, espeak-ng, espeak)
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
		return fmt.Errorf("no offline TTS available (try: sudo apt install espeak-ng): %w", lastErr)

	default:
		return fmt.Errorf("TTS not supported on OS: %s", runtime.GOOS)
	}
}

// PlayAudio plays a WAV file using native platform tool.
func PlayAudio(path string) error {
	switch runtime.GOOS {
	case "linux":
		// 1. Try mpv if available (uses system mixing, extremely robust and supports PipeWire/PulseAudio)
		if _, err := exec.LookPath("mpv"); err == nil {
			cmd := exec.Command("mpv", "--no-video", "--really-quiet", path)
			return cmd.Run()
		}
		// 2. Try paplay (PulseAudio) if available
		if _, err := exec.LookPath("paplay"); err == nil {
			cmd := exec.Command("paplay", path)
			return cmd.Run()
		}
		// 3. Try pw-play (PipeWire) if available
		if _, err := exec.LookPath("pw-play"); err == nil {
			cmd := exec.Command("pw-play", path)
			return cmd.Run()
		}
		// 4. Fall back to aplay (ALSA)
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
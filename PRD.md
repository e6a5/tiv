# 📄 PRD: Terminal Image Viewer (TIV)
## 🧭 Overview
Terminal Image Viewer (TIV) converts images to text output suitable for terminal display. It reads image files and outputs text that can be piped, redirected, or displayed directly. Following Unix philosophy: do one thing well - convert images to terminal-friendly text.

## 🎯 Goals
Convert images to text (ASCII art)

Work as a Unix filter (stdin/stdout)

Be predictable and composable

Keep it simple and portable

## 🧑‍🎓 Target Audience
Developers and terminal users

Linux/macOS users with terminal proficiency

Hobbyists learning image processing or TUI development

## 🛠️ Features
### ✅ Core Features
Feature	Description
Read images	PNG, JPEG from file or stdin
Convert to ASCII	Grayscale conversion using printable characters
Output text	Plain text to stdout (pipeable)
Size control	Simple width/height options
Unix filter	Works with pipes: `cat image.jpg | tiv | less`

### 🖼️ Output Format
**ASCII Text Only**
- Converts pixels to ASCII characters: ` .-+*#@`
- Grayscale based on pixel brightness
- Plain text output (no escape codes)
- Portable across all terminals and systems
- Can be saved, piped, or processed by other tools

**Unix Philosophy:** One tool, one job, done well.

## 🖥️ Technical Requirements
### Language
Go (Golang), for performance and easy binary distribution

### Image Processing
Use Go’s image, image/color packages

Resize with github.com/nfnt/resize or golang.org/x/image/draw

### Input/Output
**Input:** Image files (PNG, JPEG) or stdin
**Output:** Plain ASCII text to stdout

No terminal detection needed - output is always portable text

## 🔧 CLI Options
Flag	Description
-w, --width	Output width in characters (default: 80)
-h, --height	Output height in characters (auto-calculated if not set)
-i, --invert	Invert brightness levels
--help	Show usage

Example usage:
```bash
tiv image.jpg                    # Convert to 80-char width ASCII
tiv -w 40 image.png             # 40 characters wide
cat image.jpg | tiv             # Read from stdin
tiv image.jpg > ascii.txt       # Save to file
tiv image.jpg | head -20        # Preview first 20 lines
```

## 🧪 Testing Plan
Load different image formats (PNG, JPG, WebP)

Resize on various terminal sizes

Check fallback works: e.g., test on xterm vs Kitty

Confirm color accuracy using test gradients

Error handling: corrupted files, unsupported formats, very large images

## 🎯 Success Criteria
Works in 95% of modern terminals

Loads common image formats (PNG, JPEG) reliably

Processes images up to 10MB in under 3 seconds

Readable output on terminals as small as 80x24

Graceful error messages for invalid inputs

## 🧩 Future Enhancements
**Separate tools following Unix philosophy:**

`tiv-color` - ANSI color version (separate binary)

`tiv-gif` - Animated GIF support (separate binary)  

`tiv-web` - Remote image fetcher: `curl image.jpg | tiv`

Each tool does one thing well, can be combined with pipes

## 📦 Output
A simple filter that:

Reads images from files or stdin

Converts to ASCII text

Outputs to stdout

Works everywhere - from embedded systems to modern desktops

## 🚀 Distribution
Single binary releases for Linux, macOS, Windows

Optional: Homebrew formula for easy installation

Built with `go build` - no external dependencies

## 🐧 Unix Philosophy Compliance
**"Do one thing and do it well"** - Convert images to ASCII text, nothing more

**"Write programs to handle text streams"** - Reads stdin, writes stdout  

**"Write programs that work together"** - Pipeable with other Unix tools

**"Small is beautiful"** - Simple interface, minimal options

**"Choose portability over efficiency"** - Plain ASCII output works everywhere


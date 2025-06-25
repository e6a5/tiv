# 🎨 TIV - Terminal Image Viewer

[![CI](https://github.com/e6a5/tiv/workflows/CI/badge.svg)](https://github.com/e6a5/tiv/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/e6a5/tiv)](https://goreportcard.com/report/github.com/e6a5/tiv)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A fast, powerful terminal image viewer that displays images **side-by-side** with ASCII art. Features ANSI color support, Unicode blocks for 2x resolution, Floyd-Steinberg dithering, and inline image display. Perfect for viewing images in terminals with both original and ASCII representations simultaneously.

## 🐧 Philosophy

**Do one thing and do it well:** Display images in terminals with maximum clarity.

TIV follows Unix philosophy while providing modern terminal image viewing. By default, it displays images side-by-side with ASCII conversions for the best of both worlds. For traditional workflows, use `-no-split` to output pure ASCII text compatible with pipes and redirection.

## 📦 Installation

### Pre-built Binaries (Recommended)

Download from [releases page](https://github.com/e6a5/tiv/releases):

```bash
# Linux (x64)
wget https://github.com/e6a5/tiv/releases/latest/download/tiv-linux-amd64.tar.gz
tar -xzf tiv-linux-amd64.tar.gz
sudo mv tiv-linux-amd64 /usr/local/bin/tiv

# macOS (Homebrew)
brew install e6a5/tap/tiv

# macOS (Manual)
wget https://github.com/e6a5/tiv/releases/latest/download/tiv-darwin-arm64.tar.gz
tar -xzf tiv-darwin-arm64.tar.gz
sudo mv tiv-darwin-arm64 /usr/local/bin/tiv
```

### From Source

```bash
go install github.com/e6a5/tiv@latest
```

### Development

```bash
git clone https://github.com/e6a5/tiv.git
cd tiv
make build
make install  # Optional: install system-wide
```

## Usage

```bash
# Split view (default): image + ASCII side-by-side
tiv image.jpg

# Classic ASCII only
tiv -no-split image.jpg

# Pipe mode (automatically ASCII-only)  
cat image.jpg | tiv

# Save ASCII to file
tiv -no-split image.jpg > ascii.txt

# Preview with pager
tiv image.jpg | less

# Invert brightness
tiv -i image.jpg

# Increase contrast for better detail
tiv -c 1.5 image.jpg

# 🌟 High resolution Unicode blocks
tiv -b image.jpg

# 🎨 Professional dithering for smooth gradients
tiv -d image.jpg

# 🌈 ANSI Color Support
tiv -color 256 image.jpg     # 256-color mode
tiv -color 24bit image.jpg   # 24-bit truecolor

# 🚀 Ultimate quality: blocks + dithering + color + contrast
tiv -b -d -color 24bit -c 1.3 image.jpg

# 🖼️  Show actual image in terminal (instead of ASCII)
tiv -p image.jpg                             # Auto-detect best preview method
tiv -p -preview-mode terminal image.jpg      # Force terminal inline preview
tiv -p -preview-mode system image.jpg        # Force system viewer

# Pipeline example
tiv image.jpg | head -20 | tail -10
```

## Options

- `-w, --width`: Output width in characters (default: 80)
- `-h, --height`: Output height in characters (auto-calculated if not set)
- `-i, --invert`: Invert brightness levels
- `-c, --contrast`: Contrast adjustment (0.5-2.0, default: 1.0)
- `-b, --blocks`: 🌟 Use Unicode block characters for **2x higher resolution**
- `-d, --dither`: 🎨 Apply Floyd-Steinberg dithering for **professional quality**
- `--color`: 🌈 **ANSI color output** ('256' for 256-color, '24bit' for truecolor)
- `-p, --preview`: 👁️ **Show original image preview** (auto-detects best method)
- `--preview-mode`: Preview mode: 'auto', 'terminal', or 'system'
- `--no-split`: Disable split view (classic ASCII-only mode)
- `--help`: Show usage information

## Supported Formats

- PNG
- JPEG

## 🖼️ Terminal Image Preview

TIV can display the **actual image** directly in your terminal instead of ASCII art using modern terminal protocols:

### Auto Mode (Default)
```bash
tiv -p image.jpg
```
Automatically detects terminal capabilities and chooses the best preview method.

### Terminal Inline Preview
```bash
tiv -p -preview-mode terminal image.jpg
```
Shows the original image directly in compatible terminals using:
- **Kitty Graphics Protocol**: Kitty, Ghostty, Konsole
- **iTerm2 Inline Images**: iTerm2, WezTerm, Warp, VS Code, Tabby, Hyper, Bobcat  
- **Sixel Graphics**: foot, Windows Terminal, Black Box, xterm

### System Viewer
```bash
tiv -p -preview-mode system image.jpg
```
Opens the image in your system's default image viewer (Preview.app on macOS, etc.)

### Supported Terminals

Based on the [Yazi terminal compatibility research](https://yazi-rs.github.io/docs/image-preview), TIV supports:

| Terminal | Protocol | Status |
|----------|----------|--------|
| Kitty | Kitty Graphics | ✅ Full support |
| iTerm2 | Inline Images | ✅ Full support |
| WezTerm | Inline Images | ✅ Full support |
| Ghostty | Kitty Graphics | ✅ Full support |
| Windows Terminal | Sixel | ✅ Full support |
| VS Code | Inline Images | ⚠️ Limited support |
| foot | Sixel | ✅ Full support |
| Konsole | Kitty Graphics | ✅ Full support |
| Warp | Inline Images | ✅ Full support |
| Tabby | Inline Images | ✅ Full support |
| Hyper | Inline Images | ✅ Full support |
| Black Box | Sixel | ✅ Full support |

## Examples

Convert and save:
```bash
tiv photo.jpg > ascii-art.txt
```

Use in pipelines:
```bash
find . -name "*.jpg" | head -1 | xargs cat | tiv -w 60
```

Quick preview:
```bash
tiv image.png | head -20
```

Full color with high resolution:
```bash
tiv -b -d -color 24bit -c 1.3 photo.jpg
```

## 🤝 Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

- 🐛 [Report bugs](https://github.com/e6a5/tiv/issues)
- 💡 [Request features](https://github.com/e6a5/tiv/issues)
- 🔧 [Submit PRs](https://github.com/e6a5/tiv/pulls)

## 📝 License

[MIT License](LICENSE) - see the [LICENSE](LICENSE) file for details.

## 🎯 Unix Philosophy Compliance

✓ **Do one thing well**: Converts images to ASCII text  
✓ **Handle text streams**: Reads stdin, writes stdout  
✓ **Work with other programs**: Pipeable and composable  
✓ **Keep it simple**: Minimal interface and options  
✓ **Be portable**: Plain ASCII output works everywhere

## 🌟 Star History

[![Star History Chart](https://api.star-history.com/svg?repos=e6a5/tiv&type=Date)](https://star-history.com/#e6a5/tiv&Date) 
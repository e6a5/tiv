# 🎨 TIV - Terminal Image Viewer

[![CI](https://github.com/e6a5/tiv/workflows/CI/badge.svg)](https://github.com/e6a5/tiv/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/e6a5/tiv)](https://goreportcard.com/report/github.com/e6a5/tiv)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A fast, simple Unix filter that converts images to ASCII art. Perfect for viewing images in terminals, creating ASCII art, or integrating into shell workflows.

## 🐧 Philosophy

**Do one thing and do it well:** Convert images to ASCII text.

TIV follows Unix philosophy - it reads image files or stdin, converts pixels to ASCII characters based on brightness, and outputs plain text to stdout. Designed to work seamlessly with other Unix tools through pipes and redirection.

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
# Basic usage
tiv image.jpg

# Specify width
tiv -w 40 image.png

# Read from stdin
cat image.jpg | tiv

# Save to file
tiv image.jpg > ascii.txt

# Preview with pager
tiv image.jpg | less

# Invert brightness
tiv -i image.jpg

# Increase contrast for better detail
tiv -c 1.5 image.jpg

# Pipeline example
tiv image.jpg | head -20 | tail -10
```

## Options

- `-w, --width`: Output width in characters (default: 80)
- `-h, --height`: Output height in characters (auto-calculated if not set)
- `-i, --invert`: Invert brightness levels
- `-c, --contrast`: Contrast adjustment (0.5-2.0, default: 1.0)
- `--help`: Show usage information

## Supported Formats

- PNG
- JPEG

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
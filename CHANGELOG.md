# Changelog

All notable changes to TIV (Terminal Image Viewer) will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial release of TIV
- **🌈 ANSI Color Support**: Full 256-color and 24-bit truecolor output modes
- `--color` flag with support for '256' and '24bit'/'truecolor' modes
- Advanced RGB to ANSI color palette mapping for 256-color mode
- 24-bit truecolor support for maximum color fidelity
- Unicode block characters for 2x higher resolution (-b flag)
- Floyd-Steinberg dithering for professional gradients (-d flag)
- ASCII art conversion with 13-character density range
- Support for PNG and JPEG formats
- Unix filter design (stdin/stdout)
- Box filter sampling for better image quality
- Contrast adjustment (-c flag)
- Proper aspect ratio handling
- Cross-platform support (Linux, macOS, Windows)
- Version flag (--version)
- **📱 Split View (Default)**: Side-by-side display with original image on left, ASCII on right
  - Automatic terminal detection and best preview method selection
  - Graceful fallback to placeholder when inline images not supported
  - Screen-aware layout with proper positioning and sizing
- **🖼️ Terminal Inline Image Display**: Show actual images in terminal instead of ASCII
  - Kitty Graphics Protocol support (Kitty, Ghostty, Konsole)
  - iTerm2 Inline Images support (iTerm2, WezTerm, Warp, VS Code, Tabby, Hyper, Bobcat)
  - Sixel Graphics support (foot, Windows Terminal, Black Box, xterm)
- **🔄 Smart Preview Auto-Detection**: Automatically chooses best preview method
- **⚙️ Preview Mode Control**: `--preview-mode` flag with 'auto', 'terminal', 'system' options
- **📱 Multi-Protocol Support**: Based on Yazi terminal compatibility research
- **🎛️ Classic Mode**: `--no-split` flag for traditional ASCII-only output
- Cross-platform system viewer fallback (macOS `open`, Linux `xdg-open`/`eog`/`feh`, Windows `start`)
- Pure image display mode (no ASCII conversion when using preview)

### Features
- `-w, --width`: Set output width in characters
- `-h, --height`: Set output height in characters  
- `-i, --invert`: Invert brightness levels
- `-c, --contrast`: Adjust contrast (0.5-2.0 range)
- `-b, --blocks`: Use Unicode block characters for 2x higher resolution
- `-d, --dither`: Apply Floyd-Steinberg dithering for professional quality
- `--color`: Enable ANSI color output ('256' for 256-color, '24bit' for truecolor)
- `-p, --preview`: Show original image preview (auto-detects best method)
- `--preview-mode`: Control preview mode ('auto', 'terminal', 'system')
- `--no-split`: Disable split view (classic ASCII-only mode)

### Unix Philosophy Compliance
- Do one thing well: Convert images to ASCII
- Handle text streams: stdin → stdout
- Work with other programs: Fully pipeable
- Simple interface: Minimal, intuitive flags
- Portable: Plain ASCII output

## [1.0.0] - 2024-XX-XX

### Added
- Initial stable release
- Core image to ASCII conversion functionality
- Standard Unix tool behavior
- Comprehensive documentation
- Automated builds and releases

[Unreleased]: https://github.com/e6a5/tiv/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/e6a5/tiv/releases/tag/v1.0.0 
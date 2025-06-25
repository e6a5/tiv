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

### Features
- `-w, --width`: Set output width in characters
- `-h, --height`: Set output height in characters  
- `-i, --invert`: Invert brightness levels
- `-c, --contrast`: Adjust contrast (0.5-2.0 range)
- `-b, --blocks`: Use Unicode block characters for 2x higher resolution
- `-d, --dither`: Apply Floyd-Steinberg dithering for professional quality
- `--color`: Enable ANSI color output ('256' for 256-color, '24bit' for truecolor)

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
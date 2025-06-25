# Changelog

All notable changes to TIV (Terminal Image Viewer) will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial release of TIV
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
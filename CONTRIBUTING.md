# Contributing to TIV

Thank you for your interest in contributing to TIV (Terminal Image Viewer)! This document provides guidelines for contributing to the project.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Submitting Changes](#submitting-changes)
- [Issue Guidelines](#issue-guidelines)
- [Unix Philosophy](#unix-philosophy)

## 🤝 Code of Conduct

This project follows the [Contributor Covenant](https://www.contributor-covenant.org/) code of conduct. Please be respectful and inclusive in all interactions.

## 🚀 Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/your-username/tiv.git`
3. Create a feature branch: `git checkout -b feature-name`

**Note:** Replace `your-username` with your actual GitHub username in the clone command.

## 💻 Development Setup

### Prerequisites
- Go 1.21 or later
- Make (optional, for convenience)

### Building
```bash
# Quick development build
make dev

# Full build with optimizations
make build

# Multi-platform build
make build-all
```

### Testing
```bash
# Run tests
make test

# Manual testing
echo "test" | ./tiv
```

## 🛠️ Making Changes

### Code Style
- Follow standard Go formatting (`gofmt`)
- Use meaningful variable and function names
- Add comments for non-obvious code
- Keep functions small and focused

### Adding Features
Before adding new features:

1. **Check if it aligns with Unix philosophy**
   - Does it maintain the "do one thing well" principle?
   - Does it preserve the simple interface?
   - Can it be composed with other Unix tools?

2. **Consider backwards compatibility**
   - New flags should be optional
   - Default behavior should remain unchanged
   - Output format should be stable

3. **Update documentation**
   - Update `README.md` if needed
   - Update `--help` text
   - Add examples

### Bug Fixes
- Include a test case that reproduces the issue
- Keep changes minimal and focused
- Verify the fix doesn't break existing functionality

## 📝 Submitting Changes

1. **Commit Guidelines**
   ```
   type: brief description (50 chars max)
   
   Longer explanation if needed (72 chars per line)
   
   Examples:
   feat: add contrast adjustment flag
   fix: handle zero-dimension images
   docs: update installation instructions
   ```

2. **Pull Request Process**
   - Create a clear PR title and description
   - Reference any related issues
   - Include before/after examples for UI changes
   - Ensure all tests pass
   - Verify builds work on your platform

3. **Review Process**
   - Maintainers will review within a few days
   - Address feedback promptly
   - Be patient and respectful during review

## 🐛 Issue Guidelines

### Bug Reports
Include:
- Go version (`go version`)
- Operating system and terminal
- Input image format and size
- Command used
- Expected vs actual behavior
- Sample image (if safe to share)

### Feature Requests
Consider:
- Does it fit the Unix philosophy?
- Is it generally useful?
- Can it be implemented simply?
- Would it be better as a separate tool?

## 🐧 Unix Philosophy

TIV follows Unix philosophy principles:

- **Do one thing well**: Convert images to ASCII text
- **Work with other programs**: Be pipeable and composable
- **Handle text streams**: Process stdin/stdout cleanly
- **Be simple**: Minimal interface, predictable behavior
- **Be portable**: Work across different systems

When contributing, please ensure changes maintain these principles.

## 📚 Development Resources

- [Go Documentation](https://golang.org/doc/)
- [Unix Philosophy](https://en.wikipedia.org/wiki/Unix_philosophy)
- [Image Processing in Go](https://golang.org/pkg/image/)

## 💡 Ideas for Contributions

Good first contributions:
- Improve error messages
- Add input validation
- Optimize image processing
- Add more image format support
- Improve documentation
- Write tests

Advanced contributions:
- Performance optimizations
- Memory usage improvements
- New rendering algorithms
- Platform-specific enhancements

## ❓ Questions?

Feel free to open an issue for questions or discussions about:
- Implementation approaches
- Feature ideas
- Technical decisions
- Project direction

Thank you for contributing to TIV! 🎨 
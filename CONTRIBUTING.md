# Contributing to GitSentry 🤝

Thank you for your interest in contributing to GitSentry! We welcome contributions from developers of all skill levels.

## 🚀 Quick Start for Contributors

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/jinu721/git-sentry
   cd gitsentry
   ```
3. **Set up development environment**:
   ```bash
   go mod tidy
   make build
   ```

## 🐛 Reporting Issues

### Before Submitting an Issue
- Check if the issue already exists in [GitHub Issues](https://github.com/jinu721/git-sentry/issues)
- Try the latest version to see if the issue is already fixed

### When Submitting an Issue
Please include:
- **Operating System** (Windows 10, macOS 13, Ubuntu 22.04, etc.)
- **Go Version** (`go version`)
- **GitSentry Version** (`gitsentry --version`)
- **Steps to reproduce** the issue
- **Expected behavior** vs **actual behavior**
- **Error messages** or logs (check `.gitsentry/logs/gitsentry.log`)

## 💡 Suggesting Features

We love new ideas! Before suggesting a feature:

1. **Check existing discussions** in [GitHub Discussions](https://github.com/jinu721/git-sentry/gitsentry/discussions)
2. **Explain the problem** your feature would solve
3. **Describe the solution** you have in mind
4. **Consider alternatives** and why your approach is best

## 🔧 Code Contributions

### Development Workflow

1. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following our coding standards

3. **Test your changes**:
   ```bash
   make test
   make build
   ./bin/gitsentry init  # Test locally
   ```

4. **Commit with clear messages**:
   ```bash
   git commit -m "feat: add new monitoring rule for large files"
   ```

5. **Push and create a Pull Request**

### Coding Standards

- **Go formatting**: Use `go fmt` and `go vet`
- **Error handling**: Always handle errors appropriately
- **Documentation**: Add comments for public functions and complex logic
- **Testing**: Add tests for new features and bug fixes

### Commit Message Format

We follow [Conventional Commits](https://www.conventionalcommits.org/):

```
type(scope): description

[optional body]

[optional footer]
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

**Examples:**
```
feat(monitor): add support for ignoring custom file patterns
fix(cli): resolve status command crash on Windows
docs(readme): update installation instructions
```

## 🧪 Testing

### Running Tests
```bash
make test
```

### Manual Testing
```bash
# Build and test locally
make build
./bin/gitsentry init
./bin/gitsentry start
# Make some file changes and verify suggestions work
```

### Testing on Different Platforms
We appreciate testing on:
- Windows (PowerShell, Command Prompt)
- macOS (Terminal, iTerm2)
- Linux (various distributions)

## 📚 Documentation

Help us improve documentation:

- **README.md**: Main project documentation
- **Code comments**: Explain complex logic
- **Help text**: CLI command descriptions
- **Examples**: Real-world usage scenarios

## 🎯 Areas Where We Need Help

### High Priority
- **Windows compatibility** improvements
- **Performance optimization** for large repositories
- **Configuration validation** and better error messages
- **Unit tests** for core functionality

### Medium Priority
- **Commit message templates** and customization
- **Integration with popular editors** (VS Code, Vim, etc.)
- **Localization** (multiple languages)
- **Advanced Git hooks** integration

### Low Priority
- **Web dashboard** for statistics
- **Team collaboration** features
- **Plugin system** for extensibility

## 🏗️ Architecture Overview

```
GitSentry Architecture:
├── CLI Layer (cmd/gitsentry, internal/cli)
├── Core Logic (internal/core)
├── Configuration (internal/config)
├── State Management (internal/state)
├── Git Operations (internal/git)
├── File Monitoring (internal/monitor)
└── Logging (internal/logger)
```

### Key Components

- **CLI**: Command-line interface using Cobra
- **Core**: Main orchestration and business logic
- **Monitor**: File system watching with fsnotify
- **Git**: Git repository operations and status
- **Config**: YAML-based configuration management
- **State**: JSON-based state persistence

## 🤔 Questions?

- **General questions**: [GitHub Discussions](https://github.com/jinu721/git-sentry/discussions)
- **Bug reports**: [GitHub Issues](https://github.com/jinu721/git-sentry/issues)
- **Feature requests**: [GitHub Discussions](https://github.com/jinu721/git-sentry/discussions)

## 📜 Code of Conduct

We are committed to providing a welcoming and inclusive experience for everyone. Please be:

- **Respectful** of different viewpoints and experiences
- **Constructive** in feedback and discussions
- **Collaborative** in problem-solving
- **Patient** with newcomers and questions

## 🎉 Recognition

Contributors will be:
- **Listed** in our README acknowledgments
- **Mentioned** in release notes for significant contributions
- **Invited** to join our contributor team for ongoing contributors

---

Thank you for helping make GitSentry better for everyone! 🚀

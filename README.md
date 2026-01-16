# GitSentry 🛡️

> **Your Personal Git Mentor - Never Miss a Commit Again**

GitSentry is a lightweight, local-first Git assistant that helps developers maintain clean Git habits through intelligent monitoring and gentle suggestions. It watches your code changes in real-time and suggests the perfect moments to commit, without ever taking control away from you.

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey.svg)](https://github.com/yourusername/gitsentry)

---

## 🎯 **The Problem GitSentry Solves**

**Every developer faces these Git challenges:**

- 🤔 **"When should I commit?"** - Working for hours without committing, then creating massive commits
- 📝 **"What should I write in the commit message?"** - Staring at the blank commit message box
- 🔄 **"Did I forget to push?"** - Losing work because commits weren't backed up
- 📊 **"How much have I changed?"** - No visibility into current work progress
- 🎯 **"Am I following best practices?"** - Inconsistent commit patterns and messages

**GitSentry solves all of these by being your intelligent Git companion.**

---

## ✨ **What Makes GitSentry Special**

### 🧠 **Intelligent Monitoring**
- **Real-time file watching** - Knows exactly what you're changing
- **Smart thresholds** - Suggests commits based on files changed, lines modified, and time elapsed
- **Context-aware** - Understands your project structure and ignores build artifacts

### 🎯 **Perfect Timing**
- **Never interrupts** - Suggestions appear between your natural work breaks
- **Configurable rules** - Set your own thresholds for when to suggest commits
- **Respectful notifications** - Helpful hints, not annoying popups

### 🔒 **Privacy & Control**
- **100% local** - No cloud services, no data collection
- **You're in control** - Never auto-commits or auto-pushes
- **Secure** - Uses your existing Git authentication

### 🌍 **Developer-Friendly**
- **Global installation** - Install once, use in every project
- **Cross-platform** - Works on Windows, Linux, and macOS
- **Zero configuration** - Works out of the box with sensible defaults

---

## 🚀 **Quick Start**

### **1. Install GitSentry Globally**

#### **Linux/macOS:**
```bash
git clone https://github.com/jinu721/git-sentry
cd gitsentry
chmod +x install.sh
./install.sh
```

#### **Windows (PowerShell as Administrator):**
```powershell
git clone https://github.com/jinu721/git-sentry
cd gitsentry
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
.\install.ps1
```

#### **Using Go (All Platforms):**
```bash
git clone https://github.com/jinu721/git-sentry
cd gitsentry
go install ./cmd/gitsentry
```

### **2. Use in Any Project**

```bash
# Navigate to your project
cd ~/logichub-project

# Initialize GitSentry (creates .gitsentry/ folder)
gitsentry init

# Start monitoring
gitsentry start
```

### **3. Code Normally - GitSentry Watches**

```bash
# GitSentry runs in the background and suggests commits like:

💡 GitSentry suggests it's a good time to commit!
   Files changed: 4
   Lines changed: 127
   Time since last commit: 25 minutes
   Run 'git add . && git commit' when ready

# After several commits:
📤 GitSentry suggests pushing your commits for backup!
   Unpushed commits: 3
   Run 'git push' when ready
```

---

## 📖 **Complete Usage Guide**

### **Core Commands**

| Command | Description |
|---------|-------------|
| `gitsentry init` | Initialize monitoring in current project |
| `gitsentry start` | Start background monitoring |
| `gitsentry stop` | Stop monitoring |
| `gitsentry status` | View current statistics and repository info |
| `gitsentry config` | View/modify configuration settings |

### **Configuration**

GitSentry creates a `.gitsentry/config.yaml` file in each project:

```yaml
rules:
  max_files_changed: 5        # Suggest commit after N files
  max_lines_changed: 100      # Suggest commit after N lines  
  max_minutes_since_commit: 30 # Suggest commit after N minutes
  max_unpushed_commits: 3     # Suggest push after N commits

auto_suggest_commits: true    # Enable commit suggestions
auto_suggest_pushes: true     # Enable push suggestions
commit_message_format: "conventional" # conventional or simple
```

### **Working with Multiple Projects**

```bash
# Each project has independent settings
cd ~/project-1
gitsentry init && gitsentry start

cd ~/project-2  
gitsentry init && gitsentry start

# GitSentry tracks each project separately
```

---

## 🏗️ **How It Works**

1. **File System Monitoring** - Uses efficient file watchers to detect changes
2. **Rule Engine** - Applies configurable rules to determine suggestion timing
3. **Git Integration** - Reads Git status, commit history, and remote state
4. **Smart Filtering** - Ignores temporary files, build artifacts, and hidden directories
5. **Gentle Suggestions** - Provides helpful hints without interrupting your flow

---

## 🤝 **Contributing**

We welcome contributions! Here's how you can help:

### **🐛 Report Issues**
- Found a bug? [Open an issue](https://github.com/jinu721/git-sentry/issues)
- Include your OS, Go version, and steps to reproduce

### **💡 Suggest Features**
- Have an idea? [Start a discussion](https://github.com/jinu721/git-sentry/discussions)
- Explain the problem it solves and how it would work

### **🔧 Code Contributions**

1. **Fork the repository**
   ```bash
   git clone https://github.com/jinu721/git-sentry
   cd gitsentry
   ```

2. **Set up development environment**
   ```bash
   go mod tidy
   make build
   ```

3. **Make your changes**
   - Follow Go best practices
   - Add tests for new features
   - Update documentation

4. **Test your changes**
   ```bash
   make test
   make build
   ./bin/gitsentry init  # Test locally
   ```

5. **Submit a Pull Request**
   - Clear description of changes
   - Reference any related issues
   - Include tests and documentation updates

### **📚 Documentation**
- Improve README, code comments, or help text
- Add examples and use cases
- Translate to other languages

### **🧪 Testing**
- Test on different operating systems
- Try with various project types
- Report compatibility issues

---

## 🛠️ **Development**

### **Prerequisites**
- Go 1.21 or later
- Git

### **Build from Source**
```bash
# Clone and build
git clone https://github.com/jinu721/git-sentry
cd gitsentry
make build

# Run tests
make test

# Install locally
make install
```

### **Project Structure**
```
gitsentry/
├── cmd/gitsentry/           # Main application entry point
├── internal/
│   ├── cli/                 # CLI commands and interface
│   ├── core/                # Core GitSentry logic
│   ├── config/              # Configuration management
│   ├── state/               # State persistence
│   ├── git/                 # Git operations
│   ├── monitor/             # File system monitoring
│   └── logger/              # Logging utilities
├── install.sh               # Unix installation script
├── install.ps1              # Windows installation script
├── uninstall.sh             # Unix removal script
├── uninstall.ps1            # Windows removal script
└── Makefile                 # Build automation
```

---

## 🆘 **Troubleshooting**

### **Installation Issues**
- **Command not found**: Ensure Go's bin directory is in your PATH
- **Permission denied**: Use `sudo` on Unix or run PowerShell as Administrator on Windows
- **Build fails**: Check Go version (requires 1.21+)

### **Runtime Issues**
- **No suggestions**: Check `gitsentry config` and ensure monitoring is active
- **File monitoring not working**: Verify file permissions and antivirus settings
- **Git not detected**: Ensure you're in a Git repository

### **Platform-Specific**
- **Windows**: Add `%GOPATH%\bin` to your PATH environment variable
- **Linux/macOS**: Add `export PATH=$PATH:$(go env GOPATH)/bin` to your shell profile

---

## 📄 **License**

MIT License - see [LICENSE](LICENSE) file for details.

---

## 🙏 **Acknowledgments**

- Built with [Cobra](https://github.com/spf13/cobra) for CLI interface
- File monitoring powered by [fsnotify](https://github.com/fsnotify/fsnotify)
- Inspired by the need for better Git habits in development teams

---

## 🌟 **Star History**

If GitSentry helps you maintain better Git habits, please consider giving it a star! ⭐

---

**GitSentry** - Because good Git habits shouldn't be hard to maintain! 🚀

*Made with ❤️ for developers who want to write better Git history*

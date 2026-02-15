# GitSentry

> **Your Intelligent Git Workflow Assistant - Professional Git Habit Management**

GitSentry is a lightweight, local-first Git assistant that helps developers maintain clean Git habits through intelligent monitoring and smart suggestions. It watches your code changes in real-time and suggests optimal moments to commit, without ever taking control away from you.

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey.svg)](https://github.com/jinu721/git-sentry)

---

## **The Problem GitSentry Solves**

**Every developer faces these Git challenges:**

- **"When should I commit?"** - Working for hours without committing, then creating massive commits
- **"What should I write in the commit message?"** - Staring at the blank commit message box
- **"Did I forget to push?"** - Losing work because commits weren't backed up
- **"How much have I changed?"** - No visibility into current work progress
- **"Am I following best practices?"** - Inconsistent commit patterns and messages

**GitSentry solves all of these by being your intelligent Git companion.**

---

## **What Makes GitSentry Special**

### **Intelligent Monitoring**
- **Real-time file watching** - Knows exactly what you're changing
- **Smart thresholds** - Suggests commits based on files changed, lines modified, and time elapsed
- **Context-aware** - Understands your project structure and ignores build artifacts

### **Perfect Timing**
- **Never interrupts** - Suggestions appear between your natural work breaks
- **Configurable rules** - Set your own thresholds for when to suggest commits
- **Respectful notifications** - Helpful hints, not annoying popups

### **Security & Privacy**
- **100% local** - No cloud services, no data collection
- **You're in control** - Never auto-commits or auto-pushes
- **Secure** - Uses your existing Git authentication with validated commands

### **Developer-Friendly**
- **Global installation** - Install once, use in every project
- **Cross-platform** - Works on Windows, Linux, and macOS
- **Zero configuration** - Works out of the box with sensible defaults

---

## **Quick Start**

### **1. Install GitSentry Globally**

#### **Linux/macOS:**
```bash
git clone https://github.com/jinu721/git-sentry.git
cd git-sentry
chmod +x install.sh
./install.sh
```

#### **Windows (PowerShell as Administrator):**
```powershell
git clone https://github.com/jinu721/git-sentry.git
cd git-sentry
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
.\install.ps1
```

#### **Using Go (All Platforms):**
```bash
git clone https://github.com/jinu721/git-sentry.git
cd git-sentry
go install ./cmd/gitsentry
```

### **2. Use in Any Project**

```bash
# Navigate to your project
cd ~/logichub

# Initialize GitSentry with team template
gitsentry init --template=team

# Start monitoring in background
gitsentry start --daemon
```

### **3. Code Normally - GitSentry Watches**

```bash
# GitSentry runs in the background and suggests commits like:

GitSentry suggests it's a good time to commit!
   Files changed: 4
   Lines changed: 127
   Time since last commit: 25 minutes
   Run 'git add . && git commit' when ready

# After several commits:
GitSentry suggests pushing your commits for backup!
   Unpushed commits: 3
   Run 'git push' when ready
```

---

## **Complete Usage Guide**

### **Core Commands**

| Command | Description |
|---------|-------------|
| `gitsentry init [--template=TYPE]` | Initialize monitoring in current project |
| `gitsentry start [--daemon]` | Start monitoring (foreground or background) |
| `gitsentry stop` | Stop monitoring |
| `gitsentry status` | View current statistics and repository info |
| `gitsentry rules [--interactive]` | View/modify configuration settings |
| `gitsentry stats [--export=json]` | Display or export statistics |
| `gitsentry doctor` | Run comprehensive diagnostics |

### **Configuration Templates**

GitSentry provides predefined templates for different workflows:

```bash
gitsentry init --template=default   # Balanced settings for individual developers
gitsentry init --template=team      # Stricter settings for team collaboration
gitsentry init --template=strict    # Very strict settings for critical projects
gitsentry init --template=relaxed   # Relaxed settings for experimental work
```

### **Interactive Configuration**

```bash
# Configure rules interactively
gitsentry rules --interactive

# Example interactive session:
Interactive Rules Configuration
==============================
Max files changed [3]: 5
Max lines changed [75]: 100
Max minutes since commit [20]: 30
Auto-suggest commits [true]: true
Auto-suggest pushes [true]: false
```

### **Statistics and Monitoring**

```bash
# View current status
gitsentry status

# Export statistics to JSON
gitsentry stats --export=json --output=stats.json

# Run health diagnostics
gitsentry doctor
```

### **Background Daemon Mode**

```bash
# Start as background daemon
gitsentry start --daemon

# Check if daemon is running
gitsentry status

# Stop daemon
gitsentry stop
```

---

## **Configuration**

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
gitsentry init --template=team && gitsentry start --daemon

cd ~/project-2  
gitsentry init --template=strict && gitsentry start --daemon

# GitSentry tracks each project separately
```

---

## **How It Works**

1. **File System Monitoring** - Uses efficient file watchers to detect changes
2. **Rule Engine** - Applies configurable rules to determine suggestion timing
3. **Git Integration** - Reads Git status, commit history, and remote state securely
4. **Smart Filtering** - Ignores temporary files, build artifacts, and hidden directories
5. **Gentle Suggestions** - Provides helpful hints without interrupting your flow

---

## **Security Features**

GitSentry prioritizes security and privacy:

- **Input Validation** - All file paths are sanitized to prevent directory traversal
- **Command Whitelisting** - Only safe Git commands are allowed
- **Secure File Operations** - All file operations use secure permissions
- **Thread Safety** - Concurrent operations are properly synchronized
- **No External Dependencies** - Works entirely offline with local Git

---

## **Contributing**

We welcome contributions! Here's how you can help:

### **Report Issues**
- Found a bug? [Open an issue](https://github.com/jinu721/git-sentry/issues)
- Include your OS, Go version, and steps to reproduce

### **Suggest Features**
- Have an idea? [Start a discussion](https://github.com/jinu721/git-sentry/discussions)
- Explain the problem it solves and how it would work

### **Code Contributions**

1. **Fork the repository**
   ```bash
   git clone https://github.com/jinu721/git-sentry.git
   cd git-sentry
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
   ./bin/gitsentry doctor  # Test locally
   ```

5. **Submit a Pull Request**
   - Clear description of changes
   - Reference any related issues
   - Include tests and documentation updates

---

## **Development**

### **Prerequisites**
- Go 1.21 or later
- Git

### **Build from Source**
```bash
# Clone and build
git clone https://github.com/jinu721/git-sentry.git
cd git-sentry
make build

# Run tests
make test

# Install locally
make install
```

### **Project Structure**
```
git-sentry/
├── cmd/gitsentry/           # Main application entry point
├── internal/
│   ├── cli/                 # CLI commands and interface
│   ├── core/                # Core GitSentry logic
│   ├── config/              # Configuration management
│   ├── state/               # State persistence
│   ├── git/                 # Git operations
│   ├── monitor/             # File system monitoring
│   ├── security/            # Security and validation
│   ├── daemon/              # Background process management
│   └── logger/              # Logging utilities
├── install.sh               # Unix installation script
├── install.ps1              # Windows installation script
├── uninstall.sh             # Unix removal script
├── uninstall.ps1            # Windows removal script
└── Makefile                 # Build automation
```

---

## **Troubleshooting**

### **Installation Issues**
- **Command not found**: Ensure Go's bin directory is in your PATH
- **Permission denied**: Use `sudo` on Unix or run PowerShell as Administrator on Windows
- **Build fails**: Check Go version (requires 1.21+)

### **Runtime Issues**
- **No suggestions**: Run `gitsentry doctor` to diagnose issues
- **File monitoring not working**: Verify file permissions and antivirus settings
- **Git not detected**: Ensure you're in a Git repository

### **Platform-Specific**
- **Windows**: Add `%GOPATH%\bin` to your PATH environment variable
- **Linux/macOS**: Add `export PATH=$PATH:$(go env GOPATH)/bin` to your shell profile

### **Diagnostics**
```bash
# Run comprehensive health check
gitsentry doctor

# Check current status
gitsentry status

# View configuration
gitsentry rules
```

---

## **License**

MIT License - see [LICENSE](LICENSE) file for details.

---

## **Acknowledgments**

- Built with [Cobra](https://github.com/spf13/cobra) for CLI interface
- File monitoring powered by [fsnotify](https://github.com/fsnotify/fsnotify)
- Inspired by the need for better Git habits in development teams

---

**GitSentry** - Because good Git habits shouldn't be hard to maintain.

*Made with care for developers who want to write better Git history*
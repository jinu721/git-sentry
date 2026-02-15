package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"gitsentry/internal/core"
)

var (
	interactiveMode bool
)

var rulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "Manage GitSentry rules configuration",
	Long:  `View and modify GitSentry monitoring rules interactively or display current settings.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sentry := core.NewGitSentry(".")
		
		if interactiveMode {
			return runInteractiveRules(sentry)
		}
		
		config, err := sentry.GetConfig()
		if err != nil {
			return fmt.Errorf("failed to get config: %w", err)
		}
		
		fmt.Println("Current GitSentry Rules")
		fmt.Println("======================")
		fmt.Printf("Max files changed: %d\n", config.Rules.MaxFilesChanged)
		fmt.Printf("Max lines changed: %d\n", config.Rules.MaxLinesChanged)
		fmt.Printf("Max minutes since commit: %d\n", config.Rules.MaxMinutesSinceCommit)
		fmt.Printf("Max unpushed commits: %d\n", config.Rules.MaxUnpushedCommits)
		fmt.Printf("Auto-suggest commits: %t\n", config.AutoSuggestCommits)
		fmt.Printf("Auto-suggest pushes: %t\n", config.AutoSuggestPushes)
		fmt.Printf("Commit message format: %s\n", config.CommitMessageFormat)
		
		return nil
	},
}

func runInteractiveRules(sentry *core.GitSentry) error {
	config, err := sentry.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to get config: %w", err)
	}
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Interactive Rules Configuration")
	fmt.Println("==============================")
	
	if newValue, err := promptInt(reader, "Max files changed", config.Rules.MaxFilesChanged); err == nil {
		config.Rules.MaxFilesChanged = newValue
	}
	
	if newValue, err := promptInt(reader, "Max lines changed", config.Rules.MaxLinesChanged); err == nil {
		config.Rules.MaxLinesChanged = newValue
	}
	
	if newValue, err := promptInt(reader, "Max minutes since commit", config.Rules.MaxMinutesSinceCommit); err == nil {
		config.Rules.MaxMinutesSinceCommit = newValue
	}
	
	if newValue, err := promptInt(reader, "Max unpushed commits", config.Rules.MaxUnpushedCommits); err == nil {
		config.Rules.MaxUnpushedCommits = newValue
	}
	
	if newValue, err := promptBool(reader, "Auto-suggest commits", config.AutoSuggestCommits); err == nil {
		config.AutoSuggestCommits = newValue
	}
	
	if newValue, err := promptBool(reader, "Auto-suggest pushes", config.AutoSuggestPushes); err == nil {
		config.AutoSuggestPushes = newValue
	}
	
	return sentry.SaveConfig(config)
}

func promptInt(reader *bufio.Reader, prompt string, current int) (int, error) {
	fmt.Printf("%s [%d]: ", prompt, current)
	
	input, err := reader.ReadString('\n')
	if err != nil {
		return current, err
	}
	
	input = strings.TrimSpace(input)
	if input == "" {
		return current, nil
	}
	
	return strconv.Atoi(input)
}

func promptBool(reader *bufio.Reader, prompt string, current bool) (bool, error) {
	currentStr := "false"
	if current {
		currentStr = "true"
	}
	
	fmt.Printf("%s [%s]: ", prompt, currentStr)
	
	input, err := reader.ReadString('\n')
	if err != nil {
		return current, err
	}
	
	input = strings.TrimSpace(strings.ToLower(input))
	if input == "" {
		return current, nil
	}
	
	return input == "true" || input == "yes" || input == "y", nil
}

func init() {
	rulesCmd.Flags().BoolVarP(&interactiveMode, "interactive", "i", false, "Interactive configuration mode")
}
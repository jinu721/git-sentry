package cli

import (
	"fmt"
	"os"
)

func PrintSuccess(message string) {
	fmt.Printf("✅ %s\n", message)
}

func PrintError(message string) {
	fmt.Fprintf(os.Stderr, "❌ Error: %s\n", message)
}

func PrintWarning(message string) {
	fmt.Printf("⚠️  Warning: %s\n", message)
}

func PrintInfo(message string) {
	fmt.Printf("ℹ️  %s\n", message)
}

func PrintHeader(title string) {
	fmt.Printf("\n🔧 %s\n", title)
	fmt.Println(generateSeparator(len(title) + 3))
}

func PrintSubHeader(title string) {
	fmt.Printf("\n📋 %s\n", title)
	fmt.Println(generateSeparator(len(title) + 3))
}

func generateSeparator(length int) string {
	separator := ""
	for i := 0; i < length; i++ {
		separator += "="
	}
	return separator
}

func FormatKeyValue(key, value string) string {
	return fmt.Sprintf("%-25s: %s", key, value)
}

func FormatBool(value bool) string {
	if value {
		return "✅ Yes"
	}
	return "❌ No"
}

func FormatStatus(isActive bool, activeMsg, inactiveMsg string) string {
	if isActive {
		return fmt.Sprintf("🟢 %s", activeMsg)
	}
	return fmt.Sprintf("🔴 %s", inactiveMsg)
}
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the previous command to the .remember file",
	Run:   Add,
}

func Add(cmd *cobra.Command, args []string) {
	// Get the current user's home directory
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return
	}
	zshHistoryPath := usr.HomeDir + "/.zsh_history"

	// Open the Zsh history file
	file, err := os.Open(zshHistoryPath)
	if err != nil {
		fmt.Println("Error opening Zsh history file:", err)
		return
	}
	defer file.Close()

	// Read the second-to-last command from the Zsh history file
	var lastCommand, secondToLastCommand string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			secondToLastCommand = lastCommand
			lastCommand = line
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading Zsh history file:", err)
		return
	}

	// Parse the second-to-last command (Zsh history lines may include timestamps or metadata)
	secondToLastCommand = parseZshHistoryLine(secondToLastCommand)

	// Get the root of the current Git repository
	gitRoot, err := getGitRoot()
	if err != nil {
		fmt.Println("Error determining Git repository root:", err)
		return
	}

	// Construct the path to the .remember file
	rememberFilePath := gitRoot + "/.remember"
	// Open the remember history file
	remFile, err := os.OpenFile(rememberFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening or creating .remember file:", err)
		return
	}
	defer remFile.Close()

	// Write the second-to-last command to the .remember file
	_, err = remFile.WriteString(secondToLastCommand + "\n")
	if err != nil {
		fmt.Println("Error writing to .remember file:", err)
		return
	}

	fmt.Printf("Command '%s' successfully added to .remember file", secondToLastCommand)
}

func parseZshHistoryLine(line string) string {
	// Zsh history lines may include metadata like `: 1681234560:0;command`
	// Split by `;` to extract the actual command
	parts := strings.SplitN(line, ";", 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return line
}

func getGitRoot() (string, error) {
	// Run the `git rev-parse --show-toplevel` command to get the root of the Git repository
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func init() {
	rootCmd.AddCommand(addCmd)
}

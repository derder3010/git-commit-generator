package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	numCommits := 10
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Invalid number of commits: %v", err)
		}
		numCommits = n
	}

	// startDate, err := time.Parse("2006-01-02", "2019-01-01")
	// if err != nil {
	// 	log.Fatalf("Invalid start date format: %v", err)
	// }
	// endDate, err := time.Parse("2006-01-02", "2024-12-16")
	// if err != nil {
	// 	log.Fatalf("Invalid end date format: %v", err)
	// }
	
	// if startDate.After(endDate) {
	// 	log.Fatal("Start date must be before end date")
	// }

	// if err := generateCommits(numCommits, startDate, endDate); err != nil {
	// 	log.Fatal(err)
	// }
	// Get the beginning of the current year *at runtime*
    now := time.Now()
    currentYear := now.Year() // Get the current year
    startDate := time.Date(currentYear, 1, 1, 0, 0, 0, 0, time.UTC) // Use currentYear

    // Get today's date *at runtime*
    endDate := time.Now()

    if err := generateCommits(numCommits, startDate, endDate); err != nil {
        log.Fatal(err)
    }
}

func generateCommits(numCommits int, startDate, endDate time.Time) error {
	for i := 0; i < numCommits; i++ {
		randomTime := randTime(startDate, endDate)
		commitMessage := randomTime.Format(time.RFC3339)

		filePath := filepath.Join(".", "data.txt")
		err := os.WriteFile(filePath, []byte(commitMessage+"\n"), 0644)
		if err != nil {
			return fmt.Errorf("writting file: %w", err)
		}

		if err := runGitCommand("add", "."); err != nil {
			return fmt.Errorf("git add %w", err)
		}

		if err := runGitCommand("commit", "-m", commitMessage, "--date", randomTime.Format(time.RFC3339)); err != nil {
			return fmt.Errorf("git commit: %w", err)
		}

		fmt.Println("Create commit:", commitMessage)
	}

	if err := runGitCommand("push"); err != nil {
		return fmt.Errorf("git push: %w", err)
	}
	return nil
}

func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("running git command: %w", err)
	}
	return nil
}

func randTime(min, max time.Time) time.Time {
	delta := max.Sub(min)
	sec := rand.Int63n(delta.Nanoseconds())
	return max.Add((time.Duration(sec)))
}
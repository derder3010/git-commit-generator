# Git Commit Generator (Go)

This is a simple command-line tool written in Go that generates a specified number of Git commits with random timestamps within a given date range. This can be useful for simulating repository activity, testing Git tools, or creating visualizations of commit history.

## Features

*   Generates multiple commits with random timestamps.
*   Configurable number of commits via command-line arguments.
*   Uses Go's built-in `time` package for date and time handling.
*   Robust error handling.
*   Clear output messages.

## Prerequisites

*   [Go](https://go.dev/dl/)
*   Git

## Installation

1.  **Clone the repository:**

    ```bash
    git clone [https://github.com/](https://github.com/)derder3010/git-commit-generator.git # Replace with your repo URL
    cd git-commit-generator
    ```

2.  **(Optional) Install Dependencies (if any):**
    If you use any external packages, run:
        ```bash
        go mod tidy
        ```

3.  **Build the project (optional):**
    If you want to create an executable, run:

    ```bash
    go build
    ```

## Usage

1.  **Initialize a Git repository (if you haven't already):**

    ```bash
    git init
    ```

2.  **(Optional) Create an initial commit:**

    ```bash
    touch README.md
    git add .
    git commit -m "Initial commit"
    ```

3.  **Run the generator:**

    *   To generate the default 10 commits:

        ```bash
        go run .
        ```

    *   To generate a specific number of commits (e.g., 50):

        ```bash
        go run . 50
        ```

    *   To run the built binary directly:

        ```bash
        ./git-commit-generator 100 # Example: 100 commits
        ```

4.  **(Optional) Push the commits to a remote repository:**

    ```bash
    git remote add origin <your-remote-repository-url>
    git push origin main # Or git push origin master
    ```

## Configuration (Date Range)

*   The start date is set to January 1, 2019, and the end date is the current time. To change the start date, modify this line in `main.go`:

```go
startDate := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
```
    *   Change 2019 to the desired year. The month and day can also be changed if needed. For example, for March 15, 2020:

```Go

startDate := time.Date(2020, 3, 15, 0, 0, 0, 0, time.UTC)
```

### Example

*   To generate 25 commits with random timestamps between January 1, 2019, and the current time, you would run:
```Bash

go run . 25
```
### Ethical Considerations
*   Please use this tool responsibly. Artificially inflating your contribution graph can be misleading and is generally discouraged. The contribution graph is intended to reflect genuine work and activity.
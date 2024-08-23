# GitHub User Activity CLI

![Go Report Card](https://goreportcard.com/badge/github.com/Ameprizzo/github-user-activity-cli)
[![GoDoc](https://godoc.org/github.com/Ameprizzo/github-user-activity-cli?status.svg)](https://pkg.go.dev/github.com/Ameprizzo/github-user-activity-cli)
![License](https://img.shields.io/github/license/Ameprizzo/github-user-activity-cli)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/Ameprizzo/github-user-activity-cli)
![GitHub issues](https://img.shields.io/github/issues/Ameprizzo/github-user-activity-cli)

A command-line tool built with Go to fetch and display a GitHub user's recent activity. This CLI application fetches data from the GitHub API and presents it in an organized, readable format in the terminal.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
  - [Fetching User Activity](#fetching-user-activity)
- [Output Example](#output-example)
- [Contributing](#contributing)
- [License](#license)
- [Project](#project)
- [Contact](#contact)

---

## Features

- **Fetch User Activity**: Retrieve recent activity data for any GitHub user.
- **Summarized Output**: Combines similar events for a cleaner display.
- **Supports Multiple Event Types**: Handles events like pushes, creation, and deletion of branches/tags, and more.

## Prerequisites

- **Go**: You need to have Go `go1.22.6` or above installed on your system.
  - **Installation**: Download and install Go from the official website: [golang.org/dl](https://golang.org/dl/).
  - **Verify Installation**: Run `go version` to ensure Go is properly installed.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Ameprizzo/github-user-activity-cli.git

# Go Presentation

## Setup
- Install Go from [here](https://go.dev/).
- Install the Go extension for VSCode.

## Coding
- Open a folder in VSCode.
- Create a module.
  - A module is a unit of distribution and versioning in Go.
  - Run `go mod init github.com/cmm7670/godemo`. 
  - Normally you'd want to initialize with the address that the module will be published to.
  - This will create a go.mod file in the root of the project.
- Create a Go package.
  - A package is a folder that contains Go code.
  - Create a folder named 'mascot' and a file inside it named 'mascot.go'.
  - Install any missing tools using the 'Analysis Tools Missing' prompt in the bottom toolbar.
    - It may take a minute.
  - Start coding in 'mascot.go'.
  - Create 'main.go'.
  - Run the code with `go run main.go`.
  - Create a test file named 'mascot_test.go' in the 'mascot' folder.
  
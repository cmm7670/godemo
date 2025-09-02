# Demo

## Setup
- Install Go from [here](https://go.dev/).
- Install the Go extension for VSCode.

## Create Modules
- Open Bash.
- Create a new folder in C:\code
  - `mkdir goMonthlyLearn`
- Change to the new folder and open VSCode.
  - `cd goMonthlyLearn`
  - `code .`
- Create a new module.
  - A module is a unit of distribution and versioning in Go.
  - `go mod init github.com/christophermayo-omes/goMonthlyLearn`
  - Normally you'd want to initialize with the address that the module will be published to.
  - This will create a go.mod file in the root of the project.
    - The go.mod file is used to define the module and its dependencies.
    - All outside modules which are needed for the project are listed here.
- Create a Go package.
  - A package is a folder that contains Go code.
  - Create a folder named returnString and a file inside it named returnString.go.
    - Add package declaration.
    - Add a function that returns a string.
  - Create a file named main.go in the root of the project.
    - Add package declaration.
    - Import the fmt package from the standard library.
    - Add a main function that calls the function from the returnString package.
    - Add the returnString package to the import for the main.go file
    - Run the code with `go run main.go`.
    - Add a second function call to the main() function that uses the quote package.
      - `fmt.Println(quote.Go())`
      - Save the file and it will add the quote package to the imports for main.go
      - Fix the error for the new package by running `go get`
      - Check the go.mod file to see the new dependency.
    - Run the code with `go run main.go`.

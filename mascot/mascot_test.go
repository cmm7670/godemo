// 6. Package declaration
package mascot_test

import (
	"testing"

	"github.com/cmm7670/godemo/mascot"
)

// 7. Create a TestMascot function
// VS Code adds 'run test' and 'debug test' options to the function.
// This test will fail until BestMascot() returns the string "Go Gopher".
func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot...")
	}
}

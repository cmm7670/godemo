// 3. Package declaration
package main

import (
	"fmt"

	"github.com/cmm7670/godemo/mascot"
	"rsc.io/quote"
)

// 4. Create a function
// After typing fmt. you can see the methods available in the fmt package.
// After selecting the method an import statement is automatically added.
// Pass in the BestMascot() function we created earlier.
// After typing the BestMascot() function name it is automatically added to the imports.
func main() {
	fmt.Println(mascot.BestMascot())
	// 5. Use a package from the internet.
	// This adds a new package to the imports, but it is underlined in red.
	// Use the VSCode quick fix to add the package to the mod file.
	fmt.Println(quote.Go())
}

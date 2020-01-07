
package main

import (
	"fmt"
  "os"
)
func main() {
   fmt.Println(os.Getenv("Build.SourceBranchName"))
   fmt.Println("-----")
   fmt.Println(os.Getenv("$(Build.SourceBranchName)"))
}

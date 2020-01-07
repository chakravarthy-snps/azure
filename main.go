
package main

import (
	"fmt"
         "os"
)
func main() {
   fmt.Println("printing SourceBranchName...")
   fmt.Println(os.Getenv("Build.SourceBranchName"))
   fmt.Println("-----")
   fmt.Println(os.Getenv("$Build.SourceBranchName"))
   fmt.Println("____________________")
   fmt.Println("printing SourceBranch...")
   fmt.Println(os.Getenv("Build.SourceBranch"))
   fmt.Println("-----")
   fmt.Println(os.Getenv("$Build.SourceBranch"))
   fmt.Println(os.Getenv("POLARIS_SERVER_URL"))
}

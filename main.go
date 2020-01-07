
package main

import (
	"fmt"
         "os"
)
func main() {
   fmt.Println("printing SourceBranchName...")
   fmt.Println(os.Getenv("BUILD.SOURCEBRANCHNAME"))
   fmt.Println("-----")
   fmt.Println(os.Getenv("$BUILD.SOURCEBRANCHNAME"))
   fmt.Println("____________________")
   fmt.Println("printing SourceBranch...")
   fmt.Println(os.Getenv("BUILD.SOURCEBRANCH"))
   fmt.Println("-----")
   fmt.Println(os.Getenv("$BUILD.SOURCEBRANCH"))
   fmt.Println(os.Getenv("POLARIS_SERVER_URL"))
}

// cmd/artificialcuratorstudio/main.go
package main

import (
"flag"
"log"
"os"

"artificialcuratorstudio/internal/artificialcuratorstudio"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialcuratorstudio.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}

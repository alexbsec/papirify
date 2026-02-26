package main

import (
	"fmt"
	"os"

	nodes "github.com/alexbsec/papirify/nodes"
)

func main() {
    p := tea.NewProgram(NewModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

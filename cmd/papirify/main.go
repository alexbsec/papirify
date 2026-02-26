package main

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/alexbsec/papirify/internal/nodes"
)

func main() {
	root := buildTree()

	p := tea.NewProgram(root)
	if _, err := p.Run(); err != nil {
		fmt.Printf("error running program: %v\n", err)
		os.Exit(1)
	}
}

func buildTree() *nodes.OptionNode {
	// ─── ROOT ──────────────────────────────────────────────
	rootOpts := []string{
		"Save snapshot",
		"Load snapshot",
		"Delete snapshot",
		"Exit",
	}

	root := nodes.NewOptionNode(
		"Papirify\nPreserve application state\n",
		rootOpts,
	)

	// ─── SAVE SNAPSHOT ─────────────────────────────────────
	saveOpts := []string{
		"Save to default location",
		"Save to custom location",
		"Back",
	}

	saveNode := nodes.NewOptionNode("Save Snapshot", saveOpts)

	// leaf: mock save default
	saveDefault := nodes.NewOptionNode(
		"Saving snapshot (default location)…",
		[]string{},
	)
	saveDefault.SetCommand(mockCommand("Snapshot saved to default location"))

	// leaf: mock save custom
	saveCustom := nodes.NewOptionNode(
		"Saving snapshot (custom location)…",
		[]string{},
	)
	saveCustom.SetCommand(mockCommand("Snapshot saved to custom location"))

	// wire save subtree
	saveNode.AddNext(0, saveDefault)
	saveNode.AddNext(1, saveCustom)
	saveNode.AddNext(2, root) // Back

	// ─── LOAD SNAPSHOT (mock) ──────────────────────────────
	loadNode := nodes.NewOptionNode(
		"Load Snapshot (mock)",
		[]string{"Load latest snapshot", "Back"},
	)

	loadLatest := nodes.NewOptionNode(
		"Loading snapshot…",
		[]string{},
	)
	loadLatest.SetCommand(mockCommand("Snapshot loaded successfully"))

	loadNode.AddNext(0, loadLatest)
	loadNode.AddNext(1, root)

	// ─── DELETE SNAPSHOT (mock) ────────────────────────────
	deleteNode := nodes.NewOptionNode(
		"Delete Snapshot (mock)",
		[]string{"Delete latest snapshot", "Back"},
	)

	deleteLatest := nodes.NewOptionNode(
		"Deleting snapshot…",
		[]string{},
	)
	deleteLatest.SetCommand(mockCommand("Snapshot deleted"))

	deleteNode.AddNext(0, deleteLatest)
	deleteNode.AddNext(1, root)

	// ─── EXIT ──────────────────────────────────────────────
	exitNode := nodes.NewOptionNode("Exiting Papirify…", []string{})
	exitNode.SetCommand(tea.Quit)

	// ─── WIRE ROOT ─────────────────────────────────────────
	root.AddNext(0, saveNode)
	root.AddNext(1, loadNode)
	root.AddNext(2, deleteNode)
	root.AddNext(3, exitNode)

	return root
}

// ─── MOCK COMMAND ─────────────────────────────────────────

func mockCommand(msg string) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(700 * time.Millisecond)
		return mockDoneMsg(msg)
	}
}

// message emitted by mock command
type mockDoneMsg string

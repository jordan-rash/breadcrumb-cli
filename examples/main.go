package main

import (
	"breadcrumb"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	root := breadcrumb.Crumb{
		Name:    "root",
		Display: "root-display",
		Parent:  nil,
		Help:    func() string { return "root help" },
		Action:  nil,
	}
	children := breadcrumb.Crumb{
		Name:    "children-container",
		Display: "children",
		Parent:  &root,
		Help:    func() string { return "children help" },
		Action:  nil,
	}
	child1 := breadcrumb.Crumb{
		Name:    "child1",
		Display: "child1-display",
		Parent:  &children,
		Options: []string{"smile"},
		State:   map[string]interface{}{"smile": "true"},
		Help:    func() string { return "child1 help" },
		Action: func(s map[string]interface{}) error {
			fmt.Println("from child1!")
			for k, v := range s {
				fmt.Printf("\tState item: %s -> %v\n", k, v)
			}
			return nil
		},
	}
	child2 := breadcrumb.Crumb{
		Name:    "child2",
		Display: "child2-display",
		Parent:  &children,
		Options: []string{"frown"},
		State:   map[string]interface{}{"frown": "true"},
		Help:    func() string { return "child2 help" },
		Action:  func(s map[string]interface{}) error { fmt.Println("from child2!"); return nil },
	}

	children.Children = []breadcrumb.Crumb{child1, child2}
	root.Children = []breadcrumb.Crumb{children}

	breadcrumb.Breadcrumb{
		Delimiter: " > ",
		Root:      root,
	}.Start()
}

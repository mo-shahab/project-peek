package main

import (
	"os"
	"fmt"
	"flag"
	"github.com/mo-shahab/project-peek/internal/tree"
)

func main() {

	dirPtr:= flag.String("d", ".", "Directory")
	showhidden := flag.Bool("a", false, "Show Hidden files")

	flag.Parse()

	rootpath := *dirPtr 
	counter := new (tree.Counter)
	dirTree, err := tree.BuildTree(rootpath, showhidden, counter)
	if err != nil {
		fmt.Printf("Error building directory tree: %v\n", err)
		return
	}

	fmt.Println("Directory Tree:")
	currdir := ""

	if rootpath == "." {
		currdir, _ = os.Getwd()
	} else {
		currdir = rootpath
	}

	fmt.Println(currdir)

	for i, child := range dirTree.Children {
		tree.PrintTree(child, "", i == len(dirTree.Children) -1)
	}
	
	fmt.Println(counter.PrintCount())
}

package main

import (
	"os"
	"fmt"
	"flag"
	"path/filepath"
	"strings"
)

type DirEntry struct {
	Name string
	Path string
	IsDir bool
	Children []*DirEntry
}

type Counter struct {
	dirs int
	files int
}

var skiplist =  []string {
	".git",
	"node_modules",
	"dist",
	".idea",
	".vscode",
	"venv",
} 

func (counter *Counter) count (isdir bool) {
	if isdir {
		counter.dirs += 1
	} else {
		counter.files += 1
	} 
}

func (counter *Counter) printCount() string {
	return fmt.Sprintf("\n%d directories, %d files", counter.dirs, counter.files)
}

func shouldskip(name string) bool {
	for _, skip := range skiplist {
		if name == skip {
			return true
		}
	}

	return false
}

func buildTree(rootpath string, showhidden * bool, counter *Counter) (*DirEntry, error) {
	info, err := os.Stat(rootpath)
	if err != nil {
		return nil, err
	}

	entry := &DirEntry {
		Name: info.Name(),
		Path: rootpath,
		IsDir: info.IsDir(),
	}

	counter.count(entry.IsDir)

	if entry.IsDir {
		if shouldskip(entry.Name) {
			return nil, nil
		}

		entries, err := os.ReadDir(rootpath)
		if err != nil {
			return nil, err
		}

		for _, e := range entries {
			name := e.Name()

			if !*showhidden && strings.HasPrefix(name, "."){
				continue
			}
			
			childpath := filepath.Join(rootpath, e.Name())
			childEntry, err := buildTree(childpath, showhidden, counter)

			if(err != nil) {
				fmt.Println("Error reading %s: %v\n", childpath, err)
				continue
			}
			entry.Children = append(entry.Children, childEntry)
		}
	}

	return entry, nil
}

func PrintTree(entry *DirEntry, prefix string, isLast bool) {
	if(entry == nil) {
		return
	}

	connector := "|--"
	
	if isLast {
		connector = "└──"
	}

	fmt.Printf("%s%s%s\n", prefix, connector, entry.Name)

	if entry.IsDir {
		newPrefix := prefix
		if isLast {
			newPrefix += "   "
		} else {
			newPrefix += "│   "
		}

		for i, child := range entry.Children {
            PrintTree(child, newPrefix, i == len(entry.Children)-1)
        }
	} 
}


func main() {

	dirPtr:= flag.String("d", ".", "Directory")
	showhidden := flag.Bool("a", false, "Show Hidden files")

	flag.Parse()

	rootpath := *dirPtr 
	counter := new (Counter)
	tree, err := buildTree(rootpath, showhidden, counter)
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

	for i, child := range tree.Children {
		PrintTree(child, "", i == len(tree.Children) -1)
	}
	
	fmt.Println(counter.printCount())
}

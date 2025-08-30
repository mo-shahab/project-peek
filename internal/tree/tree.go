package tree

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
)

type DirEntry struct {
	Name string
	Path string
	IsDir bool
	Children []*DirEntry
}

var Skiplist =  []string {
	".git",
	"node_modules",
	"dist",
	".idea",
	".vscode",
	"venv",
} 

func Shouldskip(name string) bool {
	for _, skip := range Skiplist {
		if name == skip {
			return true
		}
	}

	return false
}

func BuildTree(rootpath string, showhidden * bool, counter *Counter) (*DirEntry, error) {
	info, err := os.Stat(rootpath)
	if err != nil {
		return nil, err
	}

	entry := &DirEntry {
		Name: info.Name(),
		Path: rootpath,
		IsDir: info.IsDir(),
	}

	counter.Count(entry.IsDir)

	if entry.IsDir {
		if Shouldskip(entry.Name) {
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
			childEntry, err := BuildTree(childpath, showhidden, counter)

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


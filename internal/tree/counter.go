package tree

import (
	"fmt"
)

type Counter struct {
	dirs int
	files int
}

func (counter *Counter) Count (isdir bool) {
	if isdir {
		counter.dirs += 1
	} else {
		counter.files += 1
	} 
}

func (counter *Counter) PrintCount() string {
	return fmt.Sprintf("\n%d directories, %d files", counter.dirs, counter.files)
}

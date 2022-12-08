package solutions

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"github.com/coryjk/advent/internal/util/collections"
)

func Day7() {
	p := Problem{
		InputPath: "input/day7.txt",
	}
	var input []string = Input(p)
	fmt.Println(day7Part1(&input))
	fmt.Println(day7Part2(&input))
}

type file struct {
	name     string
	parent   *file
	children *list.List
	size     int
}

func day7Part1(input *[]string) string {
	// populate file structure based on stdout
	root := crawl(input)

	// find deletion dirs and their sum sizes
	dirsForDeletion, _ := findDirsForDeletion(root, 100000)
	size := 0
	for _, dirs := range *dirsForDeletion {
		size += dirs.size
	}
	return strconv.Itoa(size)
}

func day7Part2(input *[]string) string {
	root, requiredUnusedSpace, totalSpace := crawl(input), 30000000, 70000000

	// reuse previous function to just populate all file sizes
	findDirsForDeletion(root, 0)

	// flatten all dirs
	dirs, usedSpace := flattenDirs(root), root.size

	// current unused space
	unusedSpace := totalSpace - usedSpace

	// find the smallest single dir that may be deleted to acheive desired space
	var smallestDirToDelete *file = nil
	for _, dir := range *dirs {
		potentialUnusedSpace := unusedSpace + dir.size

		// if requirement met, replace min with current if it is smaller
		if potentialUnusedSpace >= requiredUnusedSpace {
			if smallestDirToDelete == nil || dir.size < smallestDirToDelete.size {
				smallestDirToDelete = dir
			}
		}
	}

	return strconv.Itoa(smallestDirToDelete.size)
}

func crawl(stdout *[]string) *file {
	// init root
	root := &file{
		name:     "/",
		parent:   nil,
		children: list.New(),
		size:     0,
	}

	current := root
	for i, line := range *stdout {
		// skip i == 0 since root already initialized
		if i == 0 {
			continue
		}
		split := strings.Split(line, " ")

		// indicates some command
		if split[0] == "$" {
			command := split[1]
			switch command {
			case "cd":
				// expecting 1 arg
				arg := split[2]
				switch arg {
				case "..":
					// navigate to parent
					current = current.parent
				case "/":
					// navigate to root
					current = root
				default:
					// navigate to child dir, assume it is populated by now
					target := childByName(arg, current.children)
					current = target
				}
			case "ls":
				// expecting no args...
			}
		} else {
			// indicates display of some information, need to further
			// populate file structure with this info

			// initialize dir and populate
			if split[0] == "dir" {
				current.children.PushBack(
					newFile(split[1], current, 0),
				)
			} else {
				// got file info, populate
				size, _ := strconv.Atoi(split[0])
				current.children.PushBack(
					newFile(split[1], current, size),
				)
				// only accumulate size 1 level up at dir level
				current.size += size
			}
		}
	}

	return root
}

func newFile(name string, parent *file, size int) *file {
	return &file{name, parent, list.New(), size}
}

func childByName(name string, files *list.List) *file {
	for i := files.Front(); i != nil; i = i.Next() {
		// list stores file pointers
		file := i.Value.(*file)
		if file.name == name {
			return file
		}
	}
	return nil
}

func findDirsForDeletion(current *file, max int) (*[]*file, int) {
	// file or empty dir
	if current.children.Len() == 0 {
		return nil, current.size
	} else {
		sum := 0
		var dirsForDeletion []*file

		// dir with children, recurse
		for i := current.children.Front(); i != nil; i = i.Next() {
			dirs, size := findDirsForDeletion(i.Value.(*file), max)

			// accumulate
			sum += size
			if dirs != nil {
				dirsForDeletion = append(dirsForDeletion, *dirs...)
			}
		}

		// if current dir less than max size, it is candidate for deletion
		current.size = sum
		if sum <= max {
			dirsForDeletion = append(dirsForDeletion, current)
		}
		return &dirsForDeletion, sum
	}
}

func flattenDirs(root *file) *[]*file {
	stack, usedSpace := collections.NewStack(), 0
	var dirs []*file

	// init stack with root
	stack.Push(root)

	for stack.Len() > 0 {
		current := stack.Pop().(*file)

		// assumption: only dirs have children
		for i := current.children.Front(); i != nil; i = i.Next() {
			node := i.Value.(*file)

			// continue iteration
			stack.Push(node)

			// only append if it is a dir (has children)
			if node.children.Len() > 0 {
				dirs = append(dirs, node)

				// size of files already accounted for in dirs
				usedSpace += node.size
			}
		}
	}

	return &dirs
}

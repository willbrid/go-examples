package main

import "fmt"

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}

	fmt.Println("Print folder 1")
	folder1.print(" ")

	cloneFolder1 := folder1.clone()
	fmt.Println("Print du clone folder1")
	cloneFolder1.print(" ")

	fmt.Println("Print folder 2")
	folder2.print(" ")

	cloneFolder2 := folder2.clone()
	fmt.Println("Print du clone folder2")
	cloneFolder2.print(" ")
}

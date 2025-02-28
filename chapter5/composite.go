package main

import "fmt"

type Entry interface {
	getName() string
	getSize() int
	remove()
}

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) getName() string {
	return f.name
}

func (f *File) getSize() int {
	return f.size
}

func (f *File) remove() {
	fmt.Printf("%sを削除しました\n", f.getName())
}

type Directory struct {
	name     string
	children []Entry
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:     name,
		children: []Entry{},
	}
}

func (d *Directory) getName() string {
	return d.name
}

func (d *Directory) getSize() int {
	size := 0
	for _, child := range d.children {
		size += child.getSize()
	}
	return size
}

func (d *Directory) remove() {
	for _, child := range d.children {
		child.remove()
	}
	fmt.Printf("%sを削除しました\n", d.getName())
}

func (d *Directory) add(child Entry) {
	d.children = append(d.children, child)
}

func client(entry Entry) {
	fmt.Println(entry.getName())
	fmt.Println(entry.getSize())
	entry.remove()
}

// func main() {
// 	dir1 := NewDirectory("design-pattern-part2")
// 	dir2 := NewDirectory("composite")
// 	file1 := NewFile("composite.ts", 100)
// 	file2 := NewFile("クラス図.png", 150)

// 	dir2.add(file1)
// 	dir2.add(file2)
// 	dir1.add(dir2)

// 	client(dir1)
// }

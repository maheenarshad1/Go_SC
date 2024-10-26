package main

import (
	"fmt"
)

func main() {
	// Folder to be moved (source folder)
	src := "lmn"

	// Existing destination folder
	dst := "ghi"

	// Create the new destination path (moving "abc" inside "ghi")
	newDst := "./" + dst + "/" + src

	fmt.Println(newDst)
	// // Move the folder by renaming it to the new destination
	// err := os.Rename(src, newDst)
	// if err != nil {
	// 	fmt.Println("Error moving folder:", err)
	// 	return
	// }

	// fmt.Println("Folder moved successfully!")
}

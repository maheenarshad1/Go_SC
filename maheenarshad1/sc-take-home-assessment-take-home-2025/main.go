package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)

	folder.PrettyPrint(res)
	fmt.Printf("\n Folders for orgID: %s", orgID)
	folder.PrettyPrint(orgFolder)

	// example usage for  Get All Child Folders function
	// children, err :=  folderDriver.GetAllChildFolders(orgID, "composed-wallflower")
	// folder.PrettyPrint(children)
	// fmt.Println("\n")
	// folder.PrettyPrint(err)

	// example usage for Move Folder function
	// NewPaths, err :=  folderDriver.MoveFolder("tight-titaness", "composed-wallflower")
	// folder.PrettyPrint(NewPaths)
	// fmt.Println("\n")
	// folder.PrettyPrint(err)
	
}

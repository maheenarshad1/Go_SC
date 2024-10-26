package folder

import (
	"errors"
	"strings"
)

// Logic Notes:
// Iterate through the folders to validate if the movement is possible; otherwise, generate relevant errors.
// If no errors are found, first update the path of the source folder by appending the destination folder's path,
// followed by a '.' and the source folder's name, modifying this directly in the folders data.
// For the children, grandchildren, and further descendants of the source folder, retrieve the path of the immediate parent
// (as it would have been updated earlier) and use that to append a '.' followed by the folder's name,
// then update the path directly in the folders data. - This works as a form of recursive update.

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	folders := f.folders
	if strings.TrimSpace(name) == strings.TrimSpace(dst) {
		return nil, errors.New("Error: Cannot move a folder to itself")
	}

	// Variables used to mark the validity of Src and Dst and the validity of the movement
	OrgSrc := ""
	OrgDst := ""
	// Assuming the Src and Dst name is invalid unless found in the folder
	InvalidSrc := true 
	InvalidDst := true 

	for _,f := range folders {
		if strings.TrimSpace(f.Name) == strings.TrimSpace(name) {
			OrgSrc = f.OrgId.String()
			InvalidSrc = false
		} else if strings.TrimSpace(f.Name) == strings.TrimSpace(dst) {
			OrgDst =f.OrgId.String()
			InvalidDst = false
		} 

		if strings.TrimSpace(f.Name) == strings.TrimSpace(dst) && strings.Contains(f.Paths, name) {
			Path := strings.Split(f.Paths, ".")
			sourceIndex := GetIndex(Path, name)
			destIndex := GetIndex(Path, dst)
			if sourceIndex < destIndex {
				return nil, errors.New("Error: Cannot move a folder to a child of itself")
			}
		}
	}

	// Generating Errors
	if InvalidSrc {
		return nil, errors.New("Error: Source folder does not exist")
	} else if InvalidDst {
		return nil, errors.New("Error: Destination folder does not exist")
	} else if  OrgSrc != OrgDst {
		return nil, errors.New("Error: Cannot move a folder to a different organization")
	} else {
	
	// If no errors are encountered up to this point, it means the folder movement is valid and possible
	for i := range folders {
		// Check if the current folder's path contains the source folder
		if strings.Contains(folders[i].Paths, name) {
			// Checking that this is the main folder of the source
				if folders[i].Name == name { 
						// Get the path of the destination folder
						DestPath := GetDestination(dst, folders) 
						// Append the name of the source folder to the path and change it in the Paths of the source
						folders[i].Paths = DestPath + "." + name 
				} else {
					// These are folders were the path contains the source folder name 
						PathComponent := strings.Split(folders[i].Paths, ".")
						// Validating just in case of errorenous paths as the path length is manipulated below
						if len(PathComponent) < 2 { 
								continue 
						}
						// Getting the second last folder in the path as that is the parent of the current folder
						parent := PathComponent[len(PathComponent)-2] 
						// Get the destination of the parent, which would have been updated already ( paths are being updated recursively)
						DestPath := GetDestination(parent, folders) 
						// Update the path of the current folder incorporting the changed path of the parent
						folders[i].Paths = DestPath + "." + folders[i].Name 
				}}
		}
}
	return folders, nil
}


// Function used to get the complete path of the destination 
// The complete path is then copied into the path of the source, and the source is added as the last folder 
func GetDestination(name string, list []Folder ) string {
	DestPath := ""
	for _,f := range list {
		if f.Name == name {
			DestPath = f.Paths
		}
}
	return DestPath
}

// function courtesy of https://www.willem.dev/code-snippets/index-of-value-in-slice/
func GetIndex(s []string, v string) int {
	for i, vs := range s {
		if vs == v {
			return i
		}
	}
	return -1
}
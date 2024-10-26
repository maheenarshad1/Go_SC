package folder

// Smaple comment to check recency

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

// Logic Notes:
// A folder will contain a '.' after its name, only if it appears somewhere within the path,
// but not if it is the last item in the path. Therefore, reflecting that it has children.
// Hence string manipulation is used to handle this distinction and pull out the children

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}
	return res
}


func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	
	folders := f.folders
	RelevantPaths := []Folder{}

	
	foldername := strings.TrimSpace(name) + "."

	// Variables used for error tracking
	OrgFound := false // Checking the validity of the OrgId
	NameFound := false // Checking the validity of the folder name 
	OrgnNameFound := false // Checking the validity of name associated with the OrgId

	for _, f := range folders {

		if f.OrgId == orgID {
			OrgFound = true
		}
		if strings.TrimSpace(f.Name) == strings.TrimSpace(name) {
			NameFound = true
		}
		if f.OrgId == orgID && strings.TrimSpace(f.Name) == strings.TrimSpace(name) {
			OrgnNameFound = true
		}
		if f.OrgId == orgID && strings.Contains(f.Paths, foldername) {
			RelevantPaths = append(RelevantPaths, f)
		}
	}

	if !OrgFound {
		return nil, errors.New("Error: Organization does not exist")
	} else if !NameFound {
		return nil, errors.New("Error: Folder does not exist")
	} else if !OrgnNameFound {
		return nil, errors.New("Error: Folder does not exist in the specified organization")
	}
	
	return RelevantPaths, nil
}
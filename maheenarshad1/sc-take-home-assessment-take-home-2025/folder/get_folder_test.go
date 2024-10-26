package folder_test

// Smaple comment to check recency
import (
	"errors"
	"strings"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

// Logic Notes:
// The path will always include the source folder's name followed by a '.' if it represents a child folder.
// If the source folder's name appears in the path without a trailing '.', it indicates the folder is the main source folder,
// and the source folder will appear as the last folder in the path (indicating no children in that path).

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	res := folder.GetAllFolders()
	f := folder.NewDriver(res)
	t.Parallel()
	tests := [...]struct {
		name    		string
		neededOrg 	string
		orgID   		uuid.UUID
		folders []folder.Folder
		Error    		error
	}{
		{
		name: "Checking organization to folder match error",
		neededOrg: "prompt-dynamite",
		orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
		Error: errors.New("Error: Folder does not exist in the specified organization"),
	},{
		name: "Checking folder error",
		neededOrg: "abc",
		orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
		Error:  errors.New("Error: Folder does not exist"),
	},{
		name: "No children exist",
		neededOrg: "equipped-hypno-hustler",
		orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
		Error:   nil,
	},
	{
		name: "children exit",
		neededOrg: "stunning-horridus",
		orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
		Error:   nil,
	},
	{
		name: "children exit",
		neededOrg: "fast-watchmen",
		orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
		Error:   nil,
	},
	{
		name: "children exit",
		neededOrg: "helped-blackheart",
		orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
		Error:   nil,
	},
	
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
		get, err := f.GetAllChildFolders(tt.orgID, tt.neededOrg)

		if tt.name == "Checking organization to folder match error" && err.Error() != tt.Error.Error() {
			t.Errorf("Test %s failed. Got Error: %v, Want: %v", tt.name, err, tt.Error)
		} else if tt.name == "Checking folder error" && err.Error() != tt.Error.Error(){
		t.Errorf("Test %s failed. Got Error: %v, Want: %v", tt.name, err, tt.Error)
		} else {
			// Generating the name of the form 'xyz.' and using that to check if it is not present in the paths,
			// of the returned folders, generate an error
			foldername := tt.neededOrg + "."
			for _,f := range tt.folders {
				if !strings.Contains(f.Paths, foldername){
					t.Errorf("Test %s failed. Got: %v, Want: %v", tt.name, get, tt.folders)
				}
			}
		}
		})
	}
}

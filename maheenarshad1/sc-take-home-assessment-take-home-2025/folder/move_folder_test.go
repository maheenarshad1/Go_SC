package folder_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

// Logic Notes:
// After a valid movement, the source folder will always have an index that is
// exactly one position greater than the destination folder in all paths
// containing the source folder. This logic is used to track errors related to
// valid movements. For any other cases, the error is compared to the expected
// result/error from the function call.
func Test_folder_MoveFolder(t *testing.T) {
	res := folder.GetAllFolders()
	f := folder.NewDriver(res)
	t.Parallel()
	tests := [...]struct {
		name    string
		src 		string
		dst			string
		orgID   uuid.UUID
		Error    error
	}{
		{
			name: "Source folder does not exist",
			src: "abc",
			dst: "true-beetle",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			Error: errors.New("Error: Source folder does not exist"),
		},
		{
			name: "Destination folder does not exist",
			src: "true-beetle",
			dst: "abc",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			Error: errors.New("Error: Destination folder does not exist"),
		},
		{
			name: "Cannot move the folder to itself ",
			src: "true-beetle",
			dst: "true-beetle",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			Error: errors.New("Error: Cannot move a folder to itself"),
		},
		{
			name: "Cannot move to a different organization",
			src: "enabled-professor-monster",
			dst: "sacred-lady-shiva",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			Error: errors.New("Error: Cannot move a folder to a different organization"),
		},{
			name: "Cannot move to a child of itself", 
			src: "steady-insect",
			dst: "stirred-gunslinger",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			Error: errors.New("Error: Cannot move a folder to a child of itself"),
		},
		{
			name: "Valid Movement", 
			src: "stirred-gunslinger",
			dst: "steady-insect",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			Error: nil,
		},
		{
			name: "Valid Movement", 
			src: "national-screwball",
			dst: "endless-red-hulk",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			Error: nil,
		},
		{
			name: "Valid Movement", 
			src: "national-screwball",
			dst: "endless-red-hulk",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			Error: nil,
		},
		{
		name: "Valid Movement",
		src: "composed-wallflower",
		dst: "stunning-horridus",
		orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
		Error: nil,
	},

	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		
		get, err := f.MoveFolder(tt.src, tt.dst)

		// Error checking

		if err != nil {
			if tt.Error == nil || err.Error() != tt.Error.Error() {
					t.Errorf("Test %s failed. Got Error: %v, Want: %v", tt.name, err, tt.Error)
			}} else {
			
				// Validating the index position of the source folder in each occurrence of the source 
				// folder within the paths to ensure correctness in movement operations.

				for _,i := range get {
					if strings.Contains(i.Paths, tt.src) {
						foldersList := strings.Split(i.Paths, ".")
						dstindex := 0
						srcindex := 0
						for i, vs := range foldersList {
							if vs == tt.src {
								srcindex = i
							} else if vs == tt.dst {
								dstindex = i
							}
						}	

						if srcindex != dstindex + 1 {
							t.Errorf("Test %s failed. Got Error: %v, Want: %v", tt.name, err, tt.Error)
						}
					}
				}		
			}
	})
	}
}



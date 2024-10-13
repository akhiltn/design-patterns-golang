package structural

import (
	"testing"
)

// TestCompositePattern tests the GetDetails method for files and folders
func TestCompositePattern(t *testing.T) {
	// Create files
	file1 := &File{Name: "File1.txt", Size: 100}
	file2 := &File{Name: "File2.txt", Size: 150}
	file3 := &File{Name: "File3.txt", Size: 200}

	// Create folders and add files to them
	folder1 := &Folder{Name: "Folder1"}
	folder1.Add(file1)
	folder1.Add(file2)

	folder2 := &Folder{Name: "Folder2"}
	folder2.Add(file3)

	// Create root folder and add other folders/files
	rootFolder := &Folder{Name: "RootFolder"}
	rootFolder.Add(folder1)
	rootFolder.Add(folder2)

	// Expected output string
	expected := `Folder: RootFolder
    Folder: Folder1
        File: File1.txt (Size: 100)
        File: File2.txt (Size: 150)
    Folder: Folder2
        File: File3.txt (Size: 200)
`

	// Get details from the root folder
	actual := rootFolder.GetDetails("")

	// Compare expected and actual results
	if actual != expected {
		t.Errorf("Expected:\n%s\nbut got:\n%s", expected, actual)
	}
}

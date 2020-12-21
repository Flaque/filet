package filet

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTmpDir(t *testing.T) {
	t.Parallel()

	defer CleanUp(t)

	path := TmpDir(t, "")
	assert.Equal(t, Exists(t, path), true,
		"TmpDir should create a directory")
}

func TestTmpFile(t *testing.T) {
	t.Parallel()

	defer CleanUp(t)

	// Test that file is actually created
	file := TmpFile(t, "", "")
	assert.Equal(t, Exists(t, file.Name()), true,
		"TmpFile should create the file")

	// Test that the content exists in the file
	file = TmpFile(t, "", "hey there")
	result := FileSays(t, file.Name(), []byte("hey there"))
	assert.Equal(t, result, true,
		"TmpFile should create a file with content")
}

func TestFile(t *testing.T) {
	t.Parallel()

	defer CleanUp(t)

	// Test that file is actually created
	file := File(t, "conf.yaml", "")
	require.FileExists(t, file.Name(), "File should create the file")

	// Test that the content exists in the file
	file = File(t, "conf.yaml", "hey there")
	result := FileSays(t, file.Name(), []byte("hey there"))
	assert.True(t, result, "File should create a file with content")
}

func TestFileSays(t *testing.T) {
	t.Parallel()

	defer CleanUp(t)

	file := TmpFile(t, "", "Gandhi")
	assert.Equal(t, FileSays(t, file.Name(), []byte("Gandhi")), true,
		"FileSays can correctly read a file.")
	assert.Equal(t, FileSays(t, file.Name(), []byte("nope!")), false,
		"FileSays can correctly tell when a file does not contain content.")

	newT := testing.T{}
	FileSays(&newT, "IdontEx", []byte("hi"))
	assert.Equal(t, newT.Failed(), true,
		"FileSays should fail the testing interface")
}

func TestExists(t *testing.T) {
	t.Parallel()

	defer CleanUp(t)

	file := TmpFile(t, "", "I exist")
	assert.Equal(t, Exists(t, file.Name()), true,
		"Exists should correctly tell if a path exists")
	assert.Equal(t, Exists(t, "blahblahblah"), false,
		"Exists should correctly tell if a path does not exist")
}

func TestDirContains(t *testing.T) {
	t.Parallel()

	defer CleanUp(t)

	dir := TmpDir(t, "")
	filePath := filepath.Base(TmpFile(t, dir, "").Name())
	assert.Equal(t, DirContains(t, dir, filePath), true,
		"DirContains should correctly identify a file inside of a folder.")
	assert.Equal(t, DirContains(t, dir, "blahlleijiow"), false,
		"DirContains should correctly identify a file not inside of a folder.")
}

func TestCleanUp(t *testing.T) {
	t.Parallel()

	defer CleanUp(t) // Kind of problematic.

	// Create test files
	dir := TmpDir(t, "")
	secondDir := TmpDir(t, "")

	// Clean up and test that the files are gone
	CleanUp(t)
	assert.Equal(t, Exists(t, dir), false, "CleanUp should remove files.")
	assert.Equal(t, Exists(t, secondDir), false, "CleanUp should remove files.")

	// Create a new file and test that it's there.
	newDir := TmpDir(t, "")
	assert.Equal(t, Exists(t, newDir), true,
		"New files still exist after CleanUp")
}

func TestCleanUpEmpty(t *testing.T) {
	t.Parallel()

	defer CleanUp(t)
}

func TestMain(t *testing.T) {
	theTest(t)
}

func TestMain2(t *testing.T) {
	theTest(t)
}

func theTest(t *testing.T) {
	t.Parallel()

	defer CleanUp(t)

	TmpDir(t, "")

	time.Sleep(20 * time.Millisecond)
	assert.Len(t, Files(t), 1)
}

package filet

import (
  "github.com/spf13/afero"
  "testing"
  "bytes"
  "path/filepath"
)

// Keeps track of files that we've used so we can clean up.
var TestRegistry []string
var appFs afero.Fs = afero.NewOsFs()

/*
Creates a tmp directory for us to use.
*/
func TmpDir(t *testing.T, dir string) string {
  name, err := afero.TempDir(appFs, dir, "dir")
  if err != nil {
  t.Error("Failed to create the tmpDir: " + name, err)
  }

  TestRegistry = append(TestRegistry, name)
  return name
}

/*
Creates a tmp file for us to use when testing
*/
func TmpFile(t *testing.T, dir string, content string) afero.File {
  file, err := afero.TempFile(appFs, dir, "file")
  if err != nil {
  t.Error("Failed to create the tmpFile: " + file.Name(), err)
  }

  file.WriteString(content)
  TestRegistry = append(TestRegistry, file.Name())

  return file
}

/*
Returns true if the file at the path contains the expected byte array content.
 */
func FileSays(t *testing.T, path string, expected []byte) bool {
  content, err := afero.ReadFile(appFs, path)
  if err != nil{
    t.Error("Failed to read file: " + path, err)
  }

  return bytes.Equal(content, expected)
}

/*
Removes all files in our test registry and calls `t.Error` if something goes
wrong.
*/
func CleanUp(t *testing.T) {
  for _, path := range TestRegistry {
  if err := appFs.RemoveAll(path); err != nil {
  t.Error(appFs.Name(), err)
  }
  }

  TestRegistry = make([]string, 0)
}

/*
Returns true if the file exists. Calls t.Error if something goes wrong while
checking.
 */
func Exists(t *testing.T, path string) bool {
  exists, err := afero.Exists(appFs, path)
  if err != nil {
  t.Error("Something went wrong when checking if " + path + "exists!", err)
  }
  return exists
}

/*
Returns true if the dir contains the path. Calls t.Error if something goes wrong while
checking.
 */
func DirContains(t *testing.T, dir string, path string) bool {
  fullPath := filepath.Join(dir, path)
  return Exists(t, fullPath)
}

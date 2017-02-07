package main

import (
	. "github.com/franela/goblin"
	"testing"
  "path/filepath"
)

func TestTmpDir(t *testing.T) {
  defer CleanUp(t)

  g := Goblin(t)
  g.Describe("TmpDir", func() {
    g.It("should create the directory", func() {
      path := TmpDir(t, "")
      g.Assert(Exists(t, path)).Equal(true)
    })
  })
}

func TestTmpFile(t *testing.T) {
  defer CleanUp(t)

  g := Goblin(t)
  g.Describe("TmpFile", func() {
    g.It("should create the file", func() {
      file := TmpFile(t, "", "")
      g.Assert(Exists(t, file.Name())).Equal(true)
    })

    g.It("should create the file", func() {
      file := TmpFile(t, "", "hey there")
      g.Assert(FileSays(t, file.Name(), []byte("hey there"))).Equal(true)
    })
  })
}

func TestFileSays(t *testing.T) {
  defer CleanUp(t)

  g := Goblin(t)
  g.Describe("FileSays", func() {
    g.It("should correctly identify strings from files", func() {
      file := TmpFile(t, "", "Ghandi")
      g.Assert(FileSays(t, file.Name(), []byte("Ghandi"))).Equal(true)
      g.Assert(FileSays(t, file.Name(), []byte("nope!"))).Equal(false)
    })
  })
}

func TestExists(t *testing.T) {
  defer CleanUp(t)

  g := Goblin(t)
  g.Describe("Exists", func() {
    g.It("should correctly tell when something exists", func() {
      file := TmpFile(t, "", "I exist")
      g.Assert(Exists(t, file.Name())).Equal(true)
      g.Assert(Exists(t, "BlahBlahBlahNopeNope")).Equal(false)
    })
  })
}

func TestDirContains(t *testing.T) {
  defer CleanUp(t)

  g := Goblin(t)
  g.Describe("DirContains", func() {
    g.It("should correctly tell when a path is inside another dir", func() {
      dir := TmpDir(t, "")
      filePath := filepath.Base(TmpFile(t, dir, "").Name())
      g.Assert(DirContains(t, dir, filePath)).Equal(true)
      g.Assert(DirContains(t, dir, "bloobloobloo")).Equal(false)
    })
  })
}

func TestCleanUp(t *testing.T) {
  defer CleanUp(t) // Kind of problematic.

  // Clear testRegistry
  testRegistry = make([]string, 0)

  g := Goblin(t)
  g.Describe("CleanUp", func() {
    g.It("should correctly tell when a path is cleaned up", func() {
      dir := TmpDir(t, "")
      secondDir := TmpDir(t, "")

      // Clean up and test that the files are gone
      CleanUp(t)
      g.Assert(Exists(t, dir)).Equal(false)
      g.Assert(Exists(t, secondDir)).Equal(false)

      // Create a new file and test that it's there.
      newDir := TmpDir(t, "")
      g.Assert(Exists(t, newDir)).Equal(true)
    })
  })
}

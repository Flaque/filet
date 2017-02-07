# Filet 🍖
A small temporary file utility for Go for testing. Built on Afero and heavily
inspired by the way Afero tests itself.

Install via:
`$ go get github.com/Flaque/filet`

Then import with:
```
import (
  filet "github.com/Flaque/filet"
)
```

### Creating temporaries

#### Creating temporary directories:
```
err_1 := filet.TmpDir("") // Creates a temporary dir with no parent directory
err_2 := filet.TmpDir("myPath") // Creates a temporary dir at `myPath`
```

#### Creating temporary files:
```
err_1 := filet.TmpFile("", "") // Creates a temporary file with no parent dir

// Creates a temporary file with string "some content"
err_2 := filet.TmpFile("", "some content")

// Creates a temporary file with string "some content"
err_2 := filet.TmpFile("myDir", "some content")
```

#### Cleaning up after yourself:
Filet lets you clean up after your files with `CleanUp`, which is
most cleanly used at the beginning of a function with `defer`. For example:

```
func TestMyFoo(t *testing.T) {
  defer filet.CleanUp(t)

  // Create a bunch of temporary stuff here
}
```

`CleanUp` will call `t.Error` if something goes wrong when removing the file.

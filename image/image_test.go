package image

import (
	"io/ioutil"
	"testing"
)

func TestGet(t *testing.T) {
	targetDir := "./tmp/test"
	Get(targetDir, "go", "import \"fmt\"")

	files, _ := ioutil.ReadDir(targetDir)
	if len(files) != 3 {
		t.Errorf("files number not correct, got: %d, want: %d.", len(files), 3)
	}

	if files[0].Name() != "Dockerfile" {
		t.Errorf("Dockerfile not correct, got: %d, want: %d.", files[0], "Dockerfile")
	}

	if files[1].Name() != "app.go" {
		t.Errorf("app.go not correct, got: %d, want: %d.", files[1], "app.go")
	}

	if files[2].Name() != "fx.go" {
		t.Errorf("fx.go not correct, got: %d, want: %d.", files[2], "fx.go")
	}

	data, _ := ioutil.ReadFile(files[2].Name())
	if string(data) != "import \"fmt\"" {
		t.Errorf("content of fx.go not correct, got: %d, want: %d.", string(data), "import \"fmt\"")
	}
}

package primitive

import (
	"io"
	"os"
	"testing"
)

func TestTransform(t *testing.T) {
	var file io.Reader
	file, err := os.Open("./img/input.png")
	file, err = Transform(file, "png", "1", "100")
	if err != nil {
		t.Errorf("error")
	}
}
func TestPrimitive(t *testing.T) {
	_, err := primitive("-i ./img/in.png -o ./img/out.png -n 10 -m 0", "out.png")
	if err == nil {
		t.Error("error in primitive")
	}
	_, err = primitive("-i ./img/input.png -o ./img/output.png -n 10 -m 0", "out.png")
	if err == nil {
		t.Error("error in primitive")
	}
}

func TestTempfile(t *testing.T) {
	_, err := tempfile("", "")
	if err != nil {
		t.Error("error in tempfile")
	}
}

func TestPrimitiveError(t *testing.T) {
	_, err := primitive("-i input.png -o output.png -n 10 -m 0", "")
	if err != nil {
		t.Error("error in primitive")
	}
}

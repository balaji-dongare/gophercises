package primitive

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Transform function will transform the image first it will create two file
// input file and output file the then call the primitive function which
// actualy trasform the image and this function return io.Reader and error
func Transform(image io.Reader, ext, mode, number string) (io.Reader, error) {
	var outFile io.Reader
	// create the input file
	in, err := tempfile("in_", ext)
	if err == nil {
		defer os.Remove(in.Name())

		// create the output file
		out, err := tempfile("out_", ext)
		if err == nil {
			defer os.Remove(out.Name())
			//copy the uploaded image to input file
			_, err = io.Copy(in, image)
			if err != nil {
				fmt.Println(err)
			}
			// creating arguments
			args := fmt.Sprintf("-i %s -o %s -n %s -m %s", in.Name(), out.Name(), number, mode)
			// calling the primitive function
			outFile, err = primitive(args, out.Name())

		}
	}

	return outFile, err
}

// this function actually transform the image file it takes args of primitive library
// command and file name and return io.Reader and error
func primitive(args, fileName string) (io.Reader, error) {
	cmd := exec.Command("primitive", strings.Fields(args)...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	out, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// this function create the temprary file for us
// it takes prefix for the file and the extension for the file
// and it return file and error
func tempfile(prefix, ext string) (*os.File, error) {
	var out *os.File
	in, err := ioutil.TempFile("", prefix)
	if err == nil {
		defer os.Remove(in.Name())
		out, err = os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
	}
	return out, err
}

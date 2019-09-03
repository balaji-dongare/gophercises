package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/balaji-dongare/gophercises/ImageTransform/primitive"
)

// it will handle request at "/" show the html page for upload image
func indexHTML(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
		<body>
				<form action="/upload" method="post" enctype="multipart/form-data">
					<div class="form-group">
						<label for="image">Choose Image</label>
						<input class="form-control" type="file" name="image" id="image">
					</div>
					<button type="submit">Upload</button>
				</form>
		</body>
	</html`
	fmt.Fprint(w, html)

}

// uploadHandle is handle the request at "/upload" it  save the file for persistance
//redirect to /modify url
func uploadHandle(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	ext := filepath.Ext(header.Filename)[1:]
	saveFile, err := tempfile("", ext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	io.Copy(saveFile, file)
	http.Redirect(w, r, "modify/"+filepath.Base(saveFile.Name()), http.StatusFound)
	errorResponse(w, err)
}

// modifyHandle is used for the show all the transformed image based on the query params
// if mode is not defined in the url it will show four different image with different images
// if mode is defined and number of shapes is not defined it will show the 3 images with the same mode
// but different number of shapes
// if mode and number of shapes is in url the it will show one image
func modifyImageHandle(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./img/" + filepath.Base(r.URL.Path))
	if err == nil {
		ext := filepath.Ext(file.Name())[1:]
		mode := r.FormValue("mode")
		if mode == "" {
			transfromInAllMode(w, r, file, ext)
			return
		}
		number := r.FormValue("number")
		if number == "" {
			transformInSingleMode(w, r, file, ext, mode)
			return
		}
		log.Printf("file path :%v", filepath.Base(file.Name()))
		http.Redirect(w, r, "/img/"+string(filepath.Base(file.Name())), http.StatusFound)
	}
	errorResponse(w, err)
}

// this function will generate three image with the same mode
// but different number of shapes
func transformInSingleMode(w http.ResponseWriter, r *http.Request, file io.ReadSeeker, ext, mode string) {
	var a, b, c string
	var err error
	a, err = genImage(file, ext, mode, "50")
	//log.Printf("name :%v", a)
	if err == nil {
		file.Seek(0, 0)
		b, err = genImage(file, ext, mode, "100")
		if err == nil {
			file.Seek(0, 0)
			c, err = genImage(file, ext, mode, "150")
			if err == nil {
				html := `<html><body>
						{{range .}}
							<a href="/modify/img/{{.Name}}?mode={{.Mode}}&number={{.Number}}">
							<img style="width: 40%;" src="/img/{{.Name}}">
							</a>
						{{end}}
						</body></html>`
				tpl := template.Must(template.New("").Parse(html))
				type Images struct {
					Name   string
					Mode   int
					Number int
				}
				images := []Images{
					{a, 2, 50}, {b, 2, 100}, {c, 2, 150},
				}

				tpl.Execute(w, images)
			}
		}
	}

}

// this funcion generate the four different images with differnt modes
func transfromInAllMode(w http.ResponseWriter, r *http.Request, file io.ReadSeeker, ext string) {
	var img1, img2, img3, img4 string
	var err error
	img1, err = genImage(file, ext, "2", "100")
	if err == nil {
		file.Seek(0, 0)
		img2, err = genImage(file, ext, "3", "30")
		if err == nil {
			file.Seek(0, 0)
			img3, err = genImage(file, ext, "4", "50")
			if err == nil {
				file.Seek(0, 0)
				img4, err = genImage(file, ext, "5", "700")
				if err == nil {
					file.Seek(0, 0)
					html := `<html><body>
					{{range .}}
						<a href="/modify/img/{{.Name}}?mode={{.Mode}}">
						<img style="width: 40%;" src="/img/{{.Name}}">
						</a>
					{{end}}
					</body></html>`
					tpl := template.Must(template.New("").Parse(html))
					type Images struct {
						Name string
						Mode int
					}
					images := []Images{
						{img1, 2}, {img2, 3}, {img3, 4}, {img4, 5},
					}

					tpl.Execute(w, images)
				}
			}
		}
	}

	errorResponse(w, err)
}

// this function call the transform function from transform package
// to get trasform image
func genImage(file io.Reader, ext, mode, number string) (string, error) {
	var out io.Reader
	var err error
	var fileName string
	out, err = primitive.Transform(file, ext, mode, number)
	if err == nil {
		var outFile *os.File
		outFile, err = tempfile("", ext)
		if err == nil {
			io.Copy(outFile, out)
			fileName = outFile.Name()
		}
	}

	fileName = strings.Replace(fileName, "img\\", "", -1)
	return fileName, nil
}

func errorResponse(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// this function generate the temp file for us
func tempfile(prefix, ext string) (*os.File, error) {
	var in, out *os.File
	var err error
	in, err = ioutil.TempFile("./img/", prefix)
	if err == nil {
		defer os.Remove(in.Name())
		out, err = os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
		return out, err
	}
	return out, err
}

// getHandlers will return the router mux with handlers
func getHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHTML)
	mux.HandleFunc("/upload", uploadHandle)
	fs := http.FileServer(http.Dir("./img/"))
	mux.Handle("/img/", http.StripPrefix("/img/", fs))
	mux.HandleFunc("/modify/", modifyImageHandle)
	return mux
}

// this is main function we defined all the router here
func main() {
	http.ListenAndServe(":3000", getHandlers())
}

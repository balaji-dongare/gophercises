package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime/debug"
	"strconv"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

//debugHandler  it will handle the request on /panic and calles panic method
func debugHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	lineNo := r.FormValue("line")
	line, err := strconv.Atoi(lineNo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bffer := bytes.NewBuffer(nil)
	io.Copy(bffer, file)
	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}
	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, bffer.String())
	style := styles.Get("monokai")
	formatter := html.New(html.TabWidth(2), html.WithLineNumbers(), html.HighlightLines(lines))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<style>pre { font-size: 1.2em; }</style>")
	formatter.Format(w, style, iterator)
}

//panicHandler  it will handle the request on /debug
//shows the panic with link of  error line and line number
func panicHandler(w http.ResponseWriter, r *http.Request) {
	funcPanic()
}

func funcPanic() {
	panic("OMG !")
}

// getHandlers return the router mux
func getHandles() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/", debugHandler)
	mux.HandleFunc("/panic/", panicHandler)
	return mux
}

//recoveryHandle it handles the panic func and retrun stackstrace
func recoveryHandle(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, linkForm(string(stack)))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

//linkForm  it will parse the linke and make link using filename:linenuber
func linkForm(stack string) string {
	re := regexp.MustCompile("(\t.*:[0-9]*)")
	lines := re.FindAllString(stack, -1)
	for _, line := range lines {
		regexSplit := regexp.MustCompile(":")
		splits := regexSplit.Split(line, -1)
		link := "<a href='/debug/?path=" + splits[0] + "&line=" + splits[1] + "'>" + line + "</a>"
		reg := regexp.MustCompile(line)
		stack = reg.ReplaceAllString(stack, link)
	}
	return stack
}
func main() {
	http.ListenAndServe(":3000", recoveryHandle(getHandles()))
}

package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
)

type WorkingDay struct {
	Success bool

	ClockIn    string
	StartBreak string
	EndBreak   string
	ClockOut   string
	Total      string
}

// Embedding form as string
//go:embed forms.html
var form string

// Loads the forms and runs the server.
func runServer() {

	//tmpl := template.Must(template.ParseFiles("forms.html"))

	// Loading from an embedded form.
	tmpl, _ := template.New("forms.html").Parse(form)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		workingDay := WorkingDay{
			Success: true,

			ClockIn:    r.FormValue("clock_in"),
			StartBreak: r.FormValue("start_break"),
			EndBreak:   r.FormValue("end_break"),
			ClockOut:   r.FormValue("clock_out"),
		}

		tmpl.Execute(w, workingDay)

	})

	port := "8080"
	fmt.Println("Running on localhost:" + port)
	http.ListenAndServe(":"+port, nil)

}

func main() {
	runServer()
}

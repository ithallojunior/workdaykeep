package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type WorkingDay struct {
	// Control var
	IsValid bool

	ClockIn    string
	StartBreak string
	EndBreak   string
	ClockOut   string
	Total      string
}

// Embedding form as string
//go:embed forms.html
var form string

// Validates and adds new values for start break (4 hours after clock in)
// end break (one hour after start break) and
// calculates clock out (8 hours work journey) if none passed/invalid.
func (w *WorkingDay) ValidateAndUpdate() {

	layout := "15:04"

	clockIn, err := time.Parse(layout, w.ClockIn)
	if err != nil {
		w.IsValid = false
		return
	}

	journey := 8 * time.Hour

	startBreak, err := time.Parse(layout, w.StartBreak)
	if err != nil {
		startBreak = clockIn.Add(4 * time.Hour)
		w.StartBreak = startBreak.Format(layout)
	}

	endBreak, err := time.Parse(layout, w.EndBreak)
	if err != nil {
		endBreak = startBreak.Add(time.Hour)
		w.EndBreak = endBreak.Format(layout)
	}

	// start break -  clock in
	d1 := startBreak.Sub(clockIn)

	clockOut, err := time.Parse(layout, w.ClockOut)
	if err != nil {
		clockOut = endBreak.Add(journey - d1)
		w.ClockOut = clockOut.Format(layout)
	}

	total := d1 + clockOut.Sub(endBreak)
	w.Total = total.String()

	w.IsValid = true
}

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
			ClockIn:    r.FormValue("clock_in"),
			StartBreak: r.FormValue("start_break"),
			EndBreak:   r.FormValue("end_break"),
			ClockOut:   r.FormValue("clock_out"),
		}

		workingDay.ValidateAndUpdate()

		data := struct {
			GotData bool
			Day     WorkingDay
		}{
			GotData: true,
			Day:     workingDay,
		}

		tmpl.Execute(w, data)

	})

	port := "8080"
	fmt.Println("Running on localhost:" + port)
	http.ListenAndServe(":"+port, nil)

}

func main() {
	runServer()
}

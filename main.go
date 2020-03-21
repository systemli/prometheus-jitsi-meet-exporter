package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"text/template"
)

var (
	addr           = flag.String("web.listen-address", ":9888", "Address on which to expose metrics and web interface.")
	videoBridgeURL = flag.String("videobridge-url", "http://localhost:8888/stats", "Jitsi Videobridge /stats URL to scrape")
)

type videoBridgeStats struct {
	LargestConference       int   `json:"largest_conference"`
	TotalParticipants       int   `json:"total_participants"`
	ConferenceSizes         []int `json:"conference_sizes"`
	TotalConferencesCreated int   `json:"total_conferences_created"`
	Conferences             int   `json:"conferences"`
	Participants            int   `json:"participants"`
}

var tpl = template.Must(template.New("stats").Parse(`# HELP jitsi_total_participants Participants counter
# TYPE jitsi_total_participants counter
jitsi_total_participants {{.TotalParticipants}}
# HELP jitsi_total_conferences_created Number of conferences created
# TYPE jitsi_total_conferences_created counter
jitsi_total_conferences_created {{.TotalConferencesCreated}}
# HELP jitsi_largest_conference Participants in the largest conference
# TYPE jitsi_largest_conference gauge
jitsi_largest_conference {{.LargestConference}}
# HELP jitsi_conferences Current number of active conferences
# TYPE jitsi_conferences gauge
jitsi_conferences {{.Conferences}}
# HELP jitsi_participants Current number of active participants
# TYPE jitsi_participants gauge
jitsi_participants {{.Participants}}
`))

func serveMetrics(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get(*videoBridgeURL)
	if err != nil {
		log.Printf("scrape error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var stats videoBridgeStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		log.Printf("json decoding error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	_ = tpl.Execute(w, &stats)
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	http.HandleFunc("/metrics", serveMetrics)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Started Jitsi Meet Metrics Exporter")
}

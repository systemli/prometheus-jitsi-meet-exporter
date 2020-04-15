package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMetrics(t *testing.T) {
	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		http.HandleFunc("/stats", getJitsiVideobridgeStats)
		if err := http.ListenAndServe(":8888", nil); err != nil {
			log.Fatal(err)
		}
	}()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveMetrics)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var expected = `# HELP jitsi_total_participants Participants counter
# TYPE jitsi_total_participants counter
jitsi_total_participants 18
# HELP jitsi_total_conferences_created Number of conferences created
# TYPE jitsi_total_conferences_created counter
jitsi_total_conferences_created 14
# HELP jitsi_largest_conference Participants in the largest conference
# TYPE jitsi_largest_conference gauge
jitsi_largest_conference 3
# HELP jitsi_conferences Current number of active conferences
# TYPE jitsi_conferences gauge
jitsi_conferences 2
# HELP jitsi_participants Current number of active participants
# TYPE jitsi_participants gauge
jitsi_participants 4
`

	if rr.Body.String() != expected {
		t.Error("Response does not match the expected string")
	}
}

func getJitsiVideobridgeStats(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte(`{"largest_conference":3,"total_sip_call_failures":0,"total_participants":18,"conference_sizes":[0,1,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"bridge_selector":{"total_least_loaded_in_region":0,"total_split_due_to_load":0,"total_not_loaded_in_region_in_conference":0,"total_least_loaded_in_region_in_conference":0,"total_not_loaded_in_region":0,"total_split_due_to_region":0,"bridge_count":1,"operational_bridge_count":1,"total_least_loaded_in_conference":0,"total_least_loaded":3},"total_conferences_created":14,"total_recording_failures":0,"conferences":2,"total_live_streaming_failures":0,"participants":4}`))
}

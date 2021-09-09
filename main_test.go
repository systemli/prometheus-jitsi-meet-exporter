package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type constHandler struct {
	s string
}

func (h constHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.s))
}

func TestGetMetrics(t *testing.T) {
	tcs := []struct {
		statsJson string
		expected  string
	}{
		{
			statsJson: `{"largest_conference":3,"total_sip_call_failures":0,"total_participants":18,"conference_sizes":[0,1,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"bridge_selector":{"total_least_loaded_in_region":0,"total_split_due_to_load":0,"total_not_loaded_in_region_in_conference":0,"total_least_loaded_in_region_in_conference":0,"total_not_loaded_in_region":0,"total_split_due_to_region":0,"bridge_count":1,"operational_bridge_count":1,"total_least_loaded_in_conference":0,"total_least_loaded":3},"total_conferences_created":14,"total_conferences_completed":0,"total_recording_failures":0,"conferences":2,"p2p_conferences":1,"total_live_streaming_failures":0,"participants":4}`,
			expected: `# HELP jitsi_threads The number of Java threads that the video bridge is using.
# TYPE jitsi_threads gauge
jitsi_threads 0
# HELP jitsi_bit_rate_download The total incoming bitrate for the video bridge in kilobits per second.
# TYPE jitsi_bit_rate_download gauge
jitsi_bit_rate_download 0
# HELP jitsi_bit_rate_upload The total outgoing bitrate for the video bridge in kilobits per second.
# TYPE jitsi_bit_rate_upload gauge
jitsi_bit_rate_upload 0
# HELP jitsi_packet_rate_download The total incoming packet rate for the video bridge in packets per second.
# TYPE jitsi_packet_rate_download gauge
jitsi_packet_rate_download 0
# HELP jitsi_packet_rate_upload The total outgoing packet rate for the video bridge in packets per second.
# TYPE jitsi_packet_rate_upload gauge
jitsi_packet_rate_upload 0
# HELP jitsi_loss_rate_download The fraction of lost incoming RTP packets. This is based on RTP sequence numbers and is relatively accurate.
# TYPE jitsi_loss_rate_download gauge
jitsi_loss_rate_download 0
# HELP jitsi_loss_rate_upload The fraction of lost outgoing RTP packets. This is based on incoming RTCP Receiver Reports, and an attempt to subtract the fraction of packets that were not sent (i.e. were lost before they reached the bridge). Further, this is averaged over all streams of all users as opposed to all packets, so it is not correctly weighted. This is not accurate, but may be a useful metric nonetheless.
# TYPE jitsi_loss_rate_upload gauge
jitsi_loss_rate_upload 0
# HELP jitsi_jitter_aggregate Experimental. An average value (in milliseconds) of the jitter calculated for incoming and outgoing streams. This hasn't been tested and it is currently not known whether the values are correct or not.
# TYPE jitsi_jitter_aggregate gauge
jitsi_jitter_aggregate 0
# HELP jitsi_rtt_aggregate An average value (in milliseconds) of the RTT across all streams.
# TYPE jitsi_rtt_aggregate gauge
jitsi_rtt_aggregate 0
# HELP jitsi_largest_conference The number of participants in the largest conference currently hosted on the bridge.
# TYPE jitsi_largest_conference gauge
jitsi_largest_conference 3
# HELP jitsi_audiochannels The current number of audio channels.
# TYPE jitsi_audiochannels gauge
jitsi_audiochannels 0
# HELP jitsi_videochannels The current number of video channels.
# TYPE jitsi_videochannels gauge
jitsi_videochannels 0
# HELP jitsi_conferences The current number of conferences.
# TYPE jitsi_conferences gauge
jitsi_conferences 2
# HELP jitsi_p2p_conferences The current number of p2p conferences.
# TYPE jitsi_p2p_conferences gauge
jitsi_p2p_conferences 1
# HELP jitsi_participants The current number of participants.
# TYPE jitsi_participants gauge
jitsi_participants 4
# HELP jitsi_total_participants Total participants since running.
# TYPE jitsi_total_participants gauge
jitsi_total_participants 18
# HELP jitsi_videostreams An estimation of the number of current video streams forwarded by the bridge.
# TYPE jitsi_videostreams gauge
jitsi_videostreams 0
# HELP jitsi_stress_level Stress Level reported to Jicofo by the videobridge.
# TYPE jitsi_stress_level gauge
jitsi_stress_level 0
# HELP jitsi_total_loss_controlled_participant_seconds The total number of participant-seconds that are loss-controlled.
# TYPE jitsi_total_loss_controlled_participant_seconds counter
jitsi_total_loss_controlled_participant_seconds 0
# HELP jitsi_total_loss_limited_participant_seconds The total number of participant-seconds that are loss-limited.
# TYPE jitsi_total_loss_limited_participant_seconds counter
jitsi_total_loss_limited_participant_seconds 0
# HELP jitsi_total_loss_degraded_participant_seconds The total number of participant-seconds that are loss-degraded.
# TYPE jitsi_total_loss_degraded_participant_seconds counter
jitsi_total_loss_degraded_participant_seconds 0
# HELP jitsi_total_conference_seconds The sum of the lengths of all completed conferences, in seconds.
# TYPE jitsi_total_conference_seconds counter
jitsi_total_conference_seconds 0
# HELP jitsi_total_conferences_created The total number of conferences created on the bridge.
# TYPE jitsi_total_conferences_created counter
jitsi_total_conferences_created 14
# HELP jitsi_total_conferences_completed The total number of conferences completed on the bridge.
# TYPE jitsi_total_conferences_completed counter
jitsi_total_conferences_completed 0
# HELP jitsi_total_failed_conferences The total number of failed conferences on the bridge. A conference is marked as failed when all of its channels have failed. A channel is marked as failed if it had no payload activity.
# TYPE jitsi_total_failed_conferences counter
jitsi_total_failed_conferences 0
# HELP jitsi_total_partially_failed_conferences The total number of partially failed conferences on the bridge. A conference is marked as partially failed when some of its channels has failed. A channel is marked as failed if it had no payload activity.
# TYPE jitsi_total_partially_failed_conferences counter
jitsi_total_partially_failed_conferences 0
# HELP jitsi_total_data_channel_messages_received The total number messages received through data channels.
# TYPE jitsi_total_data_channel_messages_received counter
jitsi_total_data_channel_messages_received 0
# HELP jitsi_total_data_channel_messages_sent The total number messages sent through data channels.
# TYPE jitsi_total_data_channel_messages_sent counter
jitsi_total_data_channel_messages_sent 0
# HELP jitsi_total_colibri_web_socket_messages_received The total number messages received through COLIBRI web sockets.
# TYPE jitsi_total_colibri_web_socket_messages_received counter
jitsi_total_colibri_web_socket_messages_received 0
# HELP jitsi_total_colibri_web_socket_messages_sent The total number messages sent through COLIBRI web sockets.
# TYPE jitsi_total_colibri_web_socket_messages_sent counter
jitsi_total_colibri_web_socket_messages_sent 0
# HELP jitsi_octo_version The current running OCTO version
# TYPE jitsi_octo_version gauge
jitsi_octo_version 0
# HELP jitsi_octo_conferences The current number of OCTO conferences.
# TYPE jitsi_octo_conferences gauge
jitsi_octo_conferences 0
# HELP jitsi_octo_endpoints The current number of OCTO endpoints.
# TYPE jitsi_octo_endpoints gauge
jitsi_octo_endpoints 0
# HELP jitsi_octo_receive_bitrate The total receiving bitrate for the OCTO video bridge in kilobits per second.
# TYPE jitsi_octo_receive_bitrate gauge
jitsi_octo_receive_bitrate 0
# HELP jitsi_octo_send_bitrate The total outgoing bitrate for the OCTO video bridge in kilobits per second.
# TYPE jitsi_octo_send_bitrate gauge
jitsi_octo_send_bitrate 0
# HELP jitsi_octo_receive_packet_rate The total incoming packet rate for the OCTO video bridge in packets per second.
# TYPE jitsi_octo_receive_packet_rate gauge
jitsi_octo_receive_packet_rate 0
# HELP jitsi_octo_send_packet_rate The total outgoing packet rate for the OCTO video bridge in packets per second.
# TYPE jitsi_octo_send_packet_rate gauge
jitsi_octo_send_packet_rate 0
# HELP jitsi_total_bytes_received_octo The total incoming bit rate for the OCTO video bridge in bytes per second.
# TYPE jitsi_total_bytes_received_octo gauge
jitsi_total_bytes_received_octo 0
# HELP jitsi_total_bytes_sent_octo The total outgoing bit rate for the OCTO video bridge in bytes per second.
# TYPE jitsi_total_bytes_sent_octo gauge
jitsi_total_bytes_sent_octo 0
# HELP jitsi_total_packets_dropped_octo The total of dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_dropped_octo gauge
jitsi_total_packets_dropped_octo 0
# HELP jitsi_total_packets_received_octo The total of incoming dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_received_octo gauge
jitsi_total_packets_received_octo 0
# HELP jitsi_total_packets_sent_octo The total of sent dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_sent_octo gauge
jitsi_total_packets_sent_octo 0
# HELP jitsi_conference_sizes Distribution of conference sizes
# TYPE jitsi_conference_sizes gauge
jitsi_conference_sizes{conference_size="0"} 0
jitsi_conference_sizes{conference_size="1"} 1
jitsi_conference_sizes{conference_size="2"} 0
jitsi_conference_sizes{conference_size="3"} 1
jitsi_conference_sizes{conference_size="4"} 0
jitsi_conference_sizes{conference_size="5"} 0
jitsi_conference_sizes{conference_size="6"} 0
jitsi_conference_sizes{conference_size="7"} 0
jitsi_conference_sizes{conference_size="8"} 0
jitsi_conference_sizes{conference_size="9"} 0
jitsi_conference_sizes{conference_size="10"} 0
jitsi_conference_sizes{conference_size="11"} 0
jitsi_conference_sizes{conference_size="12"} 0
jitsi_conference_sizes{conference_size="13"} 0
jitsi_conference_sizes{conference_size="14"} 0
jitsi_conference_sizes{conference_size="15"} 0
jitsi_conference_sizes{conference_size="16"} 0
jitsi_conference_sizes{conference_size="17"} 0
jitsi_conference_sizes{conference_size="18"} 0
jitsi_conference_sizes{conference_size="19"} 0
jitsi_conference_sizes{conference_size="20"} 0
jitsi_conference_sizes{conference_size="21"} 0
`,
		},
		{
			statsJson: `{
  "audiochannels": 0,
  "bit_rate_download": 0.5,
  "bit_rate_upload": 0.5,
  "conference_sizes": [ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 ],
  "conferences": 0,
  "current_timestamp": "2019-03-14 11:02:15.184",
  "graceful_shutdown": false,
  "jitter_aggregate": 0,
  "largest_conference": 0,
  "loss_rate_download": 0.5,
  "loss_rate_upload": 0.5,
  "octo_conferences": 0,
  "octo_endpoints": 0,
  "octo_receive_bitrate": 0,
  "octo_receive_packet_rate": 0,
  "octo_send_bitrate": 0,
  "octo_send_packet_rate": 0,
  "octo_version": 1,
  "packet_rate_download": 0,
  "packet_rate_upload": 0,
  "participants": 0,
  "region": "eu-west-1",
  "relay_id": "10.0.0.5:4096",
  "rtp_loss": 0,
  "rtt_aggregate": 0,
  "stress_level": 0.6,
  "threads": 59,
  "total_bytes_received": 257628359,
  "total_bytes_received_octo": 0,
  "total_bytes_received_octo": 0,
  "total_bytes_sent": 257754048,
  "total_bytes_sent_octo": 0,
  "total_bytes_sent_octo": 0,
  "total_colibri_web_socket_messages_received": 0,
  "total_colibri_web_socket_messages_sent": 0,
  "total_conference_seconds": 470,
  "total_conferences_completed": 1,
  "total_conferences_created": 1,
  "total_data_channel_messages_received": 602,
  "total_data_channel_messages_sent": 600,
  "total_failed_conferences": 0,
  "total_ice_failed": 0,
  "total_ice_succeeded": 2,
  "total_ice_succeeded_tcp": 0,
  "total_loss_controlled_participant_seconds": 847,
  "total_loss_degraded_participant_seconds": 1,
  "total_loss_limited_participant_seconds": 0,
  "total_packets_dropped_octo": 0,
  "total_packets_dropped_octo": 0,
  "total_packets_received": 266644,
  "total_packets_received_octo": 0,
  "total_packets_received_octo": 0,
  "total_packets_sent": 266556,
  "total_packets_sent_octo": 0,
  "total_packets_sent_octo": 0,
  "total_partially_failed_conferences": 0,
  "total_participants": 2,
  "videochannels": 0,
  "videostreams": 0
}`,
			expected: `# HELP jitsi_threads The number of Java threads that the video bridge is using.
# TYPE jitsi_threads gauge
jitsi_threads 59
# HELP jitsi_bit_rate_download The total incoming bitrate for the video bridge in kilobits per second.
# TYPE jitsi_bit_rate_download gauge
jitsi_bit_rate_download 0.5
# HELP jitsi_bit_rate_upload The total outgoing bitrate for the video bridge in kilobits per second.
# TYPE jitsi_bit_rate_upload gauge
jitsi_bit_rate_upload 0.5
# HELP jitsi_packet_rate_download The total incoming packet rate for the video bridge in packets per second.
# TYPE jitsi_packet_rate_download gauge
jitsi_packet_rate_download 0
# HELP jitsi_packet_rate_upload The total outgoing packet rate for the video bridge in packets per second.
# TYPE jitsi_packet_rate_upload gauge
jitsi_packet_rate_upload 0
# HELP jitsi_loss_rate_download The fraction of lost incoming RTP packets. This is based on RTP sequence numbers and is relatively accurate.
# TYPE jitsi_loss_rate_download gauge
jitsi_loss_rate_download 0.5
# HELP jitsi_loss_rate_upload The fraction of lost outgoing RTP packets. This is based on incoming RTCP Receiver Reports, and an attempt to subtract the fraction of packets that were not sent (i.e. were lost before they reached the bridge). Further, this is averaged over all streams of all users as opposed to all packets, so it is not correctly weighted. This is not accurate, but may be a useful metric nonetheless.
# TYPE jitsi_loss_rate_upload gauge
jitsi_loss_rate_upload 0.5
# HELP jitsi_jitter_aggregate Experimental. An average value (in milliseconds) of the jitter calculated for incoming and outgoing streams. This hasn't been tested and it is currently not known whether the values are correct or not.
# TYPE jitsi_jitter_aggregate gauge
jitsi_jitter_aggregate 0
# HELP jitsi_rtt_aggregate An average value (in milliseconds) of the RTT across all streams.
# TYPE jitsi_rtt_aggregate gauge
jitsi_rtt_aggregate 0
# HELP jitsi_largest_conference The number of participants in the largest conference currently hosted on the bridge.
# TYPE jitsi_largest_conference gauge
jitsi_largest_conference 0
# HELP jitsi_audiochannels The current number of audio channels.
# TYPE jitsi_audiochannels gauge
jitsi_audiochannels 0
# HELP jitsi_videochannels The current number of video channels.
# TYPE jitsi_videochannels gauge
jitsi_videochannels 0
# HELP jitsi_conferences The current number of conferences.
# TYPE jitsi_conferences gauge
jitsi_conferences 0
# HELP jitsi_p2p_conferences The current number of p2p conferences.
# TYPE jitsi_p2p_conferences gauge
jitsi_p2p_conferences 0
# HELP jitsi_participants The current number of participants.
# TYPE jitsi_participants gauge
jitsi_participants 0
# HELP jitsi_total_participants Total participants since running.
# TYPE jitsi_total_participants gauge
jitsi_total_participants 2
# HELP jitsi_videostreams An estimation of the number of current video streams forwarded by the bridge.
# TYPE jitsi_videostreams gauge
jitsi_videostreams 0
# HELP jitsi_stress_level Stress Level reported to Jicofo by the videobridge.
# TYPE jitsi_stress_level gauge
jitsi_stress_level 0.6
# HELP jitsi_total_loss_controlled_participant_seconds The total number of participant-seconds that are loss-controlled.
# TYPE jitsi_total_loss_controlled_participant_seconds counter
jitsi_total_loss_controlled_participant_seconds 847
# HELP jitsi_total_loss_limited_participant_seconds The total number of participant-seconds that are loss-limited.
# TYPE jitsi_total_loss_limited_participant_seconds counter
jitsi_total_loss_limited_participant_seconds 0
# HELP jitsi_total_loss_degraded_participant_seconds The total number of participant-seconds that are loss-degraded.
# TYPE jitsi_total_loss_degraded_participant_seconds counter
jitsi_total_loss_degraded_participant_seconds 1
# HELP jitsi_total_conference_seconds The sum of the lengths of all completed conferences, in seconds.
# TYPE jitsi_total_conference_seconds counter
jitsi_total_conference_seconds 470
# HELP jitsi_total_conferences_created The total number of conferences created on the bridge.
# TYPE jitsi_total_conferences_created counter
jitsi_total_conferences_created 1
# HELP jitsi_total_conferences_completed The total number of conferences completed on the bridge.
# TYPE jitsi_total_conferences_completed counter
jitsi_total_conferences_completed 1
# HELP jitsi_total_failed_conferences The total number of failed conferences on the bridge. A conference is marked as failed when all of its channels have failed. A channel is marked as failed if it had no payload activity.
# TYPE jitsi_total_failed_conferences counter
jitsi_total_failed_conferences 0
# HELP jitsi_total_partially_failed_conferences The total number of partially failed conferences on the bridge. A conference is marked as partially failed when some of its channels has failed. A channel is marked as failed if it had no payload activity.
# TYPE jitsi_total_partially_failed_conferences counter
jitsi_total_partially_failed_conferences 0
# HELP jitsi_total_data_channel_messages_received The total number messages received through data channels.
# TYPE jitsi_total_data_channel_messages_received counter
jitsi_total_data_channel_messages_received 602
# HELP jitsi_total_data_channel_messages_sent The total number messages sent through data channels.
# TYPE jitsi_total_data_channel_messages_sent counter
jitsi_total_data_channel_messages_sent 600
# HELP jitsi_total_colibri_web_socket_messages_received The total number messages received through COLIBRI web sockets.
# TYPE jitsi_total_colibri_web_socket_messages_received counter
jitsi_total_colibri_web_socket_messages_received 0
# HELP jitsi_total_colibri_web_socket_messages_sent The total number messages sent through COLIBRI web sockets.
# TYPE jitsi_total_colibri_web_socket_messages_sent counter
jitsi_total_colibri_web_socket_messages_sent 0
# HELP jitsi_octo_version The current running OCTO version
# TYPE jitsi_octo_version gauge
jitsi_octo_version 1
# HELP jitsi_octo_conferences The current number of OCTO conferences.
# TYPE jitsi_octo_conferences gauge
jitsi_octo_conferences 0
# HELP jitsi_octo_endpoints The current number of OCTO endpoints.
# TYPE jitsi_octo_endpoints gauge
jitsi_octo_endpoints 0
# HELP jitsi_octo_receive_bitrate The total receiving bitrate for the OCTO video bridge in kilobits per second.
# TYPE jitsi_octo_receive_bitrate gauge
jitsi_octo_receive_bitrate 0
# HELP jitsi_octo_send_bitrate The total outgoing bitrate for the OCTO video bridge in kilobits per second.
# TYPE jitsi_octo_send_bitrate gauge
jitsi_octo_send_bitrate 0
# HELP jitsi_octo_receive_packet_rate The total incoming packet rate for the OCTO video bridge in packets per second.
# TYPE jitsi_octo_receive_packet_rate gauge
jitsi_octo_receive_packet_rate 0
# HELP jitsi_octo_send_packet_rate The total outgoing packet rate for the OCTO video bridge in packets per second.
# TYPE jitsi_octo_send_packet_rate gauge
jitsi_octo_send_packet_rate 0
# HELP jitsi_total_bytes_received_octo The total incoming bit rate for the OCTO video bridge in bytes per second.
# TYPE jitsi_total_bytes_received_octo gauge
jitsi_total_bytes_received_octo 0
# HELP jitsi_total_bytes_sent_octo The total outgoing bit rate for the OCTO video bridge in bytes per second.
# TYPE jitsi_total_bytes_sent_octo gauge
jitsi_total_bytes_sent_octo 0
# HELP jitsi_total_packets_dropped_octo The total of dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_dropped_octo gauge
jitsi_total_packets_dropped_octo 0
# HELP jitsi_total_packets_received_octo The total of incoming dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_received_octo gauge
jitsi_total_packets_received_octo 0
# HELP jitsi_total_packets_sent_octo The total of sent dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_sent_octo gauge
jitsi_total_packets_sent_octo 0
# HELP jitsi_conference_sizes Distribution of conference sizes
# TYPE jitsi_conference_sizes gauge
jitsi_conference_sizes{conference_size="0"} 0
jitsi_conference_sizes{conference_size="1"} 0
jitsi_conference_sizes{conference_size="2"} 0
jitsi_conference_sizes{conference_size="3"} 0
jitsi_conference_sizes{conference_size="4"} 0
jitsi_conference_sizes{conference_size="5"} 0
jitsi_conference_sizes{conference_size="6"} 0
jitsi_conference_sizes{conference_size="7"} 0
jitsi_conference_sizes{conference_size="8"} 0
jitsi_conference_sizes{conference_size="9"} 0
jitsi_conference_sizes{conference_size="10"} 0
jitsi_conference_sizes{conference_size="11"} 0
jitsi_conference_sizes{conference_size="12"} 0
jitsi_conference_sizes{conference_size="13"} 0
jitsi_conference_sizes{conference_size="14"} 0
jitsi_conference_sizes{conference_size="15"} 0
jitsi_conference_sizes{conference_size="16"} 0
jitsi_conference_sizes{conference_size="17"} 0
jitsi_conference_sizes{conference_size="18"} 0
jitsi_conference_sizes{conference_size="19"} 0
jitsi_conference_sizes{conference_size="20"} 0
jitsi_conference_sizes{conference_size="21"} 0
`,
		},
	}

	for _, tc := range tcs {
		srv := httptest.NewServer(constHandler{tc.statsJson})

		h := handler{
			sourceURL: srv.URL,
		}
		req, err := http.NewRequest("GET", "/metrics", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		h.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if rr.Body.String() != tc.expected {
			t.Errorf("Response does not match the expected string:\n%s", cmp.Diff(rr.Body.String(), tc.expected))
		}

		srv.Close()
	}
}

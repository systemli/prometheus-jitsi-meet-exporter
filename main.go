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
	videoBridgeURL = flag.String("videobridge-url", "http://localhost:8080/colibri/stats", "Jitsi Videobridge /stats URL to scrape")
)

type videoBridgeStats struct {
	Threads                               int     `json:"threads"`
	BitRateDownload                       float64 `json:"bit_rate_download"`
	BitRateUpload                         float64 `json:"bit_rate_upload"`
	PacketRateDownload                    float64 `json:"packet_rate_download"`
	PacketRateUpload                      float64 `json:"packet_rate_upload"`
	LossRateDownload                      float64 `json:"loss_rate_download"`
	LossRateUpload                        float64 `json:"loss_rate_upload"`
	JitterAggregate                       float64 `json:"jitter_aggregate"`
	RTTAggregate                          float64 `json:"rtt_aggregate"`
	LargestConference                     int     `json:"largest_conference"`
	ConferenceSizes                       []int   `json:"conference_sizes"`
	AudioChannels                         int     `json:"audiochannels"`
	VideoChannels                         int     `json:"videochannels"`
	Conferences                           int     `json:"conferences"`
	P2PConferences                        int     `json:"p2p_conferences"`
	Participants                          int     `json:"participants"`
	Videostreams                          int     `json:"videostreams"`
	EndpointsSendingVideo                 int     `json:"endpoints_sending_video"`
	StressLevel                           float64 `json:"stress_level"`
	TotalLossControlledParticipantSeconds int     `json:"total_loss_controlled_participant_seconds"`
	TotalLossLimitedParticipantSeconds    int     `json:"total_loss_limited_participant_seconds"`
	TotalLossDegradedParticipantSeconds   int     `json:"total_loss_degraded_participant_seconds"`
	TotalConferenceSeconds                int     `json:"total_conference_seconds"`
	TotalConferencesCreated               int     `json:"total_conferences_created"`
	TotalConferencesCompleted             int     `json:"total_conferences_completed"`
	TotalFailedConferences                int     `json:"total_failed_conferences"`
	TotalPartiallyFailedConferences       int     `json:"total_partially_failed_conferences"`
	TotalDataChannelMessagesReceived      int     `json:"total_data_channel_messages_received"`
	TotalDataChannelMessagesSent          int     `json:"total_data_channel_messages_sent"`
	TotalColibriWebSocketMessagesReceived int     `json:"total_colibri_web_socket_messages_received"`
	TotalColibriWebSocketMessagesSent     int     `json:"total_colibri_web_socket_messages_sent"`
	TotalParticipants                     int     `json:"total_participants"`
	OctoVersion                           int     `json:"octo_version"`
	OctoConferences                       int     `json:"octo_conferences"`
	OctoEndpoints                         int     `json:"octo_endpoints"`
	OctoReceiveBitrate                    float64 `json:"octo_receive_bitrate"`
	OctoReceivePacketRate                 float64 `json:"octo_receive_packet_rate"`
	OctoSendBitrate                       float64 `json:"octo_send_bitrate"`
	OctoSendPacketRate                    float64 `json:"octo_send_packet_rate"`
	TotalBytesReceivedOcto                int     `json:"total_bytes_received_octo"`
	TotalBytesSentOcto                    int     `json:"total_bytes_sent_octo"`
	TotalPacketsDroppedOcto               int     `json:"total_packets_dropped_octo"`
	TotalPacketsReceivedOcto              int     `json:"total_packets_received_octo"`
	TotalPacketsSentOcto                  int     `json:"total_packets_sent_octo"`
	TotalICESucceededRelayed              int     `json:"total_ice_succeeded_relayed"`
	TotalICESucceeded                     int     `json:"total_ice_succeeded"`
	TotalICESucceededTCP                  int     `json:"total_ice_succeeded_tcp"`
	TotalICEFailed                        int     `json:"total_ice_failed"`
	EndpointsWithSuspendedSources         int     `json:"endpoints_with_suspended_sources"`
	InactiveEndpoints                     int     `json:"inactive_endpoints"`
	InactiveConferences                   int     `json:"inactive_conferences"`
	LocalActiveEndpoints                  int     `json:"local_active_endpoints"`
	MucClientsConnected                   int     `json:"muc_clients_connected"`
	LocalEndpoints                        int     `json:"local_endpoints"`
	TotalPacketsReceived                  int     `json:"total_packets_received"`
	PreemptiveKfrSuppressed               int     `json:"preemptive_kfr_suppressed"`
	PreemptiveKfrSent                     int     `json:"preemptive_kfr_sent"`
	TotalKeyframesReceived                int     `json:"total_keyframes_received"`
	TotalDominantSpeakerChanges           int     `json:"total_dominant_speaker_changes"`
	EndpointsWithSpuriousRemb             int     `json:"endpoints_with_spurious_remb"`
	ReceiveOnlyEndpoints                  int     `json:"receive_only_endpoints"`
	TotalVisitors                         int     `json:"total_visitors"`
	Version                               string  `json:"version"`
	Visitors                              int     `json:"visitors"`
	NumEpsOversending                     int     `json:"num_eps_oversending"`
	NumEpsNoMsgTransportAfterDelay        int     `json:"num_eps_no_msg_transport_after_delay"`
	Region                                string  `json:"region"`
	MucClientsConfigured                  int     `json:"muc_clients_configured"`
	OutgoingLoss                          float64 `json:"outgoing_loss"`
	OverallLoss                           float64 `json:"overall_loss"`
	TotalLayeringChangesReceived          int     `json:"total_layering_changes_received"`
	TotalRelays                           int     `json:"total_relays"`
	EndpointsWithHighOutgoingLoss         int     `json:"endpoints_with_high_outgoing_loss"`
	Drain                                 bool    `json:"drain"`
	TotalVideoStreamMillisecondsReceived  int     `json:"total_video_stream_milliseconds_received"`
	ShuttingDown                          bool    `json:"shutting_down"`
	CurrentTimestamp                      string  `json:"current_timestamp"`
	NumRelaysNoMsgTransportAfterDelay     int     `json:"num_relays_no_msg_transport_after_delay"`
	AverageParticipantStress              float64 `json:"average_participant_stress"`
	TotalPacketsSent                      int     `json:"total_packets_sent"`
	Endpoints                             int     `json:"endpoints"`
	IncomingLoss                          float64 `json:"incoming_loss"`
	EndpointsReconnected                  int     `json:"endpoints_reconnected"`
	GracefulShutdown                      bool    `json:"graceful_shutdown"`
	TotalBytesReceived                    int     `json:"total_bytes_received"`
	EndpointsDisconnected                 int     `json:"endpoints_disconnected"`
	EndpointsSendingAudio                 int     `json:"endpoints_sending_audio"`
	DTLSFailedEndpoints                   int     `json:"dtls_failed_endpoints"`
	TotalBytesSent                        int     `json:"total_bytes_sent"`
	Healthy                               bool    `json:"healthy"`
	MucsConfigured                        int     `json:"mucs_configured"`
	MucsJoined                            int     `json:"mucs_joined"`
	RelayID                               string  `json:"relay_id"`
}

var tpl = template.Must(template.New("stats").Parse(`# HELP jitsi_threads The number of Java threads that the video bridge is using.
# TYPE jitsi_threads gauge
jitsi_threads {{.Threads}}
# HELP jitsi_bit_rate_download The total incoming bitrate for the video bridge in kilobits per second.
# TYPE jitsi_bit_rate_download gauge
jitsi_bit_rate_download {{.BitRateDownload}}
# HELP jitsi_bit_rate_upload The total outgoing bitrate for the video bridge in kilobits per second.
# TYPE jitsi_bit_rate_upload gauge
jitsi_bit_rate_upload {{.BitRateUpload}}
# HELP jitsi_packet_rate_download The total incoming packet rate for the video bridge in packets per second.
# TYPE jitsi_packet_rate_download gauge
jitsi_packet_rate_download {{.PacketRateDownload}}
# HELP jitsi_packet_rate_upload The total outgoing packet rate for the video bridge in packets per second.
# TYPE jitsi_packet_rate_upload gauge
jitsi_packet_rate_upload {{.PacketRateUpload}}
# HELP jitsi_loss_rate_download The fraction of lost incoming RTP packets. This is based on RTP sequence numbers and is relatively accurate.
# TYPE jitsi_loss_rate_download gauge
jitsi_loss_rate_download {{.LossRateDownload}}
# HELP jitsi_loss_rate_upload The fraction of lost outgoing RTP packets. This is based on incoming RTCP Receiver Reports, and an attempt to subtract the fraction of packets that were not sent (i.e. were lost before they reached the bridge). Further, this is averaged over all streams of all users as opposed to all packets, so it is not correctly weighted. This is not accurate, but may be a useful metric nonetheless.
# TYPE jitsi_loss_rate_upload gauge
jitsi_loss_rate_upload {{.LossRateUpload}}
# HELP jitsi_jitter_aggregate Experimental. An average value (in milliseconds) of the jitter calculated for incoming and outgoing streams. This hasn't been tested and it is currently not known whether the values are correct or not.
# TYPE jitsi_jitter_aggregate gauge
jitsi_jitter_aggregate {{.JitterAggregate}}
# HELP jitsi_rtt_aggregate An average value (in milliseconds) of the RTT across all streams.
# TYPE jitsi_rtt_aggregate gauge
jitsi_rtt_aggregate {{.RTTAggregate}}
# HELP jitsi_largest_conference The number of participants in the largest conference currently hosted on the bridge.
# TYPE jitsi_largest_conference gauge
jitsi_largest_conference {{.LargestConference}}
# HELP jitsi_audiochannels The current number of audio channels.
# TYPE jitsi_audiochannels gauge
jitsi_audiochannels {{.AudioChannels}}
# HELP jitsi_videochannels The current number of video channels.
# TYPE jitsi_videochannels gauge
jitsi_videochannels {{.VideoChannels}}
# HELP jitsi_conferences The current number of conferences.
# TYPE jitsi_conferences gauge
jitsi_conferences {{.Conferences}}
# HELP jitsi_p2p_conferences The current number of p2p conferences.
# TYPE jitsi_p2p_conferences gauge
jitsi_p2p_conferences {{.P2PConferences}}
# HELP jitsi_participants The current number of participants.
# TYPE jitsi_participants gauge
jitsi_participants {{.Participants}}
# HELP jitsi_total_participants Total participants since running.
# TYPE jitsi_total_participants gauge
jitsi_total_participants {{.TotalParticipants}}
# HELP jitsi_endpoint_sending_video An estimation of the number of current endpoints sending a video stream.
# TYPE jitsi_endpoint_sending_video gauge
jitsi_endpoints_sending_video {{.EndpointsSendingVideo}}
# HELP jitsi_videostreams An estimation of the number of current video streams forwarded by the bridge.
# TYPE jitsi_videostreams gauge
jitsi_videostreams {{.Videostreams}}
# HELP jitsi_stress_level Stress Level reported to Jicofo by the videobridge.
# TYPE jitsi_stress_level gauge
jitsi_stress_level {{.StressLevel}}
# HELP jitsi_total_loss_controlled_participant_seconds The total number of participant-seconds that are loss-controlled.
# TYPE jitsi_total_loss_controlled_participant_seconds counter
jitsi_total_loss_controlled_participant_seconds {{.TotalLossControlledParticipantSeconds}}
# HELP jitsi_total_loss_limited_participant_seconds The total number of participant-seconds that are loss-limited.
# TYPE jitsi_total_loss_limited_participant_seconds counter
jitsi_total_loss_limited_participant_seconds {{.TotalLossLimitedParticipantSeconds}}
# HELP jitsi_total_loss_degraded_participant_seconds The total number of participant-seconds that are loss-degraded.
# TYPE jitsi_total_loss_degraded_participant_seconds counter
jitsi_total_loss_degraded_participant_seconds {{.TotalLossDegradedParticipantSeconds}}
# HELP jitsi_total_conference_seconds The sum of the lengths of all completed conferences, in seconds.
# TYPE jitsi_total_conference_seconds counter
jitsi_total_conference_seconds {{.TotalConferenceSeconds}}
# HELP jitsi_total_conferences_created The total number of conferences created on the bridge.
# TYPE jitsi_total_conferences_created counter
jitsi_total_conferences_created {{.TotalConferencesCreated}}
# HELP jitsi_total_conferences_completed The total number of conferences completed on the bridge.
# TYPE jitsi_total_conferences_completed counter
jitsi_total_conferences_completed {{.TotalConferencesCompleted}}
# HELP jitsi_total_failed_conferences The total number of failed conferences on the bridge. A conference is marked as failed when all of its channels have failed. A channel is marked as failed if it had no payload activity.
# TYPE jitsi_total_failed_conferences counter
jitsi_total_failed_conferences {{.TotalFailedConferences}}
# HELP jitsi_total_partially_failed_conferences The total number of partially failed conferences on the bridge. A conference is marked as partially failed when some of its channels has failed. A channel is marked as failed if it had no payload activity.
# TYPE jitsi_total_partially_failed_conferences counter
jitsi_total_partially_failed_conferences {{.TotalPartiallyFailedConferences}}
# HELP jitsi_total_data_channel_messages_received The total number messages received through data channels.
# TYPE jitsi_total_data_channel_messages_received counter
jitsi_total_data_channel_messages_received {{.TotalDataChannelMessagesReceived}}
# HELP jitsi_total_data_channel_messages_sent The total number messages sent through data channels.
# TYPE jitsi_total_data_channel_messages_sent counter
jitsi_total_data_channel_messages_sent {{.TotalDataChannelMessagesSent}}
# HELP jitsi_total_colibri_web_socket_messages_received The total number messages received through COLIBRI web sockets.
# TYPE jitsi_total_colibri_web_socket_messages_received counter
jitsi_total_colibri_web_socket_messages_received {{.TotalColibriWebSocketMessagesReceived}}
# HELP jitsi_total_colibri_web_socket_messages_sent The total number messages sent through COLIBRI web sockets.
# TYPE jitsi_total_colibri_web_socket_messages_sent counter
jitsi_total_colibri_web_socket_messages_sent {{.TotalColibriWebSocketMessagesSent}}
# HELP jitsi_octo_version The current running OCTO version
# TYPE jitsi_octo_version gauge
jitsi_octo_version {{.OctoVersion}}
# HELP jitsi_octo_conferences The current number of OCTO conferences.
# TYPE jitsi_octo_conferences gauge
jitsi_octo_conferences {{.OctoConferences}}
# HELP jitsi_octo_endpoints The current number of OCTO endpoints.
# TYPE jitsi_octo_endpoints gauge
jitsi_octo_endpoints {{.OctoEndpoints}}
# HELP jitsi_octo_receive_bitrate The total receiving bitrate for the OCTO video bridge in kilobits per second.
# TYPE jitsi_octo_receive_bitrate gauge
jitsi_octo_receive_bitrate {{.OctoReceiveBitrate}}
# HELP jitsi_octo_send_bitrate The total outgoing bitrate for the OCTO video bridge in kilobits per second.
# TYPE jitsi_octo_send_bitrate gauge
jitsi_octo_send_bitrate {{.OctoSendBitrate}}
# HELP jitsi_octo_receive_packet_rate The total incoming packet rate for the OCTO video bridge in packets per second.
# TYPE jitsi_octo_receive_packet_rate gauge
jitsi_octo_receive_packet_rate {{.OctoReceivePacketRate}}
# HELP jitsi_octo_send_packet_rate The total outgoing packet rate for the OCTO video bridge in packets per second.
# TYPE jitsi_octo_send_packet_rate gauge
jitsi_octo_send_packet_rate {{.OctoSendPacketRate}}
# HELP jitsi_total_bytes_received_octo The total incoming bit rate for the OCTO video bridge in bytes per second.
# TYPE jitsi_total_bytes_received_octo gauge
jitsi_total_bytes_received_octo {{.TotalBytesReceivedOcto}}
# HELP jitsi_total_bytes_sent_octo The total outgoing bit rate for the OCTO video bridge in bytes per second.
# TYPE jitsi_total_bytes_sent_octo gauge
jitsi_total_bytes_sent_octo {{.TotalBytesSentOcto}}
# HELP jitsi_total_packets_dropped_octo The total of dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_dropped_octo gauge
jitsi_total_packets_dropped_octo {{.TotalPacketsDroppedOcto}}
# HELP jitsi_total_packets_received_octo The total of incoming dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_received_octo gauge
jitsi_total_packets_received_octo {{.TotalPacketsReceivedOcto}}
# HELP jitsi_total_packets_sent_octo The total of sent dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_sent_octo gauge
jitsi_total_packets_sent_octo {{.TotalPacketsSentOcto}}
# HELP total_ice_succeeded_relayed The total number of times an ICE Agent succeeded and the selected candidate pair included a relayed candidate.
# TYPE total_ice_succeeded_relayed gauge
total_ice_succeeded_relayed {{.TotalICESucceededRelayed}}
# HELP total_ice_succeeded The total number of times an ICE Agent succeeded.
# TYPE total_ice_succeeded gauge
total_ice_succeeded {{.TotalICESucceeded}}
# HELP total_ice_succeeded_tcp The total number of times an ICE Agent succeeded and the selected candidate was a TCP candidate.
# TYPE total_ice_succeeded_tcp gauge
total_ice_succeeded_tcp {{.TotalICESucceededTCP}}
# HELP total_ice_failed The total number of times an ICE Agent failed to establish connectivity.
# TYPE total_ice_failed gauge
total_ice_failed {{.TotalICEFailed}}
# HELP jitsi_endpoints_with_suspended_sources Number of endpoints that we have suspended sending some video streams to because of bwe.
# TYPE jitsi_endpoints_with_suspended_sources gauge
jitsi_endpoints_with_suspended_sources {{.EndpointsWithSuspendedSources}}
# HELP jitsi_inactive_endpoints Number of endpoints in inactive conferences (where no endpoint sends audio or video).
# TYPE jitsi_inactive_endpoints gauge
jitsi_inactive_endpoints {{.InactiveEndpoints}}
# HELP jitsi_inactive_conferences Number of inactive conferences (no endpoint is sending audio or video).
# TYPE jitsi_inactive_conferences gauge
jitsi_inactive_conferences {{.InactiveConferences}}
# HELP jitsi_local_active_endpoints The number of active local endpoints (in a conference where at least one endpoint sends audio or video).
# TYPE jitsi_local_active_endpoints gauge
jitsi_local_active_endpoints {{.LocalActiveEndpoints}}
# HELP jitsi_muc_clients_connected The current number of connected XMPP MUC clients.
# TYPE jitsi_muc_clients_connected gauge
jitsi_muc_clients_connected {{.MucClientsConnected}}
# HELP jitsi_local_endpoints The current number of local non-OCTO endpoints.
# TYPE jitsi_local_endpoints gauge
jitsi_local_endpoints {{.LocalEndpoints}}
# HELP jitsi_total_packets_received The total number of RTP packets received.
# TYPE jitsi_total_packets_received gauge
jitsi_total_packets_received {{.TotalPacketsReceived}}
# HELP jitsi_preemptive_kfr_suppressed Number of preemptive keyframe requests that were not sent because no endpoints were in stage view.
# TYPE jitsi_preemptive_kfr_suppressed gauge
jitsi_preemptive_kfr_suppressed {{.PreemptiveKfrSuppressed}}
# HELP jitsi_preemptive_kfr_sent Number of preemptive keyframe requests that were sent.
# TYPE jitsi_preemptive_kfr_sent gauge
jitsi_preemptive_kfr_sent {{.PreemptiveKfrSent}}
# HELP jitsi_total_keyframes_received Number of keyframes that were received (updated on endpoint expiration).
# TYPE jitsi_total_keyframes_received gauge
jitsi_total_keyframes_received {{.TotalKeyframesReceived}}
# HELP jitsi_total_dominant_speaker_changes The total number of dominant speaker changes.
# TYPE jitsi_total_dominant_speaker_changes gauge
jitsi_total_dominant_speaker_changes {{.TotalDominantSpeakerChanges}}
# HELP jitsi_endpoints_with_spurious_remb Number of endpoints that have sent a REMB packet even though REMB was not configured.
# TYPE jitsi_endpoints_with_spurious_remb gauge
jitsi_endpoints_with_spurious_remb {{.EndpointsWithSpuriousRemb}}
# HELP jitsi_receive_only_endpoints Number of endpoints that are not sending audio or video (but are receiving).
# TYPE jitsi_receive_only_endpoints gauge
jitsi_receive_only_endpoints {{.ReceiveOnlyEndpoints}}
# HELP jitsi_total_visitors The total number of visitors since startup.
# TYPE jitsi_total_visitors gauge
jitsi_total_visitors {{.TotalVisitors}}
# HELP jitsi_visitors The total number of visitor endpoints created.
# TYPE jitsi_visitors gauge
jitsi_visitors {{.Visitors}}
# HELP jitsi_num_eps_oversending Number of endpoints that we are oversending to.
# TYPE jitsi_num_eps_oversending gauge
jitsi_num_eps_oversending {{.NumEpsOversending}}
# HELP jitsi_num_eps_no_msg_transport_after_delay Number of endpoints which had not established a relay message transport even after some delay.
# TYPE jitsi_num_eps_no_msg_transport_after_delay gauge
jitsi_num_eps_no_msg_transport_after_delay {{.NumEpsNoMsgTransportAfterDelay}}
# HELP jitsi_muc_clients_configured The number of configured XMPP MUC clients.
# TYPE jitsi_muc_clients_configured gauge
jitsi_muc_clients_configured {{.MucClientsConfigured}}
# HELP jitsi_outgoing_loss Fraction of outgoing RTP packets that are lost.
# TYPE jitsi_outgoing_loss gauge
jitsi_outgoing_loss {{.OutgoingLoss}}
# HELP jitsi_overall_loss Fraction of RTP packets that are lost (incoming and outgoing combined).
# TYPE jitsi_overall_loss gauge
jitsi_overall_loss {{.OverallLoss}}
# HELP jitsi_total_layering_changes_received Number of times the layering of an incoming video stream changed (updated on endpoint expiration).
# TYPE jitsi_total_layering_changes_received gauge
jitsi_total_layering_changes_received {{.TotalLayeringChangesReceived}}
# HELP jitsi_total_relays The total number of relays created.
# TYPE jitsi_total_relays gauge
jitsi_total_relays {{.TotalRelays}}
# HELP jitsi_endpoints_with_high_outgoing_loss Number of endpoints that have high outgoing loss (>10%).
# TYPE jitsi_endpoints_with_high_outgoing_loss gauge
jitsi_endpoints_with_high_outgoing_loss {{.EndpointsWithHighOutgoingLoss}}
# HELP jitsi_drain Whether the bridge is draining and should avoid new conference allocation.
# TYPE jitsi_drain gauge
jitsi_drain {{if .Drain}}1{{else}}0{{end}}
# HELP jitsi_total_video_stream_milliseconds_received Total duration of video received, in milliseconds (each SSRC counts separately).
# TYPE jitsi_total_video_stream_milliseconds_received gauge
jitsi_total_video_stream_milliseconds_received {{.TotalVideoStreamMillisecondsReceived}}
# HELP jitsi_shutting_down Whether jitsi-videobridge is shutting down.
# TYPE jitsi_shutting_down gauge
jitsi_shutting_down {{if .ShuttingDown}}1{{else}}0{{end}}
# HELP jitsi_num_relays_no_msg_transport_after_delay Number of relays which had not established a relay message transport even after some delay.
# TYPE jitsi_num_relays_no_msg_transport_after_delay gauge
jitsi_num_relays_no_msg_transport_after_delay {{.NumRelaysNoMsgTransportAfterDelay}}
# HELP jitsi_average_participant_stress Average participant stress reported by the bridge.
# TYPE jitsi_average_participant_stress gauge
jitsi_average_participant_stress {{.AverageParticipantStress}}
# HELP jitsi_total_packets_sent The total number of RTP packets sent.
# TYPE jitsi_total_packets_sent gauge
jitsi_total_packets_sent {{.TotalPacketsSent}}
# HELP jitsi_endpoints Number of current endpoints (local and relayed).
# TYPE jitsi_endpoints gauge
jitsi_endpoints {{.Endpoints}}
# HELP jitsi_incoming_loss Fraction of incoming RTP packets that are lost.
# TYPE jitsi_incoming_loss gauge
jitsi_incoming_loss {{.IncomingLoss}}
# HELP jitsi_endpoints_reconnected Endpoints reconnected after being detected as temporarily inactive/disconnected due to inactivity.
# TYPE jitsi_endpoints_reconnected gauge
jitsi_endpoints_reconnected {{.EndpointsReconnected}}
# HELP jitsi_graceful_shutdown Whether jitsi-videobridge is in graceful shutdown mode.
# TYPE jitsi_graceful_shutdown gauge
jitsi_graceful_shutdown {{if .GracefulShutdown}}1{{else}}0{{end}}
# HELP jitsi_total_bytes_received The total number of RTP bytes received.
# TYPE jitsi_total_bytes_received gauge
jitsi_total_bytes_received {{.TotalBytesReceived}}
# HELP jitsi_endpoints_disconnected Endpoints detected as temporarily inactive/disconnected due to inactivity.
# TYPE jitsi_endpoints_disconnected gauge
jitsi_endpoints_disconnected {{.EndpointsDisconnected}}
# HELP jitsi_endpoints_sending_audio The number of local endpoints sending audio.
# TYPE jitsi_endpoints_sending_audio gauge
jitsi_endpoints_sending_audio {{.EndpointsSendingAudio}}
# HELP jitsi_dtls_failed_endpoints The total number of endpoints that failed to establish DTLS.
# TYPE jitsi_dtls_failed_endpoints gauge
jitsi_dtls_failed_endpoints {{.DTLSFailedEndpoints}}
# HELP jitsi_total_bytes_sent The total number of RTP bytes sent.
# TYPE jitsi_total_bytes_sent gauge
jitsi_total_bytes_sent {{.TotalBytesSent}}
# HELP jitsi_healthy Whether this bridge currently reports itself as healthy.
# TYPE jitsi_healthy gauge
jitsi_healthy {{if .Healthy}}1{{else}}0{{end}}
# HELP jitsi_mucs_configured The number of configured MUCs.
# TYPE jitsi_mucs_configured gauge
jitsi_mucs_configured {{.MucsConfigured}}
# HELP jitsi_mucs_joined The number of joined MUCs.
# TYPE jitsi_mucs_joined gauge
jitsi_mucs_joined {{.MucsJoined}}
# HELP jitsi_version_info Jitsi Videobridge version information.
# TYPE jitsi_version_info gauge
jitsi_version_info{version="{{.Version}}"} 1
# HELP jitsi_region_info Configured bridge region information.
# TYPE jitsi_region_info gauge
jitsi_region_info{region="{{.Region}}"} 1
# HELP jitsi_relay_id_info Relay identifier information.
# TYPE jitsi_relay_id_info gauge
jitsi_relay_id_info{relay_id="{{.RelayID}}"} 1
# HELP jitsi_conference_sizes Distribution of conference sizes
# TYPE jitsi_conference_sizes gauge
{{ range $key, $value := .ConferenceSizes -}}
jitsi_conference_sizes{conference_size="{{$key}}"} {{ $value }}
{{ end -}}

`))

type handler struct {
	sourceURL string
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get(h.sourceURL)
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

	http.Handle("/metrics", handler{sourceURL: *videoBridgeURL})
	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte(`ok`))
	})
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Started Jitsi Meet Metrics Exporter")
}

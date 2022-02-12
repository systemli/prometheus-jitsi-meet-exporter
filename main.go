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
	AverageParticipantStress              float64 `json:"average_participant_stress"`
	BitRateDownload                       float64 `json:"bit_rate_download"`
	BitRateUpload                         float64 `json:"bit_rate_upload"`
	ConferenceSizes                       []int   `json:"conference_sizes"`
	Conferences                           int     `json:"conferences"`
	ConferenceAudioSizes                  []int   `json:"conferences_by_audio_senders"`
	ConferenceVideoSizes                  []int   `json:"conferences_by_video_senders"`
	DTLSFailedEndpoints                   int     `json:"dtls_failed_endpoints"`
	Endpoints                             int     `json:"endpoints"`
	EndpointsSendingAudio                 int     `json:"endpoints_sending_audio"`
	EndpointsSendingVideo                 int     `json:"endpoints_sending_video"`
	EndpointsWithHighOutgoingLoss         int     `json:"endpoints_with_high_outgoing_loss"`
	EndpointsWithSpuriousREMB             int     `json:"endpoints_with_spurious_remb"`
	InactiveConferences                   int     `json:"inactive_conferences"`
	InactiveEndpoints                     int     `json:"inactive_endpoints"`
	IncomingLoss                          float64 `json:"incoming_loss"`
	JitterAggregate                       float64 `json:"jitter_aggregate"`
	LargestConference                     int     `json:"largest_conference"`
	LocalActiveEndpoints                  int     `json:"local_active_endpoints"`
	LocalEndpoints                        int     `json:"local_endpoints"`
	MucClientsConfigured                  int     `json:"muc_clients_configured"`
	MucClientsConnected                   int     `json:"muc_clients_connected"`
	MucsConfigured                        int     `json:"mucs_configured"`
	MucsJoined                            int     `json:"mucs_joined"`
	NumEndpointsOversending               int     `json:"num_eps_oversending"`
	NumEndpointsNoMsgTransportAfterDelay  int     `json:"num_eps_no_msg_transport_after_delay"`
	OctoConferences                       int     `json:"octo_conferences"`
	OctoEndpoints                         int     `json:"octo_endpoints"`
	OctoReceiveBitrate                    float64 `json:"octo_receive_bitrate"`
	OctoReceivePacketRate                 int     `json:"octo_receive_packet_rate"`
	OctoSendBitrate                       float64 `json:"octo_send_bitrate"`
	OctoSendPacketRate                    int     `json:"octo_send_packet_rate"`
	OctoVersion                           int     `json:"octo_version"`
	OutgoingLoss                          float64 `json:"outgoing_loss"`
	OverallLoss                           float64 `json:"overall_loss"`
	P2PConferences                        int     `json:"p2p_conferences"`
	PacketRateDownload                    int     `json:"packet_rate_download"`
	PacketRateUpload                      int     `json:"packet_rate_upload"`
	Participants                          int     `json:"participants"`
	PreEmptiveKfrSent                     int     `json:"preemptive_kfr_sent"`
	PreEmptiveKfrSuppressed               int     `json:"preemptive_kfr_suppressed"`
	ReceiveOnlyEndpoints                  int     `json:"receive_only_endpoints"`
	RTTAggregate                          float64 `json:"rtt_aggregate"`
	StressLevel                           float64 `json:"stress_level"`
	TotalAimdBweExpirations               int     `json:"total_aimd_bwe_expirations"`
	TotalBytesReceived                    int     `json:"total_bytes_received"`
	TotalBytesReceivedOcto                int     `json:"total_bytes_received_octo"`
	TotalBytesSent                        int     `json:"total_bytes_sent"`
	TotalBytesSentOcto                    int     `json:"total_bytes_sent_octo"`
	TotalColibriWebSocketMessagesReceived int     `json:"total_colibri_web_socket_messages_received"`
	TotalColibriWebSocketMessagesSent     int     `json:"total_colibri_web_socket_messages_sent"`
	TotalConferenceSeconds                int     `json:"total_conference_seconds"`
	TotalConferencesCompleted             int     `json:"total_conferences_completed"`
	TotalConferencesCreated               int     `json:"total_conferences_created"`
	TotalDataChannelMessagesReceived      int     `json:"total_data_channel_messages_received"`
	TotalDataChannelMessagesSent          int     `json:"total_data_channel_messages_sent"`
	TotalDominantSpeakerChanges           int     `json:"total_dominant_speaker_changes"`
	TotalFailedConferences                int     `json:"total_failed_conferences"`
	TotalIceFailed                        int     `json:"total_ice_failed"`
	TotalIceSucceeded                     int     `json:"total_ice_succeeded"`
	TotalIceSucceededRelayed              int     `json:"total_ice_succeeded_relayed"`
	TotalIceSucceededTCP                  int     `json:"total_ice_succeeded_tcp"`
	TotalKeyframesReceived                int     `json:"total_keyframes_received"`
	TotalLayeringChangesReceived          int     `json:"total_layering_changes_received"`
	TotalLossControlledParticipantSeconds int     `json:"total_loss_controlled_participant_seconds"`
	TotalLossDegradedParticipantSeconds   int     `json:"total_loss_degraded_participant_seconds"`
	TotalLossLimitedParticipantSeconds    int     `json:"total_loss_limited_participant_seconds"`
	TotalPacketsDroppedOcto               int     `json:"total_packets_dropped_octo"`
	TotalPacketsReceived                  int     `json:"total_packets_received"`
	TotalPacketsReceivedOcto              int     `json:"total_packets_received_octo"`
	TotalPacketsSent                      int     `json:"total_packets_sent"`
	TotalPacketsSentOcto                  int     `json:"total_packets_sent_octo"`
	TotalPartiallyFailedConferences       int     `json:"total_partially_failed_conferences"`
	TotalParticipants                     int     `json:"total_participants"`
	TotalVideoStreamMilliSecReceived      int     `json:"total_video_stream_milliseconds_received"`
	VideoChannels                         int     `json:"videochannels"`
}

var tpl = template.Must(template.New("stats").Parse(`# HELP jitsi_threads The number of Java threads that the video bridge is using.
# TYPE jitsi_threads gauge
jitsi_threads {{.Threads}}
# HELP jitsi_average_participant_stress The average load of the videobrige according to the stress per user set in, and therfore reported by, jicofo.
# TYPE jitsi_average_participant_stress gauge
jitsi_average_participant_stress {{.AverageParticipantStress}}
# HELP jitsi_bit_rate_download The total incoming bitrate for the video bridge in kilobits per second.
# TYPE jitsi_bit_rate_download gauge
jitsi_bit_rate_download {{.BitRateDownload}}
# HELP jitsi_bit_rate_upload The total outgoing bitrate for the video bridge in kilobits per second.
# TYPE jitsi_bit_rate_upload gauge
jitsi_bit_rate_upload {{.BitRateUpload}}
# HELP jitsi_conference_sizes Distribution of conference sizes.
# TYPE jitsi_conference_sizes gauge
{{ range $key, $value := .ConferenceSizes -}}
jitsi_conference_sizes{conference_size="{{$key}}"} {{ $value }}
{{ end -}}
# HELP jitsi_conferences The current number of conferences.
# TYPE jitsi_conferences gauge
jitsi_conferences {{.Conferences}}
# HELP jitsi_conferences_by_audio_senders Distribution of conference sizes by audio senders.
# TYPE jitsi_conferences_by_audio_senders gauge
{{ range $key, $value := .ConferenceAudioSizes -}}
jitsi_conferences_by_audio_senders{conferences_by_audio="{{$key}}"} {{ $value }}
{{ end -}}
# HELP jitsi_conferences_by_video_senders Distribution of conference sizes by video senders.
# TYPE jitsi_conferences_by_video_senders gauge
{{ range $key, $value := .ConferenceVideoSizes -}}
jitsi_conferences_by_video_senders{conferences_by_video="{{$key}}"} {{ $value }}
{{ end -}}
# HELP jitsi_dtls_failed_endpoints The number of endpoints with failed Datagram Transport Layer Security connection.
# TYPE jitsi_dtls_failed_endpoints gauge
jitsi_dtls_failed_endpoints {{.DTLSFailedEndpoints}}
# HELP jitsi_endpoints The number of active endpoints.
# TYPE jitsi_endpoints gauge
jitsi_endpoints {{.Endpoints}}
# HELP jitsi_endpoints_sending_audio The number of active endpoints sending audio.
# TYPE jitsi_endpoints_sending_audio gauge
jitsi_endpoints_sending_audio {{.EndpointsSendingAudio}}
# HELP jitsi_endpoints_sending_video The number of active endpoints sending video.
# TYPE jitsi_endpoints_sending_video gauge
jitsi_endpoints_sending_video {{.EndpointsSendingVideo}}
# HELP jitsi_endpoints_with_high_outgoing_loss The number of endpoints with high outgoing loss.
# TYPE jitsi_endpoints_with_high_outgoing_loss gauge
jitsi_endpoints_with_high_outgoing_loss {{.EndpointsWithHighOutgoingLoss}}
# HELP jitsi_endpoints_with_spurious_remb The total number of endpoints which have sent an RTCP REMB packet when REMB was not signaled.
# TYPE jitsi_endpoints_with_spurious_remb gauge
jitsi_endpoints_with_spurious_remb {{.EndpointsWithSpuriousREMB}}
# HELP jitsi_inactive_conferences The number of conferences in which no endpoints are sending audio nor video. Note that this includes conferences which are currently using a peer-to-peer transport. 
# TYPE jitsi_inactive_conferences gauge
jitsi_inactive_conferences {{.InactiveConferences}}
# HELP jitsi_inactive_endpoints The number endpoints sending no audio and no video. 
# TYPE jitsi_inactive_endpoints gauge
jitsi_inactive_endpoints {{.InactiveEndpoints}}
# HELP jitsi_incoming_loss The average value (in percent) of lost packages from incoming connections.
# TYPE jitsi_incoming_loss gauge
jitsi_incoming_loss {{.IncomingLoss}}
# HELP jitsi_jitter_aggregate Experimental. An average value (in milliseconds) of the jitter calculated for incoming and outgoing streams.
# TYPE jitsi_jitter_aggregate gauge
jitsi_jitter_aggregate {{.JitterAggregate}}
# HELP jitsi_largest_conference The number of participants in the largest conference currently hosted on the bridge.
# TYPE jitsi_largest_conference gauge
jitsi_largest_conference {{.LargestConference}}
# HELP jitsi_local_active_endpoints The number of local endpoints (not octo) which are in an active conference. This includes endpoints which are not sending audio or video, but are in an active conference (i.e. they are receive-only).
# TYPE jitsi_local_active_endpoints gauge
jitsi_local_active_endpoints {{.LocalActiveEndpoints}}
# HELP jitsi_local_endpoints The number of local (non-octo) endpoints.
# TYPE jitsi_local_endpoints gauge
jitsi_local_endpoints {{.LocalEndpoints}}
# HELP jitsi_muc_clients_configured The number of muc clients configured on the JVB.
# TYPE jitsi_muc_clients_configured gauge
jitsi_muc_clients_configured {{.MucClientsConfigured}}
# HELP jitsi_muc_clients_connected The number of muc clients connected to the JVB.
# TYPE jitsi_muc_clients_connected gauge
jitsi_muc_clients_connected {{.MucClientsConnected}}
# HELP jitsi_mucs_configured The number of mucs configured on the JVB.
# TYPE jitsi_mucs_configured gauge
jitsi_mucs_configured {{.MucsConfigured}}
# HELP jitsi_mucs_joined The number of mucs the JVB joined.
# TYPE jitsi_mucs_joined gauge
jitsi_mucs_joined {{.MucsJoined}}
# HELP jitsi_num_eps_oversending The number of endpoints to which we are oversending.
# TYPE jitsi_num_eps_oversending gauge
jitsi_num_eps_oversending {{.NumEndpointsOversending}}
# HELP jitsi_num_eps_no_msg_transport_after_delay No help found.
# TYPE jitsi_num_eps_no_msg_transport_after_delay gauge
jitsi_num_eps_no_msg_transport_after_delay {{.NumEndpointsNoMsgTransportAfterDelay}}
# HELP jitsi_octo_conferences The current number of OCTO conferences.
# TYPE jitsi_octo_conferences gauge
jitsi_octo_conferences {{.OctoConferences}}
# HELP jitsi_octo_endpoints The current number of OCTO endpoints.
# TYPE jitsi_octo_endpoints gauge
jitsi_octo_endpoints {{.OctoEndpoints}}
# HELP jitsi_octo_receive_bitrate The total receiving bitrate for the OCTO video bridge in kilobits per second.
# TYPE jitsi_octo_receive_bitrate gauge
jitsi_octo_receive_bitrate {{.OctoReceiveBitrate}}
# HELP jitsi_octo_receive_packet_rate The total incoming packet rate for the OCTO video bridge in packets per second.
# TYPE jitsi_octo_receive_packet_rate gauge
jitsi_octo_receive_packet_rate {{.OctoReceivePacketRate}}
# HELP jitsi_octo_send_bitrate The total outgoing bitrate for the OCTO video bridge in kilobits per second.
# TYPE jitsi_octo_send_bitrate gauge
jitsi_octo_send_bitrate {{.OctoSendBitrate}}
# HELP jitsi_octo_send_packet_rate The total outgoing packet rate for the OCTO video bridge in packets per second.
# TYPE jitsi_octo_send_packet_rate gauge
jitsi_octo_send_packet_rate {{.OctoSendPacketRate}}
# HELP jitsi_octo_version The current running OCTO version
# TYPE jitsi_octo_version gauge
jitsi_octo_version {{.OctoVersion}}
# HELP jitsi_outgoing_loss The average value (in percent) of lost packages from outgoing connections.
# TYPE jitsi_outgoing_loss gauge
jitsi_outgoing_loss {{.OutgoingLoss}}
# HELP jitsi_overall_loss The average value (in percent) of lost packages combined from incoming and outgoing connections.
# TYPE jitsi_overall_loss gauge
jitsi_overall_loss {{.OverallLoss}}
# HELP jitsi_p2p_conferences The current number of p2p conferences.
# TYPE jitsi_p2p_conferences gauge
jitsi_p2p_conferences {{.P2PConferences}}
# HELP jitsi_packet_rate_download The total incoming packet rate for the video bridge in packets per second.
# TYPE jitsi_packet_rate_download gauge
jitsi_packet_rate_download {{.PacketRateDownload}}
# HELP jitsi_packet_rate_upload The total outgoing packet rate for the video bridge in packets per second.
# TYPE jitsi_packet_rate_upload gauge
jitsi_packet_rate_upload {{.PacketRateUpload}}
# HELP jitsi_participants The current number of participants.
# TYPE jitsi_participants gauge
jitsi_participants {{.Participants}}
# HELP jitsi_preemptive_kfr_sent The number of preemptive keyframe requests sent.
# TYPE jitsi_preemptive_kfr_sent gauge
jitsi_preemptive_kfr_sent {{.PreEmptiveKfrSent}}
# HELP jitsi_preemptive_kfr_suppressed The number of preemptive keyframe requests suppressed.
# TYPE jitsi_preemptive_kfr_suppressed gauge
jitsi_preemptive_kfr_suppressed {{.PreEmptiveKfrSuppressed}}
# HELP jitsi_receive_only_endpoints The number of endpoints which are not sending audio nor video.
# TYPE jitsi_receive_only_endpoints gauge
jitsi_receive_only_endpoints {{.ReceiveOnlyEndpoints}}
# HELP jitsi_rtt_aggregate An average value (in milliseconds) of the RTT across all streams.
# TYPE jitsi_rtt_aggregate gauge
jitsi_rtt_aggregate {{.RTTAggregate}}
# HELP jitsi_stress_level Stress Level reported to Jicofo by the videobridge.
# TYPE jitsi_stress_level gauge
jitsi_stress_level {{.StressLevel}}
# HELP jitsi_total_aimd_bwe_expirations The name of the stat that tracks the total number of times our AIMDs have expired the incoming bitrate (and which would otherwise result in video suspension).
# TYPE jitsi_total_aimd_bwe_expirations counter
jitsi_total_aimd_bwe_expirations {{.TotalAimdBweExpirations}}
# HELP jitsi_total_bytes_received The total number of bytes received by the JVB  since it started.
# TYPE jitsi_total_bytes_received counter
jitsi_total_bytes_received {{.TotalBytesReceived}}
# HELP jitsi_total_bytes_received_octo The total number of octo-bytes received by the JVB since it started.
# TYPE jitsi_total_bytes_received_octo counter
jitsi_total_bytes_received_octo {{.TotalBytesReceivedOcto}}
# HELP jitsi_total_bytes_sent The total number of bytes sent by the JVB  since it started.
# TYPE jitsi_total_bytes_sent counter
jitsi_total_bytes_sent {{.TotalBytesSent}}
# HELP jitsi_total_bytes_sent_octo The total number of octo-bytes sent by the JVB since it started.
# TYPE jitsi_total_bytes_sent_octo counter
jitsi_total_bytes_sent_octo {{.TotalBytesSentOcto}}
# HELP jitsi_total_colibri_web_socket_messages_received The total number messages received through COLIBRI web sockets.
# TYPE jitsi_total_colibri_web_socket_messages_received counter
jitsi_total_colibri_web_socket_messages_received {{.TotalColibriWebSocketMessagesReceived}}
# HELP jitsi_total_colibri_web_socket_messages_sent The total number messages sent through COLIBRI web sockets.
# TYPE jitsi_total_colibri_web_socket_messages_sent counter
jitsi_total_colibri_web_socket_messages_sent {{.TotalColibriWebSocketMessagesSent}}
# HELP jitsi_total_conference_seconds The sum of the lengths of all completed conferences, in seconds.
# TYPE jitsi_total_conference_seconds counter
jitsi_total_conference_seconds {{.TotalConferenceSeconds}}
# HELP jitsi_total_conferences_created The total number of conferences created on the bridge.
# TYPE jitsi_total_conferences_created counter
jitsi_total_conferences_created {{.TotalConferencesCreated}}
# HELP jitsi_total_conferences_completed The total number of conferences completed on the bridge.
# TYPE jitsi_total_conferences_completed counter
jitsi_total_conferences_completed {{.TotalConferencesCompleted}}
# HELP jitsi_total_data_channel_messages_received The total number messages received through data channels.
# TYPE jitsi_total_data_channel_messages_received counter
jitsi_total_data_channel_messages_received {{.TotalDataChannelMessagesReceived}}
# HELP jitsi_total_data_channel_messages_sent The total number messages sent through data channels.
# TYPE jitsi_total_data_channel_messages_sent counter
jitsi_total_data_channel_messages_sent {{.TotalDataChannelMessagesSent}}
# HELP jitsi_total_dominant_speaker_changes The total number of speaker changes.
# TYPE jitsi_total_dominant_speaker_changes counter
jitsi_total_dominant_speaker_changes {{.TotalDominantSpeakerChanges}}
# HELP jitsi_total_failed_conferences The total number of failed conferences on the bridge. A conference is marked as failed when all of its channels have failed. A channel is marked as failed if it had no payload activity.
# TYPE jitsi_total_failed_conferences counter
jitsi_total_failed_conferences {{.TotalFailedConferences}}
# HELP jitsi_total_ice_failed The total number of endpoints which failed to establish an ICE connection.
# TYPE jitsi_total_ice_failed counter
jitsi_total_ice_failed {{.TotalIceFailed}}
# HELP jitsi_total_ice_succeeded The number of endpoints which successfully established an ICE connection.
# TYPE jitsi_total_ice_succeeded counter
jitsi_total_ice_succeeded {{.TotalIceSucceeded}}
# HELP jitsi_total_ice_succeeded_relayed The number of endpoints which successfully established an relayed connection.
# TYPE jitsi_total_ice_succeeded_relayed counter
jitsi_total_ice_succeeded_relayed {{.TotalIceSucceededRelayed}}
# HELP jitsi_total_ice_succeeded_tcp The number of endpoints which successfully established an TCP-ICE connection.
# TYPE jitsi_total_ice_succeeded_tcp counter
jitsi_total_ice_succeeded_tcp {{.TotalIceSucceededTCP}}
# HELP jitsi_total_keyframes_received The total number of keyframes received.
# TYPE jitsi_total_keyframes_received counter
jitsi_total_keyframes_received {{.TotalKeyframesReceived}}
# HELP jitsi_total_layering_changes_received The total number of received changes in layering.
# TYPE jitsi_total_layering_changes_received counter
jitsi_total_layering_changes_received {{.TotalLayeringChangesReceived}}
# HELP jitsi_total_loss_controlled_participant_seconds The total number of participant-seconds that are loss-controlled.
# TYPE jitsi_total_loss_controlled_participant_seconds counter
jitsi_total_loss_controlled_participant_seconds {{.TotalLossControlledParticipantSeconds}}
# HELP jitsi_total_loss_degraded_participant_seconds The total number of participant-seconds that are loss-degraded.
# TYPE jitsi_total_loss_degraded_participant_seconds counter
jitsi_total_loss_degraded_participant_seconds {{.TotalLossDegradedParticipantSeconds}}
# HELP jitsi_total_loss_limited_participant_seconds The total number of participant-seconds that are loss-limited.
# TYPE jitsi_total_loss_limited_participant_seconds counter
jitsi_total_loss_limited_participant_seconds {{.TotalLossLimitedParticipantSeconds}}
# HELP jitsi_total_packets_dropped_octo The total of dropped packets handled by the OCTO JVB.
# TYPE jitsi_total_packets_dropped_octo counter
jitsi_total_packets_dropped_octo {{.TotalPacketsDroppedOcto}}
# HELP jitsi_total_packets_received The total of received packets handled by the JVB.
# TYPE jitsi_total_packets_received counter
jitsi_total_packets_received {{.TotalPacketsReceived}}
# HELP jitsi_total_packets_received_octo The total of incoming dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_received_octo counter
jitsi_total_packets_received_octo {{.TotalPacketsReceivedOcto}}
# HELP jitsi_total_packets_sent The total outgoing packets sent from the JVB.
# TYPE jitsi_total_packets_sent counter
jitsi_total_packets_sent {{.TotalPacketsSent}}
# HELP jitsi_total_packets_sent_octo The total outgoing packets sent from the Octo-JVB.
# TYPE jitsi_total_packets_sent_octo counter
jitsi_total_packets_sent_octo {{.TotalPacketsSentOcto}}
# HELP jitsi_total_partially_failed_conferences The total number of partially failed conferences on the bridge. A conference is marked as partially failed when some of its channels has failed. A channel is marked as failed if it had no payload activity.
# TYPE jitsi_total_partially_failed_conferences counter
jitsi_total_partially_failed_conferences {{.TotalPartiallyFailedConferences}}
# HELP jitsi_total_participants Total participants since running.
# TYPE jitsi_total_participants gauge
jitsi_total_participants {{.TotalParticipants}}
# HELP jitsi_total_video_stream_milliseconds_received The total number (in milliseconds) of videostreams received by the JVB.
# TYPE jitsi_total_video_stream_milliseconds_received counter
jitsi_total_video_stream_milliseconds_received {{.TotalVideoStreamMilliSecReceived}}
# HELP jitsi_videochannels An estimation of the number of current video streams forwarded by the bridge.
# TYPE jitsi_videochannels gauge
jitsi_videochannels {{.VideoChannels}}

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

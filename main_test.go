package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type constHandler struct {
	s string
}

func (h constHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(h.s))
	if err != nil {
		log.Fatal(err)
	}
}

func TestGetMetrics(t *testing.T) {
	tcs := []struct {
		statsJson string
		expected  string
	}{
		{ // no participant
			statsJson: `{"xmpp_service":{"total_recv":100,"total_sent":100},"jibri_detector":{"count":1,"available":1},"largest_conference":0,"conference_sizes":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"total_conferences_created":1,"threads":222,"xmpp":{},"jingle":{"received":{},"sent":{}},"bridge_failures":{"participants_moved":0,"bridges_removed":0},"avg_allocate_channels_req_time_nanos":0.0,"total_participants":1,"participant_notifications":{"ice_failed":0,"request_restart":0},"bridge_selector":{"total_least_loaded_in_region":0,"total_split_due_to_load":0,"total_not_loaded_in_region_in_conference":0,"in_shutdown_bridge_count":0,"total_least_loaded_in_region_in_conference":0,"total_not_loaded_in_region":0,"total_split_due_to_region":0,"bridge_count":1,"operational_bridge_count":1,"total_least_loaded_in_conference":0,"total_least_loaded":0},"jibri":{"total_sip_call_failures":0,"live_streaming_pending":0,"recording_pending":0,"live_streaming_active":0,"total_recording_failures":0,"sip_call_pending":0,"sip_call_active":0,"total_live_streaming_failures":0,"recording_active":0},"conferences":0,"participants":0,"slow_health_check":2}`,
			expected: `# HELP xmpp_service_total_recv stats about xmpp_service.
# TYPE xmpp_service_total_recv counter
jitsi_xmpp_service_total_recv 100
# HELP xmpp_service_total_sent stats about xmpp_service.
# TYPE xmpp_service_total_sent counter
jitsi_xmpp_service_total_sent 100
# HELP total jibris registered.
# TYPE jibri_count gauge
jitsi_jibri_count 1
# HELP total jibri available stats.
# TYPE jibri_available gauge
jitsi_jibri_available 1
# HELP jicofo largest_conferences stats.
# TYPE largest_conference_jicofo_stats gauge
jitsi_jicofo_largest_conferences 0
# HELP jitsi_conference_sizes Distribution of conference sizes on jicofo
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
# HELP jitsi_total_conferences_created The total number of conferences created on jicofo.
# TYPE jitsi_total_conferences_created counter
jitsi_total_conferences_created 1
# HELP jitsi_threads_on_jicofo Threads count on jicofo
# TYPE jitsi_jicofo_thread gauge
jitsi_jicofo_threads 222
# HELP jitsi_birdge_failure_participants_moved Participants moved after bridge failure
# TYPE jitsi_bridge_failure_participants_moved gauge
jitsi_bridge_failure_participants_moved 0
# HELP jitsi_birdge_removed after bridge failure
# TYPE jitsi_bridge_failure_bridge_removed gauge
jitsi_bridge_failure_bridge_removed 0
# HELP jitsi_avg_allocate_channels_req_time_nanos Avg allocate channel request time.
# TYPE jitsi_avg_allocate_channels_req_time_nanos gauge
jitsi_avg_allocate_channels_req_time_nanos 0
# HELP jitsi_total_participants Total number of participants joined so far.
# TYPE jitsi_total_participants counter
jitsi_total_participants 1
# HELP jitsi_participants_notification_ice_failed Stats about ice failure.
# TYPE jitsi_participants_notification_ice_failed counter
jitsi_participants_notification_ice_failed 0
# HELP jitsi_participants_notification_request_restart.
# TYPE jitsi_participants_notification_request_restart counter
jitsi_participants_notification_request_restart 0
# HELP jitsi_bridge_selector_total_least_loaded_in_region Bridges that are least loaded in region.
# TYPE jitsi_bridge_selector_total_least_loaded_in_region gauge
jitsi_bridge_selector_total_least_loaded_in_region 0
# HELP jitsi_bridge_selector_total_split_due_to_load Bridges splitted due to load.
# TYPE jitsi_bridge_selector_total_split_due_to_load gauge
jitsi_bridge_selector_total_split_due_to_load 0
# HELP jitsi_total_not_loaded_in_region_in_conference Bridges not loaded in a region in a conference.
# TYPE jitsi_bridge_selector_total_not_loaded_in_region_in_conference gauge
jitsi_bridge_selector_total_not_loaded_in_region_in_conference 0
# HELP jitsi_in_shutdown_bridge_count Bridges count that are in shutdown.
# TYPE jitsi_bridge_selector_in_shutdown_bridge_count gauge
jitsi_bridge_selector_in_shutdown_bridge_count 0
# HELP jitsi_total_least_loaded_in_region_in_conference Bridges that are lease loaded in a region in a conference.
# TYPE jitsi_bridge_selector_total_least_loaded_in_region_in_conference gauge
jitsi_bridge_selector_total_least_loaded_in_region_in_conference 0
# HELP jitsi_total_not_loaded_in_region Bridges that are not loaded in a region.
# TYPE jitsi_bridge_selector_total_not_loaded_in_region gauge
jitsi_bridge_selector_total_not_loaded_in_region 0
# HELP jitsi_total_split_due_to_region Bridges splitted due to region.
# TYPE jitsi_bridge_selector_total_not_loaded_in_region gauge
jitsi_bridge_selector_total_split_due_to_region 0
# HELP jitsi_bridge_count Number of bridges registered.
# TYPE jitsi_bridge_count gauge
jitsi_bridge_count 1
# HELP jitsi_operational_bridge_count Bridges that are operational.
# TYPE jitsi_operational_bridge_count gauge
jitsi_operational_bridge_count 1
# HELP jitsi_total_least_loaded_in_conference Bridges that are least loaded in a conference.
# TYPE jitsi_total_least_loaded_in_conference gauge
jitsi_total_least_loaded_in_conference 0
# HELP jitsi_total_least_loaded Bridges that are least loaded.
# TYPE jitsi_total_least_loaded gauge
jitsi_total_least_loaded 0
# HELP jitsi_total_sip_call_failures Total Sip call failures.
# TYPE jitsi_total_least_loaded counter
jitsi_total_sip_call_failures 0
# HELP jitsi_live_streaming_pending Live streaming pending state.
# TYPE jitsi_live_streaming_pending gauge
jitsi_live_streaming_pending 0
# HELP jitsi_recording_pending Jibri recording pending.
# TYPE jitsi_recording_pending gauge
jitsi_recording_pending 0
# HELP jitsi_live_streaming_active Jibri live streaming active.
# TYPE jitsi_live_streaming_active gauge
jitsi_live_streaming_active 0
# HELP jitsi_total_recording_failures Jibri total recording failures.
# TYPE jitsi_live_streaming_active counter
jitsi_total_recording_failures 0
# HELP jitsi_sip_call_pending Jibri sip call pending state.
# TYPE jitsi_sip_call_pending gauge
jitsi_sip_call_pending 0
# HELP jitsi_sip_call_active Jibri sip call active.
# TYPE jitsi_sip_call_active gauge
jitsi_sip_call_active 0
# HELP jitsi_total_live_streaming_failures Jibri live streaming failures.
# TYPE jitsi_total_live_streaming_failures gauge
jitsi_total_live_streaming_failures 0
# HELP jitsi_recording_active Total recording active.
# TYPE jitsi_recording_active gauge
jitsi_recording_active 0
# HELP jitsi_conferences Current running conferences.
# TYPE jitsi_conferences gauge
jitsi_conferences 0
# HELP jitsi_participants Current participants.
# TYPE jitsi_participants gauge
jitsi_participants 0`,
		},
		{ // 1 participant
			statsJson: `{"xmpp_service":{"total_recv":1000,"total_sent":1000},"jibri_detector":{"count":1,"available":1},"largest_conference":1,"conference_sizes":[0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"total_conferences_created":2,"threads":222,"xmpp":{},"jingle":{"received":{},"sent":{}},"bridge_failures":{"participants_moved":0,"bridges_removed":0},"avg_allocate_channels_req_time_nanos":0.0,"total_participants":2,"participant_notifications":{"ice_failed":0,"request_restart":0},"bridge_selector":{"total_least_loaded_in_region":0,"total_split_due_to_load":0,"total_not_loaded_in_region_in_conference":0,"in_shutdown_bridge_count":0,"total_least_loaded_in_region_in_conference":0,"total_not_loaded_in_region":0,"total_split_due_to_region":0,"bridge_count":1,"operational_bridge_count":1,"total_least_loaded_in_conference":0,"total_least_loaded":0},"jibri":{"total_sip_call_failures":0,"live_streaming_pending":0,"recording_pending":0,"live_streaming_active":0,"total_recording_failures":0,"sip_call_pending":0,"sip_call_active":0,"total_live_streaming_failures":0,"recording_active":0},"conferences":1,"participants":1,"slow_health_check":2}`,
			expected: `# HELP xmpp_service_total_recv stats about xmpp_service.
# TYPE xmpp_service_total_recv counter
jitsi_xmpp_service_total_recv 1000
# HELP xmpp_service_total_sent stats about xmpp_service.
# TYPE xmpp_service_total_sent counter
jitsi_xmpp_service_total_sent 1000
# HELP total jibris registered.
# TYPE jibri_count gauge
jitsi_jibri_count 1
# HELP total jibri available stats.
# TYPE jibri_available gauge
jitsi_jibri_available 1
# HELP jicofo largest_conferences stats.
# TYPE largest_conference_jicofo_stats gauge
jitsi_jicofo_largest_conferences 1
# HELP jitsi_conference_sizes Distribution of conference sizes on jicofo
# TYPE jitsi_conference_sizes gauge
jitsi_conference_sizes{conference_size="0"} 0
jitsi_conference_sizes{conference_size="1"} 1
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
# HELP jitsi_total_conferences_created The total number of conferences created on jicofo.
# TYPE jitsi_total_conferences_created counter
jitsi_total_conferences_created 2
# HELP jitsi_threads_on_jicofo Threads count on jicofo
# TYPE jitsi_jicofo_thread gauge
jitsi_jicofo_threads 222
# HELP jitsi_birdge_failure_participants_moved Participants moved after bridge failure
# TYPE jitsi_bridge_failure_participants_moved gauge
jitsi_bridge_failure_participants_moved 0
# HELP jitsi_birdge_removed after bridge failure
# TYPE jitsi_bridge_failure_bridge_removed gauge
jitsi_bridge_failure_bridge_removed 0
# HELP jitsi_avg_allocate_channels_req_time_nanos Avg allocate channel request time.
# TYPE jitsi_avg_allocate_channels_req_time_nanos gauge
jitsi_avg_allocate_channels_req_time_nanos 0
# HELP jitsi_total_participants Total number of participants joined so far.
# TYPE jitsi_total_participants counter
jitsi_total_participants 2
# HELP jitsi_participants_notification_ice_failed Stats about ice failure.
# TYPE jitsi_participants_notification_ice_failed counter
jitsi_participants_notification_ice_failed 0
# HELP jitsi_participants_notification_request_restart.
# TYPE jitsi_participants_notification_request_restart counter
jitsi_participants_notification_request_restart 0
# HELP jitsi_bridge_selector_total_least_loaded_in_region Bridges that are least loaded in region.
# TYPE jitsi_bridge_selector_total_least_loaded_in_region gauge
jitsi_bridge_selector_total_least_loaded_in_region 0
# HELP jitsi_bridge_selector_total_split_due_to_load Bridges splitted due to load.
# TYPE jitsi_bridge_selector_total_split_due_to_load gauge
jitsi_bridge_selector_total_split_due_to_load 0
# HELP jitsi_total_not_loaded_in_region_in_conference Bridges not loaded in a region in a conference.
# TYPE jitsi_bridge_selector_total_not_loaded_in_region_in_conference gauge
jitsi_bridge_selector_total_not_loaded_in_region_in_conference 0
# HELP jitsi_in_shutdown_bridge_count Bridges count that are in shutdown.
# TYPE jitsi_bridge_selector_in_shutdown_bridge_count gauge
jitsi_bridge_selector_in_shutdown_bridge_count 0
# HELP jitsi_total_least_loaded_in_region_in_conference Bridges that are lease loaded in a region in a conference.
# TYPE jitsi_bridge_selector_total_least_loaded_in_region_in_conference gauge
jitsi_bridge_selector_total_least_loaded_in_region_in_conference 0
# HELP jitsi_total_not_loaded_in_region Bridges that are not loaded in a region.
# TYPE jitsi_bridge_selector_total_not_loaded_in_region gauge
jitsi_bridge_selector_total_not_loaded_in_region 0
# HELP jitsi_total_split_due_to_region Bridges splitted due to region.
# TYPE jitsi_bridge_selector_total_not_loaded_in_region gauge
jitsi_bridge_selector_total_split_due_to_region 0
# HELP jitsi_bridge_count Number of bridges registered.
# TYPE jitsi_bridge_count gauge
jitsi_bridge_count 1
# HELP jitsi_operational_bridge_count Bridges that are operational.
# TYPE jitsi_operational_bridge_count gauge
jitsi_operational_bridge_count 1
# HELP jitsi_total_least_loaded_in_conference Bridges that are least loaded in a conference.
# TYPE jitsi_total_least_loaded_in_conference gauge
jitsi_total_least_loaded_in_conference 0
# HELP jitsi_total_least_loaded Bridges that are least loaded.
# TYPE jitsi_total_least_loaded gauge
jitsi_total_least_loaded 0
# HELP jitsi_total_sip_call_failures Total Sip call failures.
# TYPE jitsi_total_least_loaded counter
jitsi_total_sip_call_failures 0
# HELP jitsi_live_streaming_pending Live streaming pending state.
# TYPE jitsi_live_streaming_pending gauge
jitsi_live_streaming_pending 0
# HELP jitsi_recording_pending Jibri recording pending.
# TYPE jitsi_recording_pending gauge
jitsi_recording_pending 0
# HELP jitsi_live_streaming_active Jibri live streaming active.
# TYPE jitsi_live_streaming_active gauge
jitsi_live_streaming_active 0
# HELP jitsi_total_recording_failures Jibri total recording failures.
# TYPE jitsi_live_streaming_active counter
jitsi_total_recording_failures 0
# HELP jitsi_sip_call_pending Jibri sip call pending state.
# TYPE jitsi_sip_call_pending gauge
jitsi_sip_call_pending 0
# HELP jitsi_sip_call_active Jibri sip call active.
# TYPE jitsi_sip_call_active gauge
jitsi_sip_call_active 0
# HELP jitsi_total_live_streaming_failures Jibri live streaming failures.
# TYPE jitsi_total_live_streaming_failures gauge
jitsi_total_live_streaming_failures 0
# HELP jitsi_recording_active Total recording active.
# TYPE jitsi_recording_active gauge
jitsi_recording_active 0
# HELP jitsi_conferences Current running conferences.
# TYPE jitsi_conferences gauge
jitsi_conferences 1
# HELP jitsi_participants Current participants.
# TYPE jitsi_participants gauge
jitsi_participants 1`,
		},
		{ // 2 participants
			statsJson: `{"xmpp_service":{"total_recv":2000,"total_sent":2000},"jibri_detector":{"count":1,"available":1},"largest_conference":2,"conference_sizes":[0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"total_conferences_created":2,"threads":224,"xmpp":{},"jingle":{"received":{"transport-info":3,"session-accept":2},"sent":{"source-add":2,"session-initiate":2}},"bridge_failures":{"participants_moved":0,"bridges_removed":0},"avg_allocate_channels_req_time_nanos":9.60278E8,"total_participants":3,"participant_notifications":{"ice_failed":0,"request_restart":0},"bridge_selector":{"total_least_loaded_in_region":0,"total_split_due_to_load":0,"total_not_loaded_in_region_in_conference":0,"in_shutdown_bridge_count":0,"total_least_loaded_in_region_in_conference":0,"total_not_loaded_in_region":0,"total_split_due_to_region":0,"bridge_count":1,"operational_bridge_count":1,"total_least_loaded_in_conference":0,"total_least_loaded":1},"jibri":{"total_sip_call_failures":0,"live_streaming_pending":0,"recording_pending":0,"live_streaming_active":0,"total_recording_failures":0,"sip_call_pending":0,"sip_call_active":0,"total_live_streaming_failures":0,"recording_active":0},"conferences":1,"participants":2,"slow_health_check":2}`,
			expected: `# HELP xmpp_service_total_recv stats about xmpp_service.
# TYPE xmpp_service_total_recv counter
jitsi_xmpp_service_total_recv 2000
# HELP xmpp_service_total_sent stats about xmpp_service.
# TYPE xmpp_service_total_sent counter
jitsi_xmpp_service_total_sent 2000
# HELP total jibris registered.
# TYPE jibri_count gauge
jitsi_jibri_count 1
# HELP total jibri available stats.
# TYPE jibri_available gauge
jitsi_jibri_available 1
# HELP jicofo largest_conferences stats.
# TYPE largest_conference_jicofo_stats gauge
jitsi_jicofo_largest_conferences 2
# HELP jitsi_conference_sizes Distribution of conference sizes on jicofo
# TYPE jitsi_conference_sizes gauge
jitsi_conference_sizes{conference_size="0"} 0
jitsi_conference_sizes{conference_size="1"} 0
jitsi_conference_sizes{conference_size="2"} 1
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
# HELP jitsi_total_conferences_created The total number of conferences created on jicofo.
# TYPE jitsi_total_conferences_created counter
jitsi_total_conferences_created 2
# HELP jitsi_threads_on_jicofo Threads count on jicofo
# TYPE jitsi_jicofo_thread gauge
jitsi_jicofo_threads 224
# HELP jitsi_birdge_failure_participants_moved Participants moved after bridge failure
# TYPE jitsi_bridge_failure_participants_moved gauge
jitsi_bridge_failure_participants_moved 0
# HELP jitsi_birdge_removed after bridge failure
# TYPE jitsi_bridge_failure_bridge_removed gauge
jitsi_bridge_failure_bridge_removed 0
# HELP jitsi_avg_allocate_channels_req_time_nanos Avg allocate channel request time.
# TYPE jitsi_avg_allocate_channels_req_time_nanos gauge
jitsi_avg_allocate_channels_req_time_nanos 9.60278e&#43;08
# HELP jitsi_total_participants Total number of participants joined so far.
# TYPE jitsi_total_participants counter
jitsi_total_participants 3
# HELP jitsi_participants_notification_ice_failed Stats about ice failure.
# TYPE jitsi_participants_notification_ice_failed counter
jitsi_participants_notification_ice_failed 0
# HELP jitsi_participants_notification_request_restart.
# TYPE jitsi_participants_notification_request_restart counter
jitsi_participants_notification_request_restart 0
# HELP jitsi_bridge_selector_total_least_loaded_in_region Bridges that are least loaded in region.
# TYPE jitsi_bridge_selector_total_least_loaded_in_region gauge
jitsi_bridge_selector_total_least_loaded_in_region 0
# HELP jitsi_bridge_selector_total_split_due_to_load Bridges splitted due to load.
# TYPE jitsi_bridge_selector_total_split_due_to_load gauge
jitsi_bridge_selector_total_split_due_to_load 0
# HELP jitsi_total_not_loaded_in_region_in_conference Bridges not loaded in a region in a conference.
# TYPE jitsi_bridge_selector_total_not_loaded_in_region_in_conference gauge
jitsi_bridge_selector_total_not_loaded_in_region_in_conference 0
# HELP jitsi_in_shutdown_bridge_count Bridges count that are in shutdown.
# TYPE jitsi_bridge_selector_in_shutdown_bridge_count gauge
jitsi_bridge_selector_in_shutdown_bridge_count 0
# HELP jitsi_total_least_loaded_in_region_in_conference Bridges that are lease loaded in a region in a conference.
# TYPE jitsi_bridge_selector_total_least_loaded_in_region_in_conference gauge
jitsi_bridge_selector_total_least_loaded_in_region_in_conference 0
# HELP jitsi_total_not_loaded_in_region Bridges that are not loaded in a region.
# TYPE jitsi_bridge_selector_total_not_loaded_in_region gauge
jitsi_bridge_selector_total_not_loaded_in_region 0
# HELP jitsi_total_split_due_to_region Bridges splitted due to region.
# TYPE jitsi_bridge_selector_total_not_loaded_in_region gauge
jitsi_bridge_selector_total_split_due_to_region 0
# HELP jitsi_bridge_count Number of bridges registered.
# TYPE jitsi_bridge_count gauge
jitsi_bridge_count 1
# HELP jitsi_operational_bridge_count Bridges that are operational.
# TYPE jitsi_operational_bridge_count gauge
jitsi_operational_bridge_count 1
# HELP jitsi_total_least_loaded_in_conference Bridges that are least loaded in a conference.
# TYPE jitsi_total_least_loaded_in_conference gauge
jitsi_total_least_loaded_in_conference 0
# HELP jitsi_total_least_loaded Bridges that are least loaded.
# TYPE jitsi_total_least_loaded gauge
jitsi_total_least_loaded 1
# HELP jitsi_total_sip_call_failures Total Sip call failures.
# TYPE jitsi_total_least_loaded counter
jitsi_total_sip_call_failures 0
# HELP jitsi_live_streaming_pending Live streaming pending state.
# TYPE jitsi_live_streaming_pending gauge
jitsi_live_streaming_pending 0
# HELP jitsi_recording_pending Jibri recording pending.
# TYPE jitsi_recording_pending gauge
jitsi_recording_pending 0
# HELP jitsi_live_streaming_active Jibri live streaming active.
# TYPE jitsi_live_streaming_active gauge
jitsi_live_streaming_active 0
# HELP jitsi_total_recording_failures Jibri total recording failures.
# TYPE jitsi_live_streaming_active counter
jitsi_total_recording_failures 0
# HELP jitsi_sip_call_pending Jibri sip call pending state.
# TYPE jitsi_sip_call_pending gauge
jitsi_sip_call_pending 0
# HELP jitsi_sip_call_active Jibri sip call active.
# TYPE jitsi_sip_call_active gauge
jitsi_sip_call_active 0
# HELP jitsi_total_live_streaming_failures Jibri live streaming failures.
# TYPE jitsi_total_live_streaming_failures gauge
jitsi_total_live_streaming_failures 0
# HELP jitsi_recording_active Total recording active.
# TYPE jitsi_recording_active gauge
jitsi_recording_active 0
# HELP jitsi_conferences Current running conferences.
# TYPE jitsi_conferences gauge
jitsi_conferences 1
# HELP jitsi_participants Current participants.
# TYPE jitsi_participants gauge
jitsi_participants 2`,
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
			t.Logf("\n\nValue of body String is: %s ", rr.Body.String())
			t.Logf("\n\nValue of expected String is: %s ", tc.expected)
			t.Log("\n\n")
			t.Errorf("Response does not match the expected string:\n%s", cmp.Diff(rr.Body.String(), tc.expected))
		}

		srv.Close()
	}
}

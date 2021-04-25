package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
	webflag "github.com/prometheus/exporter-toolkit/web/kingpinflag"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	tlsConfig       = webflag.AddFlags(kingpin.CommandLine)
	listenAddress   = kingpin.Flag("web.listen-address", "Address on which to expose metrics").Default(":9996").String()
	metricsPath     = kingpin.Flag("web.metrics-path", "Path under which to expose metrics.").Default("/metrics").String()
	jicofoScrapeURI = kingpin.Flag("jicofo.scrape-uri", "Jitsi jicofo URL to scrape").Default("http://localhost:8888/stats").String()
	tpl             = template.Must(template.New("stats").Parse(`# HELP xmpp_service stats about xmpp_service.
# TYPE xmpp_service gauge
jitsi_xmpp_service {{.XmppService}}
# HELP total jibris registered.
# TYPE jibri_count gauge
jitsi_jibri_count {{.JibriDetector.Count}}
# HELP total jibri available stats.
# TYPE jibri_available gauge
jitsi_jibri_available {{.JibriDetector.Available}}
# HELP jicofo largest_conferences stats.
# TYPE largest_conference_jicofo_stats  gauge
jitsi_jicofo_largest_conferences {{.LargestConferences}}
# HELP jitsi_conference_sizes Distribution of conference sizes on jicofo
# TYPE jitsi_conference_sizes gauge
{{ range $key, $value := .ConferenceSizes -}}
jitsi_conference_sizes{conference_size="{{$key}}"} {{ $value }}
{{ end -}}
# HELP jitsi_total_conferences_created The total number of conferences created on jicofo.
# TYPE jitsi_total_conferences_created counter
jitsi_total_conferences_created {{.TotalConferencesCreated}}
# HELP jitsi_threads_on_jicofo Threads count on jicofo
# TYPE jitsi_jicofo_thread gauge
jitsi_jicofo_threads {{.Threads}}
# HELP jitsi_birdge_failure_participants_moved Participants moved after bridge failure
# TYPE jitsi_bridge_failure_participants_moved gauge
jitsi_bridge_failure_participants_moved {{.BridgeFailure.ParticipantsMoved}}
# HELP jitsi_birdge_removed after bridge failure
# TYPE jitsi_bridge_failure_bridge_removed gauge
jitsi_bridge_failure_bridge_removed {{.BridgeFailure.BridgesRemoved}}
# HELP jitsi_avg_allocate_channels_req_time_nanos Avg allocate channel request time.
# TYPE jitsi_avg_allocate_channels_req_time_nanos gauge
jitsi_avg_allocate_channels_req_time_nanos {{.AvgAllocateChannelsReqTimeNanos}}
# HELP jitsi_total_participants Total number of participants joined so far.
# TYPE jitsi_total_participants counter
jitsi_total_participants {{.TotalParticipants}}
# HELP jitsi_participants_notification_ice_failed Stats about ice failure.
# TYPE jitsi_participants_notification_ice_failed counter
jitsi_participants_notification_ice_failed {{.ParticipantNotifications.IceFailed}}
# HELP jitsi_participants_notification_request_restart.
# TYPE jitsi_participants_notification_request_restart counter
jitsi_participants_notification_request_restart {{.ParticipantNotifications.RequestRestart}}
# HELP jitsi_bridge_selector_total_least_loaded_in_region Bridges that are least loaded in region.
# TYPE jitsi_bridge_selector_total_least_loaded_in_region gauge
jitsi_bridge_selector_total_least_loaded_in_region {{.BridgeSelector.TotalLeastLoadedInRegion}}
# HELP jitsi_bridge_selector_total_split_due_to_load Bridges splitted due to load.
# TYPE jitsi_bridge_selector_total_split_due_to_load gauge
jitsi_bridge_selector_total_split_due_to_load {{.BridgeSelector.TotalSplitDueToLoad}}
# HELP jitsi_total_not_loaded_in_region_in_conference Bridges not loaded in a region in a conference.
# TYPE jitsi_bridge_selector_total_not_loaded_in_region_in_conference gauge
jitsi_bridge_selector_total_not_loaded_in_region_in_conference {{.BridgeSelector.TotalNotLoadedInRegionInConf}}
# HELP jitsi_in_shutdown_bridge_count Bridges count that are in shutdown.
# TYPE jitsi_bridge_selector_in_shutdown_bridge_count gauge
jitsi_bridge_selector_in_shutdown_bridge_count {{.BridgeSelector.InShutdownCount}}
# HELP jitsi_total_least_loaded_in_region_in_conference Bridges that are lease loaded in a region in a conference.
# TYPE jitsi_bridge_selector_total_least_loaded_in_region_in_conference gauge
jitsi_bridge_selector_total_least_loaded_in_region_in_conference {{.BridgeSelector.TotalLeastLoadedInRegionInConf}}
# HELP jitsi_total_not_loaded_in_region Bridges that are not loaded in a region.
# TYPE jitsi_bridge_selector_total_not_loaded_in_region gauge
jitsi_bridge_selector_total_not_loaded_in_region {{.BridgeSelector.TotalNotLoadedInRegion}}
# HELP jitsi_total_split_due_to_region Bridges splitted due to region.
# TYPE jitsi_bridge_selector_total_not_loaded_in_region gauge
jitsi_bridge_selector_total_split_due_to_region {{.BridgeSelector.TotalSplitDueToregion}}
# HELP jitsi_bridge_count Number of bridges registered.
# TYPE jitsi_bridge_count gauge
jitsi_bridge_count {{.BridgeSelector.BridgeCount}}
# HELP jitsi_operational_bridge_count Bridges that are operational.
# TYPE jitsi_operational_bridge_count gauge
jitsi_operational_bridge_count {{.BridgeSelector.OperationalBridgeCount}}
# HELP jitsi_total_least_loaded_in_conference Bridges that are least loaded in a conference.
# TYPE jitsi_total_least_loaded_in_conference gauge
jitsi_total_least_loaded_in_conference {{.BridgeSelector.TotalLeastLoadedInConference}}
# HELP jitsi_total_least_loaded Bridges that are least loaded.
# TYPE jitsi_total_least_loaded gauge
jitsi_total_least_loaded {{.BridgeSelector.TotalLeastLoaded}}
# HELP jitsi_total_sip_call_failures Total Sip call failures.
# TYPE jitsi_total_least_loaded counter
jitsi_total_sip_call_failures {{.Jibri.TotalSipcallFailure}}
# HELP jitsi_live_streaming_pending Live streaming pending state.
# TYPE jitsi_live_streaming_pending gauge
jitsi_live_streaming_pending {{.Jibri.LiveStreamingPending}}
# HELP jitsi_recording_pending Jibri recording pending.
# TYPE jitsi_recording_pending gauge
jitsi_recording_pending {{.Jibri.RecordingPending}}
# HELP jitsi_live_streaming_active Jibri live streaming active.
# TYPE jitsi_live_streaming_active gauge
jitsi_live_streaming_active {{.Jibri.LiveStreamingActive}}
# HELP jitsi_total_recording_failures Jibri total recording failures.
# TYPE jitsi_live_streaming_active counter
jitsi_total_recording_failures {{.Jibri.TotalRecordingFailures}}
# HELP jitsi_sip_call_pending Jibri sip call pending state.
# TYPE jitsi_sip_call_pending gauge
jitsi_sip_call_pending {{.Jibri.SipcallPending}}
# HELP jitsi_sip_call_active Jibri sip call active.
# TYPE jitsi_sip_call_active gauge
jitsi_sip_call_active {{.Jibri.SipcallActive}}
# HELP jitsi_total_live_streaming_failures Jibri live streaming failures.
# TYPE jitsi_total_live_streaming_failures gauge
jitsi_total_live_streaming_failures {{.Jibri.TotalLiveStreamingFailures}}
# HELP jitsi_recording_active Total recording active.
# TYPE jitsi_recording_active gauge
jitsi_recording_active {{.Jibri.RecordingActive}}
# HELP jitsi_conferences Current running conferences.
# TYPE jitsi_conferences gauge
jitsi_conferences {{.Conferences}}
# HELP jitsi_participants Current participants.
# TYPE jitsi_participants gauge
jitsi_participants {{.Participants}}`))
)

// jicofoStats
type jicofoStats struct {
	XmppService   struct{} `json:"xmpp_service"`
	JibriDetector struct {
		Count     int `json:"count"`
		Available int `json:"available"`
	} `json:"jibri_detector"`

	LargestConferences      int   `json:"largest_conference"`
	ConferenceSizes         []int `json:"conference_sizes"`
	TotalConferencesCreated int   `json:"total_conferences_created"`
	Threads                 int   `json:"threads"`
	Jingle                  struct {
		Received struct{} `json:"received"`
		Sent     struct{} `json:"sent"`
	} `json:"jingle"`
	BridgeFailure struct {
		ParticipantsMoved int `json:"participants_moved"`
		BridgesRemoved    int `json:"bridges_removed"`
	} `json:"bridge_failures"`
	AvgAllocateChannelsReqTimeNanos float32 `json:"avg_allocate_channels_req_time_nanos"`
	TotalParticipants               int     `json:"total_participants"`
	ParticipantNotifications        struct {
		IceFailed      int `json:"ice_failed"`
		RequestRestart int `json:"request_restart"`
	} `json:"participant_notifications"`
	BridgeSelector struct {
		TotalLeastLoadedInRegion       int `json:"total_least_loaded_in_region"`
		TotalSplitDueToLoad            int `json:"total_split_due_to_load"`
		TotalNotLoadedInRegionInConf   int `json:"total_not_loaded_in_region_in_conference"`
		InShutdownCount                int `json:"in_shutdown_bridge_count"`
		TotalLeastLoadedInRegionInConf int `json:"total_least_loaded_in_region_in_conference"`
		TotalNotLoadedInRegion         int `json:"total_not_loaded_in_region"`
		TotalSplitDueToregion          int `json:"total_split_due_to_region"`
		BridgeCount                    int `json:"bridge_count"`
		OperationalBridgeCount         int `json:"operational_bridge_count"`
		TotalLeastLoadedInConference   int `json:"total_least_loaded_in_conference"`
		TotalLeastLoaded               int `json:"total_least_loaded"`
	} `json:"bridge_selector"`
	Jibri struct {
		TotalSipcallFailure        int `json:"total_sip_call_failures"`
		LiveStreamingPending       int `json:"live_streaming_pending"`
		RecordingPending           int `json:"recording_pending"`
		LiveStreamingActive        int `json:"live_streaming_active"`
		TotalRecordingFailures     int `json:"total_recording_failures"`
		SipcallPending             int `json:"sip_call_pending"`
		SipcallActive              int `json:"sip_call_active"`
		TotalLiveStreamingFailures int `json:"total_live_streaming_failures"`
		RecordingActive            int `json:"recording_active"`
	} `json:"jibri"`
	Conferences     int `json:"conferences"`
	Participants    int `json:"participants"`
	SlowHealthCheck int `json:"slow_health_check"`
}

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

	var stats jicofoStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		log.Printf("json decoding error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	_ = tpl.Execute(w, &stats)
}

func main() {

	promlogConfig := &promlog.Config{}
	flag.AddFlags(kingpin.CommandLine, promlogConfig)
	kingpin.Version(version.Print("Jicofo Exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	logger := promlog.New(promlogConfig)

	level.Info(logger).Log("msg", "Starting jicofo_exporter", "version", version.Info())
	level.Info(logger).Log("msg", "Build context", "context", version.BuildContext())
	level.Info(logger).Log("msg", "Listening on address", "address", *listenAddress)
	// prometheus metrics endpoint
	http.Handle(*metricsPath, handler{sourceURL: *jicofoScrapeURI})
	// default page screen when someone opens "http:localhost:9996/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Jicofo Exporter</title></head>
			<body>
			<h1>Jicofo Exporter</h1>
			<p><a href='` + *metricsPath + `'>Metrics</a></p>
			</body>
			</html>`))
	})

	// Listening on the server port
	srv := &http.Server{Addr: *listenAddress}
	if err := web.ListenAndServe(srv, *tlsConfig, logger); err != nil {
		level.Error(logger).Log("msg", "Error starting HTTP server", "err", err)
		os.Exit(1)
	}
}

# Jicofo Metrics Exporter

[![Release](https://img.shields.io/github/v/release/prayagsingh/prometheus-jicofo-exporter?color=dark-green)](https://github.com/prayagsingh/prometheus-jicofo-exporter/releases/)
[![Integration](https://github.com/prayagsingh/prometheus-jicofo-exporter/workflows/Integration/badge.svg?branch=main)](https://github.com/prayagsingh/prometheus-jicofo-exporter/workflows/Integration/badge.svg?branch=main) [![Quality](https://github.com/prayagsingh/prometheus-jicofo-exporter/workflows/Quality/badge.svg?branch=main)](https://github.com/prayagsingh/prometheus-jicofo-exporter/workflows/Quality/badge.svg?branch=main) [![Docker Image Version](https://img.shields.io/docker/v/prayagsingh/prometheus-jicofo-exporter/latest)](https://hub.docker.com/r/prayagsingh/prometheus-jicofo-exporter) [![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/prayagsingh/prometheus-jicofo-exporter)](https://hub.docker.com/r/prayagsingh/prometheus-jicofo-exporter)

Prometheus Exporter for Jibri written in Go. Special thanks to `Systemli` team since this project is based on [Prometheus-jitsi-meet-exporter](https://github.com/systemli/prometheus-jitsi-meet-exporter) repository.

There's only one GET endpoint to check the stats of jicofo (like /stats); you can configure the used URL with the `jicofo-scrape-uri`.
The exporter will handle it.

## Usage

```
go get github.com/prayagsingh/prometheus-jicofo-exporter
go install github.com/prayagsingh/prometheus-jicofo-exporter
$GOPATH/bin/prometheus-jicofo-exporter
```

### Docker

```
docker run -p 9996:9996 prayagsingh/prometheus-jicofo-exporter:latest -jicofo.scrape-uri http://localhost:8888/stats
```

## Metrics

```
# HELP xmpp_service stats about xmpp_service.
# TYPE xmpp_service gauge
jitsi_xmpp_service {}
# HELP total jibris registered.
# TYPE jibri_count gauge
jitsi_jibri_count 1
# HELP total jibri available stats.
# TYPE jibri_available gauge
jitsi_jibri_available 1
# HELP jicofo largest_conferences stats.
# TYPE largest_conference_jicofo_stats  gauge
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
jitsi_total_conferences_created 3
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
jitsi_total_participants 4
# HELP jitsi_participants_notification_ice_failed Stats about ice failure.
# TYPE jitsi_participants_notification_ice_failed counter
jitsi_participants_notification_ice_failed 0
# HELP jitsi_participants_notification_request_restart.
# TYPE jitsi_participants_notification_request_restart counter
jitsi_participants_notification_request_restart 3
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
jitsi_total_least_loaded 2
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
jitsi_participants 0
```

## TODO

Add xmpp_service stats and jingle stats

## License

GPLv3

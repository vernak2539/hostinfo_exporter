# HostInfo Exporter

Meant to be used with [Prometheus](https://prometheus.io/).

I know this isn't the right way to go about this, but wanted to experiment with Go, Github Actions, and building a 
Prometheus exporter.

## Motivation

I wanted to export host information into Prometheus. Things like hostname, ip (internal + external), and various
other metrics

## Gotchas

Since none of these values satisfy the constraint of being a numeric value, I went along the line of using a 
Unix timestamp of the last time this exporter was started. All other metrics will be delivered as labels.

:rotating_light: This means that if any of these values were to change, the exporter would continue to report the
old values until it was restarted.



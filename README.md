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

## Usage

If you really want to use this (don't use in any production grade system), then you can follow the information below.


Check the latest release on the [release](https://github.com/vernak2539/hostinfo_exporter/releases) page to get the link
for the architecture you're wanting (limited selection now).

```
cd /opt/hostinfo_exporter
sudo wget https://github.com/vernak2539/hostinfo_exporter/releases/download/v0.2.0/hostinfo_exporter-v0.2.0-linux-arm64
sudo mv hostinfo_exporter-v0.2.0-linux-arm64 hostinfo_exporter
sudo chmod +x ./hostinfo_exporter
```

I would suggest setting up a systemd service and running it under it's own user

**Create User**
```sh
sudo useradd -M hostinfo_exporter
sudo chown -R hostinfo_exporter:root /opt/hostinfo_exporter
```

**Setup systemd**
```sh
sudo vi /etc/systemd/system/hostinfo_exporter.service
```

Insert following information into service file:

```
[Unit]
Description=HostInfo Exporter
After=network.target

[Service]
User=hostinfo_exporter
Type=simple
ExecStart=/opt/hostinfo_exporter/hostinfo_exporter

[Install]
WantedBy=multi-user.target
```

Run following commands:
```sh
sudo systemctl daemon-reload
sudo systemctl start hostinfo_exporter
```

Maybe add a cronjob to restart it every so often so the labels get updated :shrug:

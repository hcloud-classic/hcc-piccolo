[![pipeline status](http://210.207.104.150:8100/iitp-sds/piccolo/badges/master/pipeline.svg)](http://210.207.104.150:8100/iitp-sds/piccolo/pipelines)
[![coverage report](http://210.207.104.150:8100/iitp-sds/piccolo/badges/master/coverage.svg)](http://210.207.104.150:8100/iitp-sds/piccolo/commits/master)
[![go report](http://210.207.104.150:8100/iitp-sds/hcloud-badge/raw/feature/dev/hcloud-badge_piccolo.svg)](http://210.207.104.150:8100/iitp-sds/hcloud-badge/raw/feature/dev/goreport_piccolo)



API Gateway 역할



### Service 등록

`/etc/systemd/system/piccolo.service`
```shell
[Unit]
Description=HCC Piccolo Service

[Service]
Type=simple
User=root
ExecStart=/usr/local/bin/piccolo
ExecStop=/usr/bin/killall piccolo
Restart=on-failure
RestartPreventExitStatus=100

[Install]
WantedBy=multi-user.target
# Alias 를 설정 하면 해당 경로에 하단 이름으로 symlink를 만든다.
Alias=Piccolo.service
```

```shell
$ systemctl enable piccolo.service
Synchronizing state of piccolo.service with SysV service script with /lib/systemd/systemd-sysv-install.
Executing: /lib/systemd/systemd-sysv-install enable piccolo
Created symlink /etc/systemd/system/Piccolo.service → /etc/systemd/system/piccolo.service.
$ systemctl start piccolo
```
all: clean build

clean:
	rm -f smartctl_telegraf
	rm -f nvme_telegraf

build:
	go build -ldflags="-X 'main.Command=/usr/sbin/smartctl'" -o smartctl_telegraf commandwrapper.go
	chown root:telegraf smartctl_telegraf
	chmod 4550 smartctl_telegraf
	go build -ldflags="-X 'main.Command=/usr/sbin/nvme'" -o nvme_telegraf commandwrapper.go
	chown root:telegraf nvme_telegraf
	chmod 4550 nvme_telegraf

all: clean build

clean:
	rm -f smartctl_telegraf
	rm -f nvme_telegraf

build:
	go build -ldflags="-X 'main.Command=/usr/sbin/smartctl'" -o smartctl_telegraf commandwrapper.go
	chown telegraf smartctl_telegraf
	chmod 4500 smartctl_telegraf
	go build -ldflags="-X 'main.Command=/usr/sbin/nvme'" -o nvme_telegraf commandwrapper.go
	chown telegraf nvme_telegraf
	chmod 4500 nvme_telegraf

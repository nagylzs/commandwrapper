
# commandwrapper

This is a utility that can wrap another program. You can then set the suid flag on the wrapper.

This was originally developed for the S.M.A.R.T. input module for Telegraf, see here:

    https://github.com/influxdata/telegraf/issues/8690

## How to compile

    git clone git@github.com:nagylzs/commandwrapper.git
    cd commandwrapper
    go build -o your_wrapper commandwrapper.go -ldflags="-X 'main.Command=/usr/sbin/your_wrapped_command'"

## Real life example

The telegraf S.M.A.R.T. input plugin needs to run `smartctl` and `nvme` commands as root. Instead of setting up
`sudo` (which is problematic, see https://github.com/influxdata/telegraf/issues/8690) you can compile two wrappers
that are owner by uid=root, gid=telegraf, and set the suid flag on the wrappers. This allows you to use them
directly without sudo:

```
[[inputs.smart]]
    path_smartctl = "/usr/sbin/smartctl_telegraf"
    path_nvme = "/usr/sbin/nvme_telegraf"
    use_sudo = false
```

For this complete example, please see `Makefile` or just use `make` to compile them.


# commandwrapper

This is a utility that can wrap another program. You can then set the suid flag on the wrapper.

This was originally developed for the S.M.A.R.T. input module for Telegraf, see here:

    https://github.com/influxdata/telegraf/issues/8690

## How to compile

    git clone git@github.com:nagylzs/commandwrapper.git
    cd commandwrapper
    go build -o your_wrapper commandwrapper.go -ldflags="-X 'main.Command=/usr/sbin/your_wrapped_command'"

For more real-life examples, please see the `Makefile`

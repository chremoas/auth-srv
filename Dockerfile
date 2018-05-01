FROM scratch
MAINTAINER Brian Hechinger <wonko@4amlunch.net>

ADD auth-srv-linux-amd64 auth-srv
VOLUME /etc/chremoas

ENTRYPOINT ["/auth-srv", "--configuration_file", "/etc/chremoas/auth-bot.yaml"]
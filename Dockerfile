FROM arm32v6/alpine
LABEL maintainer="maurer.it@gmail.com"

RUN apk update && apk upgrade && apk add ca-certificates

ADD ./auth-srv /
WORKDIR /

RUN mkdir /etc/chremoas
VOLUME /etc/chremoas

RUN rm -rf /var/cache/apk/*

ENV MICRO_REGISTRY_ADDRESS chremoas-consul:8500

CMD [""]
ENTRYPOINT ["./auth-srv", "--configuration_file", "/etc/chremoas/auth-srv.yaml"]

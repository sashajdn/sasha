### --- Service --- ###
FROM alpine:latest
MAINTAINER alexperkins.dev@gmail.com
ADD sasha.service.discord /
ADD ./config /service.discord/config
EXPOSE 8080
RUN apk --no-cache add ca-certificates
ENTRYPOINT ["/sasha.service.discord"]

### --- Service --- ###
FROM alpine:latest
MAINTAINER alexperkins.dev@gmail.com
ADD sasha.service.locationtracker /
ADD ./config /service.locationtracker/config
EXPOSE 8080
RUN apk --no-cache add ca-certificates
ENTRYPOINT ["/sasha.service.locationtracker"]

### --- Service --- ###
FROM alpine:latest
MAINTAINER alexperkins.dev@gmail.com
ADD sasha.service.github /
ADD ./config /service.github/config
EXPOSE 8080
RUN apk --no-cache add ca-certificates
ENTRYPOINT ["/sasha.service.github"]

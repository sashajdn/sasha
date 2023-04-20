### --- Service --- ###
FROM alpine:latest
MAINTAINER alexperkins.dev@gmail.com
ADD sasha.service.openai /
ADD ./config /service.openai/config
EXPOSE 8080
RUN apk --no-cache add ca-certificates
ENTRYPOINT ["/sasha.service.openai"]

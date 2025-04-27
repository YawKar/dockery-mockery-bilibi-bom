FROM alpine:3.21.3 AS base

VOLUME [ "/data" ] 

COPY script.sh script.sh

ENTRYPOINT [ "sh", "./script.sh" ]

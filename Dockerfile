FROM ubuntu:latest
LABEL authors="LGB"

ENTRYPOINT ["top", "-b"]
FROM golang as buildStage
WORKDIR /build
COPY ./main /build
RUN cd /build/main & go build -o binlog-printer

FROM debian:11-slim
WORKDIR /app
COPY --from=buildStage /build/binlog-printer /app
CMD ["sh", "./binlog-printer"]
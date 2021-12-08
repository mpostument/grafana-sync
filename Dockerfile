FROM alpine

RUN apk update \
  && apk add ca-certificates \
  && update-ca-certificates

RUN apk --no-cache add openssl wget 

RUN wget https://github.com/mpostument/grafana-sync/releases/download/1.4.3/grafana-sync_1.4.3_Linux_x86_64.tar.gz \
  && tar -xvzf grafana-sync_1.4.3_Linux_x86_64.tar.gz \
  && rm grafana-sync_1.4.3_Linux_x86_64.tar.gz \
  && chmod +x ./grafana-sync

ENTRYPOINT ["./grafana-sync"]

ARG VERSION=latest

FROM  keppel.eu-de-1.cloud.sap/ccloud/grafana-operator-binaries:$VERSION as grafana-operator-binaries

FROM alpine:3.8 as grafana-operator
LABEL maintainer "Stefan Hipfel <stefan.hipfel@sap.com>"
LABEL source_repository "https://github.com/sapcc/grafana-operator"
RUN apk add --no-cache curl
RUN curl -Lo /bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.1/dumb-init_1.2.1_amd64 \
	&& chmod +x /bin/dumb-init \
	&& dumb-init -V
COPY --from=grafana-operator-binaries /apiserver /manager /usr/local/bin/
ENTRYPOINT ["dumb-init", "--"]
CMD ["apiserver"]

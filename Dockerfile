FROM registry.ci.openshift.org/openshift/release:golang-1.16 AS builder
ARG VERSION
WORKDIR /go/src/github.com/openshift/cluster-api-provider-powervs
COPY . .

RUN GOPROXY=off NO_DOCKER=1 GOARCH=ppc64le GOOS=linux VERSION=$VERSION make build

FROM --platform=ppc64le registry.access.redhat.com/ubi8/ubi:8.4
COPY --from=builder /go/src/github.com/openshift/cluster-api-provider-powervs/bin/machine-controller-manager /

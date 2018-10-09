FROM alpine
LABEL maintainer="Richard Case <richard.case@outlook.com>"

RUN apk --no-cache add openssl musl-dev ca-certificates
COPY serverless-operator /usr/local/bin

ENTRYPOINT ["/usr/local/bin/serverless-operator"]

# Builder

ARG GITHUB_PATH=github.com/Julia1505/RedditCloneBack

FROM golang:1.19.1-alpine AS builder
RUN apk add --update make git gcc
COPY . /home/${GITHUB_PATH}
WORKDIR /home/${GITHUB_PATH}
RUN make build

FROM alpine:latest as server
WORKDIR /root/
COPY --from=builder /home/${GITHUB_PATH}/bin/redditclone .
COPY --from=builder /home/${GITHUB_PATH}/static ./static
RUN chown root:root redditclone
EXPOSE 8080

CMD ["./redditclone"]

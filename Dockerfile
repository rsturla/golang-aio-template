ARG GO_VERSION=1.21
ARG NODE_VERSION=20

FROM node:${NODE_VERSION}-alpine AS node-builder

WORKDIR /app
COPY ui/package.json ui/yarn.lock ./
RUN yarn install --immutable
COPY ui/ .
ENV NEXT_TELEMETRY_DISABLED=1
RUN yarn run export


FROM golang:${GO_VERSION}-alpine AS go-builder
WORKDIR /app
COPY go.mod main.go ./
COPY --from=node-builder /app/dist ./ui/dist
RUN go build -o golang-aio .


FROM scratch
WORKDIR /app
COPY --from=go-builder /app/golang-aio .
ENTRYPOINT ["./golang-aio"]
EXPOSE 8080

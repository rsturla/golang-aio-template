ARG GO_VERSION=1.21
ARG NODE_VERSION=20


FROM node:${NODE_VERSION}-alpine AS node-builder
WORKDIR /app
COPY web/package.json web/yarn.lock web/.yarnrc.yml ./
COPY web/.yarn ./.yarn
RUN yarn install --immutable
COPY web/ .
RUN yarn export


FROM golang:${GO_VERSION}-alpine AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=node-builder /app/dist ./web/dist
RUN go build -o golang-aio .


FROM scratch
WORKDIR /app
COPY --from=go-builder /app/golang-aio .
ENTRYPOINT ["./golang-aio"]
EXPOSE 8080

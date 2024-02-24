FROM golang:1.22-alpine3.19 AS buildStage
WORKDIR /api-gateway
COPY . ./
RUN go mod download
RUN go build -o ./api-gateway ./cmd

FROM scratch AS release-stage 
WORKDIR /
COPY --from=buildStage /api-gateway/api-gateway /api-gateway
COPY --from=buildStage /api-gateway/dev.env /

EXPOSE 3000

ENTRYPOINT [ "/api-gateway" ]
FROM --platform=amd64 golang:alpine as builder
ENV VERSION v1.0.0
WORKDIR /app
COPY go.mod go.sum  /app/
RUN go mod download
COPY cmd /app/cmd
COPY pkg /app/pkg
COPY _resources /app/_resources
RUN ls
RUN go build -C cmd/ -buildvcs=false -v -ldflags="-X 'main.Version=$VERSION'"  -o go8ball

FROM --platform=amd64 alpine as prod
WORKDIR /app
COPY --from=builder /app/cmd/go8ball /app
COPY _resources /app/_resources
EXPOSE 8080
CMD [ "/app/go8ball" ]
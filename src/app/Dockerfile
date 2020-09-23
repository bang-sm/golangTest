FROM golang:1.15.2 AS build

WORKDIR /src/
COPY hello.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
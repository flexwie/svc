FROM golang:1.19 as build
ARG SVC

WORKDIR /app

COPY . .
RUN cd $SVC && go mod download && cd ..

RUN GOOS=linux go build -o /svc $SVC/cmd/main.go

FROM alpine

WORKDIR /app
COPY --from=build /svc .
# TODO: copy env config files

CMD ["/app/svc"]

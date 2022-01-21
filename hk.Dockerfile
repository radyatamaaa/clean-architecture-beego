FROM golang:1.14-alpine as builder

# Utilities
RUN apk add git gcc g++ tzdata zip ca-certificates
# Add dep for package management
RUN go get -u -f -v github.com/golang/dep/...

#set workdir
RUN mkdir -p /go/src/PROJECT
WORKDIR /go/src/PROJECT

# COPY go.mod and go.sum files to the workspace
COPY go.mod . 
COPY go.sum .

## Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
RUN go mod tidy

COPY . .
RUN go get
RUN go test ./... -coverprofile=coverage.out
RUN golangci-lint run

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/PROJECT main.go 

# final stage
FROM alpine:latest

#Env                                                                                                                                                2d21h ✖ ⚑ ◒   
ENV TIMEZONE Asia/Jakarta

ARG BRANCH
# set environment based on args
ENV BRANCH $BRANCH
#set timezone

#set timezone
RUN apk --no-cache add tzdata && echo "Asia/Jakarta" > /etc/timezone                                                                                2d21h ✖ ⚑ ◒   
RUN apk add --update tzdata && \
cp /usr/share/zoneinfo/${TIMEZONE} /etc/localtime && \
echo "${TIMEZONE}" > /etc/timezone && apk del tzdata

EXPOSE $PORT
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/PROJECT .

#set env file
COPY --from=builder /go/src/PROJECT/.env.${BRANCH} .env


RUN printf "#!/bin/sh\n\nwhile true; do\n\techo \"[INFO] Starting Service at \$(date)\"\n\t(./PROJECT >> ./history.log || echo \"[ERROR] Restarting Service at \$(date)\")\ndone" > run.sh
RUN printf "#!/bin/sh\n./run.sh & tail -F ./history.log" > up.sh
RUN chmod +x up.sh run.sh
CMD ["./up.sh"]

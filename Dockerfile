FROM golang:1.23

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.35.2
RUN mkdir -p "explore"
RUN protoc --go_out=./explore --go_opt=paths=source_relative \
         --go-grpc_out=./explore --go-grpc_opt=paths=source_relative \
         ./proto/*.proto --experimental_allow_proto3_optional

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Run
CMD ["/docker-gs-ping"]

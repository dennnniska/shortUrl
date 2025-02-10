FROM golang as builder
WORKDIR /shortUrl
ENV DB_PASSWORD=qwerty
ENV POSTGRES_USER=postgres
ENV POSTGRES_DB=postgres
ENV CONFIG_PATH=./config/local.yaml
ENV POSTGRES_PASSWORD=qwerty
# Copy the Go module files
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o /shortUrl ./cmd/shortUrl

FROM postgres
ENV DB_PASSWORD=qwerty
ENV POSTGRES_USER postgres
ENV POSTGRES_DB postgres
ENV CONFIG_PATH=./config/docker.yaml
ENV POSTGRES_PASSWORD qwerty
ENV POSTGRES_HOST localhost
ENV POSTGRES_PORT 5436
WORKDIR /shortUrl
COPY --from=builder /shortUrl /shortUrl
ENTRYPOINT  ["./shortUrl" ]
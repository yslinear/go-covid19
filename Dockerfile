# Golang Server's Dockerfile
FROM golang:1.18.1-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh build-base

# Add Maintainer Info
LABEL maintainer="Ying-Shan Lin <yslinear@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy && go mod download

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags '-extldflags "-static"' -o main .

# Run the executable
CMD ["./main"]

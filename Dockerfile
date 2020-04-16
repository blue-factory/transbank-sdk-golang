# Base container for compile service
FROM golang:alpine AS builder

# Define service name
ARG SVC=webpay-api

# Install dependencies
RUN apk add make

# Go to builder workdir
WORKDIR /go/src/github.com/microapis/${SVC}/

# Copy go modules files
COPY go.mod .
COPY go.sum .

# Install dependencies
RUN go mod download

# Copy all source code
COPY . .


# Compile service
RUN make linux

#####################################################################
#####################################################################

# Base container for run service
FROM alpine

# Define service name
ARG SVC=webpay-api

# Go to workdir
WORKDIR /src/${SVC}

# Copy binary
COPY --from=builder /go/src/github.com/microapis/${SVC}/bin/${SVC} /usr/bin/${SVC}

# Expose service port
EXPOSE 5040

# Run service
CMD ["/bin/sh", "-l", "-c", "$SVC"]
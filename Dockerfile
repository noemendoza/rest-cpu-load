# Add image from and build
FROM golang:1.16.5 as build

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

#Make project
RUN make build


# Start app
#CMD ["/bin/bash"]
CMD ["./build/rest-cpu-load"]
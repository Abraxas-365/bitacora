FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Luis F. Miranda"

# Install git.
# Git is required for fetching the dependencies.
# Setup folders
RUN mkdir /Backend
WORKDIR /Backend

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Install the package
RUN go mod tidy

# Build the Go app
RUN go build cmd/main.go

# Expose port 3000 to the outside world
EXPOSE 1234:1234

# Run the executable
CMD ["./main"]

# Use the official Golang image as the base image
FROM golang:1.22 as build-stage

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o book-service .

# Create a minimal runtime image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files from the builder image
COPY --from=build-stage /app/book-service /app/book-service
COPY /config.yaml /app/config.yaml

# Expose the port on which the Go application will run
EXPOSE 8081

# Define the command to run the Go application
CMD ["./book-service"]
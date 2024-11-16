# Use the official Golang image as the base
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to download dependencies
COPY go.mod  ./

# Download all dependencies
RUN go mod download

# Copy the entire project code into the container
COPY . .

# Build the application inside the container
RUN go build -o main .

# Expose the port your application will use (replace 8080 with your actual port if different)
EXPOSE 8080

# Command to run your application when the container starts
CMD ["./main"]

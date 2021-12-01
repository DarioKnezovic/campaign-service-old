FROM golang:1.16.5 as development

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# Install Reflex for development
RUN go install github.com/cespare/reflex@latest

# Get modules
RUN go get github.com/gin-gonic/gin

# Expose port
EXPOSE 4000

# Start app
CMD reflex -r '\.go' go run main.go --start-service
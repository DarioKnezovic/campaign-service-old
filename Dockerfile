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
RUN go get github.com/joho/godotenv
RUN go get github.com/lib/pq
RUN go get github.com/jinzhu/gorm
RUN go get golang.org/x/sys

# Expose port
EXPOSE 4000

# Start app
CMD reflex -r '\.go' go run main.go --start-service
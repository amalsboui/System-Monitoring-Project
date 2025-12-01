FROM golang:1.25

WORKDIR /app

# Copy ONLY required folders
COPY go.mod go.sum ./
COPY kafkahelper ./kafkahelper
COPY models ./models
COPY processor ./processor
COPY utils ./utils
COPY consumer ./consumer

# Create logs folder
RUN mkdir -p /app/logs


# Download dependencies
RUN go mod download

# Build consumer
RUN cd consumer && go build -o /consumer

CMD ["/consumer"]

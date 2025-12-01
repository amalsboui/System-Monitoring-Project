FROM golang:1.25

WORKDIR /app

# Copy ONLY required folders
COPY go.mod go.sum ./
COPY kafkahelper ./kafkahelper
COPY models ./models
COPY utils ./utils
COPY producer ./producer

# Create logs folder
RUN mkdir -p /app/logs


# Download dependencies
RUN go mod download

# Build producer
RUN cd producer && go build -o /producer

CMD ["/producer"]

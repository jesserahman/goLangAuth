FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code into the working directory
COPY *.go ./

# Add all your sources to container
COPY . ./
RUN ls

# Build
RUN go build -o /goLangAuth

EXPOSE 8081

# Run
CMD [ "/goLangAuth" ]

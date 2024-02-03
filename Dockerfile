# Use a multi-stage build to keep the final image as small as possible
# Start with a Node.js base image for building the UI
FROM node:latest as uibuilder

# Set the working directory in the container
WORKDIR /app

# Copy the UI directory to the container
COPY ui ./ui
COPY templates ./templates

# Install npm dependencies and build the UI
RUN cd ui && npm install && npm run build

# Use a Go base image for the Go application
FROM golang:latest as gobuilder

# Enable Go modules
ENV GO111MODULE=on

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o cloud .

# Final stage: Use a small base image
FROM alpine:latest

WORKDIR /root/

# Copy the Go executable from the gobuilder stage
COPY --from=gobuilder /go/src/app/cloud .

# Copy the templates directory
COPY --from=gobuilder /go/src/app/templates ./templates

# Copy the static directory
COPY --from=gobuilder /go/src/app/static ./static

# Copy the built UI from the uibuilder stage
COPY --from=uibuilder /app/static/dist ./static/dist

# Expose the port your app runs on
EXPOSE 8080

ENV APP_ENV=production

# Command to run the executable
ENTRYPOINT ["/root/cloud"]
FROM golang:1.21.3-bullseye AS builder

COPY . /dineflow-notification-service
WORKDIR /dineflow-notification-service
RUN go mod tidy
RUN go build

# Create a new stage for the final image
FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y ca-certificates wget
RUN wget -O /usr/local/bin/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

COPY --from=builder /dineflow-notification-service/dineflow-notification-service /app/dineflow-notification-service

EXPOSE 8093

CMD ["wait-for-it.sh", "rabbit:5672", "--", "/app/dineflow-notification-service"]
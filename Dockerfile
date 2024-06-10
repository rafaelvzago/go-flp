# Step 1: Build
FROM golang:1.20 AS build
WORKDIR /src
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /app/app
RUN mkdir -p /app/data

# Step 2: Run
FROM scratch AS runtime
COPY --from=build /app /app
CMD ["/app/app"]
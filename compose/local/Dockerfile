From golang:1.20-bullseye

COPY ./compose/local/entrypoint /entrypoint
RUN sed -i 's/\r$//g' /entrypoint
RUN chmod +x /entrypoint

COPY ./compose/local/start /start
RUN sed -i 's/\r$//g' /start
RUN chmod +x /start

WORKDIR /build

COPY go.mod go.sum main.go ./
COPY app ./app
COPY web/web.go ./web/web.go
COPY web/dist ./web/dist/

RUN go mod download && \
    go install github.com/cosmtrek/air@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest
 
ENTRYPOINT ["/entrypoint"]



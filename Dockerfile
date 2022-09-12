FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN go build -o ./server ./cmd/app 


FROM golang:1.18

WORKDIR /app
ENV GIN_MODE=release

COPY --from=builder /app/static/ ./static
COPY --from=builder /app/server ./

RUN chmod 777 /app/server 
CMD [ "./server" ]
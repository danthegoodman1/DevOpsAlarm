# Change image appropriately for the board you are using, see options at: https://www.balena.io/docs/reference/base-images/base-images-ref/
FROM balenalib/raspberry-pi-golang as build

WORKDIR /app

COPY go.* /app/

RUN go mod download

COPY . .

RUN go build -o /app/devOpsAlarm

# Change image appropriately for the board you are using, see options at: https://www.balena.io/docs/reference/base-images/base-images-ref/
FROM balenalib/raspberry-pi-debian
COPY --from=build /app/devOpsAlarm /app/

CMD ["/app/devOpsAlarm" ]

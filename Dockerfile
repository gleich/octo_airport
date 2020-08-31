FROM golang:1.15-alpine3.12

# Meta data:
LABEL maintainer="matthewgleich@gmail.com"
LABEL description="🐙 CLI to view the status of GitHub actions for all your repos"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing dependencies/
RUN go get -v -t -d ./...

# Build the app
RUN go build -o app .

CMD ["./app"]

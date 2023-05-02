FROM golang:1.20.0 as gobase

WORKDIR /usr/src/app

RUN apt-get update
RUN apt-get -y install curl gnupg
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && apt-get install -y nodejs

COPY . .

RUN docker/setup-npm-env.sh
RUN cd ./web/static/src && npm install

RUN go mod tidy
RUN make build

FROM golang:alpine AS build-env
WORKDIR /opt/src
LABEL stage=builder
ENV GO111MODULE=on
ADD  . .
RUN go build -o /opt/src/bin/restmock /opt/src/cmd/restmock/


FROM alpine
WORKDIR /app/
USER restmock
COPY --from=build-env /opt/src/bin/restmock .
COPY ./config/conf.yaml /etc/restmock/conf.yaml
EXPOSE 8080
CMD [ "/app/restmock"]

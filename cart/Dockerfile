FROM golang:1.16-alpine
RUN mkdir /opt/app
WORKDIR /opt/app
COPY  src/ . 
CMD go mod init cart/src
RUN go build -o /src
CMD [ "/src" ]

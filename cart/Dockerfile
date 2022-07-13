FROM golang:1.16-alpine
RUN mkdir /opt/app
WORKDIR /opt/app
COPY  src/ . 
RUN go build -o /src
CMD [ "/src" ]

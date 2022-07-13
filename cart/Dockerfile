FROM golang:1.16-alpine
RUN mkdir /opt/app
WORKDIR /opt/app
COPY  cart/src/ . 
RUN go build -o /cart
CMD [ "/cart" ]

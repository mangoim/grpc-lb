FROM alpine:3.14
RUN mkdir /app
WORKDIR /app
ADD bin/grpc-client-svc /app

EXPOSE 8080

CMD ["/app/grpc-client-svc"]
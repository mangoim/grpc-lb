FROM alpine:3.14
RUN mkdir /app
WORKDIR /app
ADD bin/grpc-server-svc /app

EXPOSE 8081

CMD ["/app/grpc-server-svc"]
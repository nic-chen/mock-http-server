FROM debian:bullseye-slim

WORKDIR /usr/local/mock-server

COPY ./mock-server ./

CMD ["/usr/local/mock-server/mock-server"]

FROM alpine:latest

ADD data /data

ENV WISHLIST_FE_PORT 5000
ENV WISHLIST_BASE_URL https://wishlist.example.ch

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

CMD ["/data/run_server_linux"]

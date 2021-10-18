FROM alpine:latest
RUN mkdir /data
COPY run_server_linux /data/run_server_linux
ADD templates /data/templates
COPY certs /certs
ENV WISHLIST_FE_PORT 5000
ENV WISHLIST_SSL_CERT /certs/fullchain.pem
ENV WISHLIST_SSL_KEY /certs/privkey.pem
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /data
CMD ["./run_server_linux"]

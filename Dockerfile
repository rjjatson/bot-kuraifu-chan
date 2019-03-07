FROM alpine:3.7

ENV HOME /home
ENV CONTENT_DIR content
ENV EXECUTABLE bin

RUN apk add --no-cache curl
RUN curl --remote-name --time-cond cacert.pem https://curl.haxx.se/ca/cacert.pem

COPY $EXECUTABLE $HOME/$EXECUTABLE

ENTRYPOINT ["/home/bin"]
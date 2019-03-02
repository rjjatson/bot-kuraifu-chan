FROM alpine:3.7
ENV HOME /home
COPY bin $HOME/bin
USER nobody
ENTRYPOINT ["/home/bin"]
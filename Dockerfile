FROM scratch

WORKDIR /app

ADD ./main /app

EXPOSE 4040

ENTRYPOINT [ "./main" ]
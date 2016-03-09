FROM ubuntu
WORKDIR /app
COPY hellogo /app/
EXPOSE 8080
ENTRYPOINT ["./hellogo"]
FROM ubuntu
WORKDIR /main
# copy binary into image
COPY main /main/
EXPOSE 8080
ENTRYPOINT ["./main"]
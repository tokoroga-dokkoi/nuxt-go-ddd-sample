FROM golang:latest

RUN mkdir -p app
WORKDIR /app
COPY ./src .

#set variable　for HotReload
ENV GO111MODULE=on

#Only get this library instead of these library
RUN go get github.com/pilu/fresh

# Expose default port (3306)
EXPOSE 9000

# CMD [ "./main" ]
CMD ["fresh"]

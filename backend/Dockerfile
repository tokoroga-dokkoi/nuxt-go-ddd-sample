FROM golang:latest

RUN mkdir -p app
WORKDIR /app
COPY . .

#set variable　for HotReload
ENV GO111MODULE=on

#Only get this library instead of these library
RUN go get github.com/pilu/fresh 

# Expose default port (3306)
EXPOSE 9000 

WORKDIR /go/src/github.com/tnkyk/clean_book_go
# CMD [ "./main" ]
CMD ["fresh"]
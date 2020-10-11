FROM golang:latest

# set GOPATH
RUN mkdir -p /go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

WORKDIR /go/src/github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/
ADD ./src /go/src/github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/
#set variable　for HotReload
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
RUN go get -u github.com/rubenv/sql-migrate/sql-migrate
# Copy Entrypoint
COPY entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]


# Expose default port
EXPOSE 3001

# CMD [ "./main" ]
CMD ["fresh"]

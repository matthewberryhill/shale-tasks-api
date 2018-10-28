FROM iron/go:dev

WORKDIR /app

# Build API
ENV SRC_DIR=/go/src/github.com/matthewberryhill/shale-tasks-api
ADD . $SRC_DIR
RUN cd $SRC_DIR; go get
RUN cd $SRC_DIR; go build -o api; cp api /app/

ENTRYPOINT ./api

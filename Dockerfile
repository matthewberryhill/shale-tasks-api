FROM iron/go:dev

WORKDIR /app
ARG MONGO_ARG
ENV MONGO=$MONGO_ARG

# Build API
ENV SRC_DIR=/go/src/github.com/matthewberryhill/shale-tasks-api
ADD . $SRC_DIR
ADD ./config ./config
RUN cd $SRC_DIR; go get
RUN cd $SRC_DIR; go build -o api; cp api /app/

ENTRYPOINT ./api -env=$MONGO

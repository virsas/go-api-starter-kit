# Build container run following command
# source .env
# docker build --build-arg api_port=$API_PORT --build-arg prom_port=$PRM_PORT Dockerfile 
FROM alpine:3.15

# default values 8080 and 8081, if you add --build-arg it will be overwritten
ARG api_port=8080
ARG prom_port=7081

WORKDIR /app

# install tzdata and configure container with UTC timezone
RUN apk add --no-cache --update tzdata && \
    rm -rf /var/cache/apk/*
RUN cp /usr/share/zoneinfo/UTC /etc/localtime && \
    echo "UTC" >  /etc/timezone

# upload keys (eg JWT or DKIM keys would go here if needed)
ADD keys ./keys
# upload templates (eg PDF or Emailing templates)
ADD templates ./templates
# upload migrations for database
ADD migrations ./migrations
# upload assets directory with all images
ADD assets ./assets

# copy the service
COPY main .

# check and expose application ports
ENV API_PORT=$api_port
ENV PROM_PORT=$prom_port
EXPOSE $API_PORT $PROM_PORT

# run the service
CMD ["/app/main"]
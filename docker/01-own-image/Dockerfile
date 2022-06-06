# use an base image
FROM alpine

# download and install dependency
RUN apk add --update redis
RUN apk add --update gcc


# tell the image what to do when it starts as a container
CMD ["redis-server"]
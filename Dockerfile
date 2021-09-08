# Start from golang v1.13.4 base image to have access to go modules
FROM golang:1.15.3

# create a working directory
WORKDIR /app

# Fetch dependencies on separate layer as they are less likely to
# change on every build and will therefore be cached for speeding
# up the next build
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy source from the host to the working directory inside
# the container
COPY . .

RUN go build -o /uniswap181

# This container exposes port 7777 to the outside world
EXPOSE 8000

CMD ["/uniswap181"]

# TODO: Create a multistage build if time permits
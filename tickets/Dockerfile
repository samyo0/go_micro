FROM golang

WORKDIR '/app'

# COPY go.mod .
# COPY go.sum .
COPY . .
RUN go get .


CMD ["make", "run", "main.go"]
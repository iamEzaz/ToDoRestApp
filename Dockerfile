# FROM golang:latest
# WORKDIR /app/src/task1
# ENV GOPATH=/app
# COPY . /app/src/task1
# RUN go get -u github.com/go-sql-driver/mysql
# RUN go get -u github.com/jinzhu/gorm
# RUN go get -u github.com/gorilla/mux
# RUN go get -u github.com/gorilla/handlers
# RUN go build -o main .
# CMD [ "./main" ]

FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
#RUN go get github.com/stretchr/testify/assert
RUN go get go.uber.org/zap
RUN go build -o main .
CMD ["/app/main"]
 
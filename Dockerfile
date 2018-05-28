FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/github.com/ujjkumsi/docker-go/app
COPY ./best-practices /go/src/github.com/ujjkumsi/docker-go/best-practices
COPY ./models /go/src/github.com/ujjkumsi/docker-go/models
COPY ./dao /go/src/github.com/ujjkumsi/docker-go/dao
WORKDIR /go/src/github.com/ujjkumsi/docker-go/app

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8080
# FROM golang
# WORKDIR /usr/src/app
# ENV MODE=production
# COPY . ./
# RUN go build .
# CMD [ "./poly-career-back" ]
# EXPOSE 4300


##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.19-alpine AS build

# create a working directory inside the image
ENV MODE=production
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod go.sum ./

# download Go modules and dependencies
RUN go mod download

COPY . .

# compile application
RUN  go build -o main .


##
## STEP 2 - DEPLOY
##
FROM alpine:latest
RUN apk --no-cache add ca-certificates

ENV MODE=production
WORKDIR /

# COPY --from=build /app/main /main
COPY --from=build /app/ /
COPY --from=build /app/.env.production /.env.production

EXPOSE 4300

CMD ["/main"]

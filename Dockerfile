
FROM node AS web

COPY web ./
RUN  "npm build"

FROM golang AS server
COPY --from=web ./public ./public
EXPOSE 3000
RUN "go run main.go"

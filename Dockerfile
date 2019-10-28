FROM node:10.16-alpine as nodeBuilder
WORKDIR /app
COPY frontend/ .
RUN npm config set unsafe-perm true
RUN npm install -g @angular/cli
RUN npm install && npm run build

FROM golang:1.13.2-alpine AS goBuilder
WORKDIR /builder
ADD backend/ .
RUN go build -o mbb
RUN ls -l

FROM alpine
WORKDIR /app
COPY --from=goBuilder /builder/mbb /app/
RUN ls -l
COPY --from=nodeBuilder /app/dist/frontend/ assets/
COPY backend/migrations/ migrations/
HEALTHCHECK --interval=5s --timeout=3s --retries=5 CMD curl --fail http://localhost:8080/api/alive || exit 1
ENTRYPOINT ["./mbb"]

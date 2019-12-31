FROM node:10.16-alpine as nodeBuilder
WORKDIR /app
COPY frontend/ .
RUN npm config set unsafe-perm true
RUN npm install -g @angular/cli
RUN npm install --production && npm run build

FROM golang:1.13.5-alpine AS goBuilder
WORKDIR /builder
ADD backend/ .
RUN go build -o mbb

FROM alpine
WORKDIR /app
COPY --from=goBuilder /builder/mbb /app/
COPY --from=nodeBuilder /app/dist/frontend/ assets/
COPY backend/migrations/ migrations/
ENTRYPOINT ["./mbb"]

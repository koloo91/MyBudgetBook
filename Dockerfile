FROM node:10.16-alpine as node
WORKDIR /app
COPY frontend/ .
RUN npm install -g @angular/cli
RUN npm install && npm run build

FROM openjdk:11.0.5-jdk as java
WORKDIR /app
COPY . .
COPY --from=node /app/dist/frontend/ src/main/resources/static/
RUN ./gradlew build

FROM openjdk:11-jre-slim
WORKDIR /app
COPY --from=java /app/build/libs/*.jar app.jar
ENTRYPOINT ["java","-jar","app.jar"]

# For Java 11, try this
FROM openjdk:17-jdk-slim-buster

# Refer to Maven build -> finalName
ARG JAR_FILE=target/configmap-*.jar


# cd /opt/app
WORKDIR /app

# cp target/spring-boot-web.jar /opt/app/app.jar
COPY ${JAR_FILE} configmap-test.jar

# java -jar /opt/app/app.jar
ENTRYPOINT ["java","-jar","configmap-test.jar"]
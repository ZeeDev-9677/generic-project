FROM alpine:latest
RUN apk --no-cache add openjdk17-jdk openjdk17-jmods
ENV JAVA_HOME=/opt/java-minimal

ENV PATH="$PATH:$JAVA_HOME/bin"

COPY /home/ec2-user/Dockerfile/dockerzee.jar dockerzee.jar

ENTRYPOINT ["java","-jar","/dockerzee.jar"]

# sudo docker build -t dockerzee .
# sudo docker run --rm -d -p8080:8080  dockerzee:latest

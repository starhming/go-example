FROM ubuntu:latest
LABEL authors="minghai1024"
RUN apt-get update; apt-get -y install fortune
ADD fortuneloop.sh /bin/fortuneloop.sh
RUN chmod +x /bin/fortuneloop.sh
ENTRYPOINT /bin/fortuneloop.sh

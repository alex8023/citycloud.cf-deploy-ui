#
#  citycloud.cf-deploy-ui Dockerfile
#
FROM  ubuntu:latest

MAINTAINER hanghy

WORKDIR /home/
RUN useradd -d /home/ubuntu -m -s /bin/bash ubuntu

WORKDIR ubuntu/webapp

COPY ./ ./

RUN chown -R ubuntu:ubuntu /home/ubuntu

USER ubuntu

EXPOSE 8080

CMD ["nohup","./citycloud.cf-deploy-ui","&"]




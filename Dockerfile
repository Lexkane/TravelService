FROM ubuntu:latest

COPY src/TravelService /opt/TravelService
COPY config.json /opt/TravelService/config/

RUN chmod +x /opt/TravelService/TravelService
EXPOSE 8080

WORKDIR /opt

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

CMD cat /opt/TravelService/config/config.json && /wait && /opt/TravelService/TravelService -config /opt/TravelService/config/config.json

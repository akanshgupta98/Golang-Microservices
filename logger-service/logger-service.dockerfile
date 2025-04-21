FROM alpine:latest
RUN mkdir /app
COPY ./loggerApp /app
RUN chmod +x /app/loggerApp
CMD [ "/app/loggerApp" ]





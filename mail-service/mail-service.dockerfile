FROM alpine:latest
RUN mkdir /app

COPY ./mailApp /app/
# RUN mkdir -p /app/templates
COPY ./templates /templates

CMD [ "/app/mailApp" ]

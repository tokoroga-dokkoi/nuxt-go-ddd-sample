# Nginx Container
# 開発環境用　本番環境ではBeanstalkにインストールされてるNginxを利用する
FROM nginx:latest

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

RUN rm -f /etc/nginx/conf.d/*
RUN mkdir /var/lib/nginx
RUN chown nginx:nginx /var/lib/nginx

ADD ./config/proxy.conf /etc/nginx/conf.d/proxy.conf

EXPOSE 8080

#CMD /usr/sbin/nginx -g 'daemon off:' -c /etc/nginx/nginx.conf

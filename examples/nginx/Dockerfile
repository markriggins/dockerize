FROM nginx:1.9
MAINTAINER Jason Wilder mail@jasonwilder.com

RUN apt-get update -qq && apt-get install -y vim

# RUN wget https://github.com/markriggins/dockerize/releases/download/v0.2.0/dockerize-linux-amd64-v0.0.4.tar.gz
# RUN tar -C /usr/local/bin -xvzf dockerize-linux-amd64-v0.2.0.tar.gz

COPY .dockerize_linux_amd64.tar.gz /tmp/
RUN  cd /usr/local/bin && \
    tar -xzv --strip-components=1 -f /tmp/.d*.gz dockerize_linux_amd64/dockerize && \
    rm /tmp/.d*.gz &&  \
    echo "daemon off;" >> /etc/nginx/nginx.conf

COPY default.conf.tmpl /etc/nginx/conf.d/default.conf.tmpl

COPY overlays /tmp/overlays
 
EXPOSE 80

ENV DEPLOYMENT_ENV=staging

#  NOTE: These options do not work well on virtualbox osx -- it pegs one CPU
#     "-poll", \
#     "-stdout", "/var/log/nginx/access.log", \
#     "-stderr", "/var/log/nginx/error.log", \

ENTRYPOINT [ "dockerize", \
     "-template", "/etc/nginx/conf.d/default.conf.tmpl:/etc/nginx/conf.d/default.conf", \
     "-poll", \
     "-stdout", "/var/log/nginx/access.log", \
     "-stderr", "/var/log/nginx/error.log" ]


CMD [ "--", "nginx"]

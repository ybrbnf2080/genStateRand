FROM nginx

WORKDIR /app

ADD ybrbnf@localhost:/home/ybrbnf/tmp/ny.tar.gz ./

RUN ls 

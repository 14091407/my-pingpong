FROM fholzer/nginx-brotli:v1.19.1
RUN rm /etc/nginx/conf.d/*
COPY ./conf/default.conf /etc/nginx/conf.d
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
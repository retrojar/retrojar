FROM node:slim
WORKDIR /usr/src/app
RUN npm install --quiet --global vue-cli
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build


FROM nginx
WORKDIR /srv/app
COPY --from=0 /usr/src/app/dist/ .
COPY docker/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf.template
COPY docker/nginx_start.sh /usr/local/bin/
CMD ["nginx_start.sh"]

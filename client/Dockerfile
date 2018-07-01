FROM node:slim

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

RUN npm install --quiet --global vue-cli

COPY . .

EXPOSE 8080
CMD [ "npm", "start" ]

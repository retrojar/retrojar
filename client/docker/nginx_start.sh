#!/usr/bin/env bash

envsubst '$API_HOST' < /etc/nginx/conf.d/default.conf.template > /etc/nginx/conf.d/default.conf

nginx -g "daemon off;"

#!/bin/bash

# prepend 'VITE_' on each line
cat .env | sed -e 's/^/VITE_/' | sed -e 's/VITE_$//' > ./web/static/src/.env

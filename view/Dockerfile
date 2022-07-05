FROM node:16-alpine3.15
RUN mkdir /app
WORKDIR /app

# downloading the packages
COPY package.json package-lock.json ./
RUN npm ci
COPY . .

# port
EXPOSE 3000

# run
CMD ["npm", "start"]
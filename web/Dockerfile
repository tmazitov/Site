FROM node:10-alpine as builder
WORKDIR /vue-ui

# Copy the package.json and install dependencies
COPY package*.json ./
RUN npm install
RUN npm install yarn

# Copy rest of the files
COPY . .

# Build the project
RUN npm run build

EXPOSE 8080

# production stage
CMD [ "yarn", "serve" ]

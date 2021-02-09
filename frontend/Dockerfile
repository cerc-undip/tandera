FROM node:15.7.0-buster

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
# COPY package-lock.json ./
RUN npm install
RUN npm install react-scripts@3.4.1 -g --silent

# more dependencies here

# add app
COPY . ./

# start app
CMD ["npm", "start"]

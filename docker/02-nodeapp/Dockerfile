#base image
FROM node:14.16.0-alpine

#make a directory and make this working directory
WORKDIR /usr/app

#COPY ./local to ./container 
#just copies the package.json file to run the npm install for dependencies
#this approach reduces the re-running of npm run install when building 
#again after some file change in the project
COPY ./package.json ./

#install the depemdencies
RUN npm install
#Copy the all other files except package.json
COPY ./ ./

#default commands
CMD ["npm", "start"]
FROM node:18.12.1
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
WORKDIR /app
COPY package.json package.json
COPY package-lock.json package-lock.json

RUN npm install
COPY . .
CMD ["npm", "run", "start"]
FROM node:16-alpine3.11 AS build-env

WORKDIR /work

ARG ENV_FILE_NAME=".env.development"
COPY ./${ENV_FILE_NAME} ./.env

COPY ./package.json ./yarn.lock  ./
RUN yarn install

COPY . .
RUN yarn build

FROM node:16-alpine3.11

WORKDIR /work
COPY --from=build-env /work/.output /work/.output

CMD ["node", ".output/server/index.mjs"]
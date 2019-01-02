# ever-post-mixin-bot

post your content on Telegraph and IPFS,  based on Mixin Network

## server


use `./ever-post-mixin-bot --service http` start the http server
and `./ever-post-mixin-bot --service blaze` start the bot message server


## web

web build with [quasar framework](https://quasar-framework.org/)

```
cd web

npm install || yarn

# Node.js >= 8.9.0 is required.
$ yarn global add quasar-cli
# or:
$ npm install -g quasar-cli

quasar dev # develop mode

quasar build # product mode and output /dist/spa-mat

```
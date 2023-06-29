# ベースイメージを指定
FROM node

# node.js の環境変数を定義する
# 本番環境では production
ENV NODE_ENV=development

# 雛形を生成するのに必要なパッケージのインストール

# ディレクトリを移動する
WORKDIR /var/www/application

COPY ./front /var/www/application

RUN yarn install
RUN npm install pm2 -g

# ポート3000番を開放する
EXPOSE 3000

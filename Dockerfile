FROM golang:1.23 AS test-backend

RUN useradd -ms /bin/sh user
USER user

WORKDIR /app

COPY --chown=user . .
RUN make install-backend-dependencies

CMD ["make", "test-backend"]


FROM node:20.18.0 AS test-unit-frontend

RUN useradd -ms /bin/sh user
USER user

WORKDIR /app

COPY --chown=user . .
RUN make install-frontend-dependencies

CMD ["make", "test-frontend-unit"]


FROM node:20.18.0 AS test-e2e-frontend
RUN apt update
RUN apt install -y xvfb libgtk-3-dev libwebkit2gtk-4.0-dev

RUN wget https://go.dev/dl/go1.23.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz

RUN useradd -ms /bin/sh user
USER user

RUN mkdir /home/user/go
ENV PATH=$PATH:/usr/local/go/bin:/home/user/go/bin
ENV DISPLAY=:1
ENV GOPATH=/home/user/go

WORKDIR /app

COPY --chown=user . .

RUN make install
RUN make frontend-build
RUN make install-frontend-playwright

CMD ["make", "test-frontend-e2e-container"]


FROM golang:bookworm AS linux-build

RUN echo "deb http://deb.debian.org/debian bookworm-backports main" > /etc/apt/sources.list.d/debian-12-backports.list
RUN apt update && apt install -y upx-ucl libwebkit2gtk-4.0-dev

RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash
ENV NVM_DIR="/root/.nvm"

WORKDIR /app

COPY . .

RUN . $NVM_DIR/nvm.sh && cd frontend && nvm install
RUN . $NVM_DIR/nvm.sh && make install

RUN go install mvdan.cc/garble@latest

RUN . $NVM_DIR/nvm.sh && wails build -platform linux/amd64 -upx

RUN ls -lah build/bin

CMD ["tail", "-f", "/dev/null"]

FROM golang:1.25.2-alpine3.22 AS alpine-builder

RUN apk update && apk add -q webkit2gtk-4.1-dev wget bash gcc g++ make

ENV NODE_PACKAGE_URL  https://unofficial-builds.nodejs.org/download/release/v24.5.0/node-v24.5.0-linux-x64-musl.tar.gz

RUN apk add libstdc++
WORKDIR /opt
RUN wget $NODE_PACKAGE_URL
RUN mkdir -p /opt/nodejs
RUN tar -zxvf *.tar.gz --directory /opt/nodejs --strip-components=1
RUN rm *.tar.gz
RUN ln -s /opt/nodejs/bin/node /usr/local/bin/node
RUN ln -s /opt/nodejs/bin/npm /usr/local/bin/npm

WORKDIR /app

COPY . .

RUN make install

RUN touch frontend/build/.gitignore
RUN wails build -platform linux/amd64 -tags webkit2_41

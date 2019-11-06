FROM buildpack-deps:cosmic

### base ###
RUN yes | unminimize \
    && apt-get install -yq \
        asciidoctor \
        bash-completion \
        build-essential \
        htop \
        jq \
        less \
        llvm \
        locales \
        man-db \
        nano \
        software-properties-common \
        sudo \
        vim \
    && locale-gen en_US.UTF-8 \
    && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/*
ENV LANG=en_US.UTF-8

### Gitpod user ###
# '-l': see https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#user
RUN useradd -l -u 33333 -G sudo -md /home/gitpod -s /bin/bash -p gitpod gitpod \
    # passwordless sudo for users in the 'sudo' group
    && sed -i.bkp -e 's/%sudo\s\+ALL=(ALL\(:ALL\)\?)\s\+ALL/%sudo ALL=NOPASSWD:ALL/g' /etc/sudoers
ENV HOME=/home/gitpod
WORKDIR $HOME
# custom Bash prompt
RUN { echo && echo "PS1='\[\e]0;\u \w\a\]\[\033[01;32m\]\u\[\033[00m\] \[\033[01;34m\]\w\[\033[00m\] \\\$ '" ; } >> .bashrc

### Go ###
ENV GO_VERSION=1.13.4 \
    GOPATH=$HOME/go-packages \
    GOROOT=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH
RUN curl -fsSL https://storage.googleapis.com/golang/go$GO_VERSION.linux-amd64.tar.gz | tar -xzv
# user Go packages
ENV GOPATH=/workspace:$GOPATH \
    PATH=/workspace/bin:$PATH

### Node.js ###
ARG NODE_VERSION=12.13.0
ENV PATH=/home/gitpod/.nvm/versions/node/v12.13.0/bin:$PATH
RUN curl -fsSL https://raw.githubusercontent.com/creationix/nvm/v0.33.11/install.sh | bash \
    && bash -c ". .nvm/nvm.sh \
        && npm config set python /usr/bin/python --global \
        && npm config set python /usr/bin/python \
        && npm install -g typescript"
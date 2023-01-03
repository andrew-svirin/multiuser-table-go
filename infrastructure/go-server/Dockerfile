FROM golang:1.19

RUN apt-get update && apt-get install -y \
    sudo

ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ENV GOROOT=/usr/local/go
ENV GOPATH=/var/www/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

ARG USER_ID=1000
ARG GROUP_ID=1000
ARG USER_NAME=www-data
ARG GROUP_NAME=www-data

# Add user
RUN userdel -f ${USER_NAME} \
    && if getent group ${GROUP_NAME} ; then groupdel ${GROUP_NAME}; fi \
    && groupadd -g ${GROUP_ID} ${GROUP_NAME} \
    && useradd -l -u ${USER_ID} -g ${GROUP_NAME} ${USER_NAME} && echo "${USER_NAME}:${GROUP_NAME}" | chpasswd \
    && adduser ${USER_NAME} sudo \
    && install -d -m 0755 -o ${USER_NAME} -g ${GROUP_NAME} "/home/${USER_NAME}" \
    && echo "\n%sudo ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && chown --changes --silent --no-dereference --recursive ${USER_ID}:${GROUP_ID}  \
      "/home/${USER_NAME}" /var/log /var/www || true

USER ${USER_NAME}:${GROUP_NAME}

WORKDIR /var/www/go/multiuser-table-go/server

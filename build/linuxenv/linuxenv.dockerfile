ARG OS_VER

FROM ${OS_VER}

ARG USER_NAME

# Install sudo and add password
# to root user
RUN apt update && \
    apt install sudo && \
    echo 'root:Docker' | chpasswd

# Create user and add it to sudo group
RUN useradd ${USER_NAME} -p ${USER_NAME} -m && \
    echo ${USER_NAME}:${USER_NAME} | chpasswd && \
    usermod -aG sudo ${USER_NAME}

USER ${USER_NAME}

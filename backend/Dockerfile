FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupmod --gid $USER_GID $USERNAME \
    && usermod --uid $USER_UID --gid $USER_GID $USERNAME \
    && chown -R $USER_UID:$USER_GID /home/$USERNAME

WORKDIR /app

COPY . .

# RUN chown -R $USER_UID:$USER_GID /app

RUN go mod download

RUN go build -o main .

CMD ["./main"]

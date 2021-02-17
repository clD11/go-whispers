FROM utensils/opengl:stable

# Install all needed runtime dependencies.
RUN \
  set -xe; \
  apk --update add --no-cache --virtual .runtime-deps \
    bash \
    ffmpeg \
    git \
    gource \
    imagemagick;

# Copy scripts into image
COPY ./entrypoint.sh /usr/local/bin/

# Add executable right to scripts
RUN chmod +x /usr/local/bin/entrypoint.sh

# Create working directories
RUN \
  mkdir /gource; \
  mkdir /gource/logs; \
  mkdir /gource/avatars; \
  mkdir /gource/git_repo; \
  mkdir /gource/git_repos; \
  mkdir /gource/output;

# Set working directory with full access
WORKDIR /gource
RUN chmod -R 777 /gource

# Set our entrypoint.
ENTRYPOINT  ["bash", "/usr/local/bin/entrypoint.sh"]
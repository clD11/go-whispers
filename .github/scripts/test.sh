printf 'Running Test Script'

apk --update add --no-cache --virtual .runtime-deps \
    bash \
    ffmpeg;


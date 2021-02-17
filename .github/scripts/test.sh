#!/bin/bash
printf 'Running Test Script'

apt-get update
apt-get upgrade

ffmpeg -version
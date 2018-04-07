OS=$(shell uname -s)
.DEFAULT_GOAL := run
.SILENT: # silence!!

export PATH := ./bin:$(PATH)

setup:
	mkdir -p bin
ifeq ($(OS), Darwin)
	brew install hugo || brew upgrade hugo
	wget -O ./bin/htmltest https://github.com/wjdp/htmltest/releases/download/v0.8.1/htmltest-osx
else
	curl -sL https://raw.githubusercontent.com/goreleaser/godownloader/master/samples/godownloader-hugo.sh | sh
	wget -O ./bin/htmltest https://github.com/wjdp/htmltest/releases/download/v0.8.1/htmltest-linux
endif
	chmod +x ./bin/*

run:
	hugo server -w

ci:
	rm -rf ./public || true
	hugo server -d ./public >/dev/null &
	htmltest ./public
	pkill hugo

avatar:
	wget -O static/avatar.jpg https://avatars0.githubusercontent.com/u/245435
	convert static/avatar.jpg \
		-bordercolor white -border 0 \
		\( -clone 0 -resize 16x16 \) \
		\( -clone 0 -resize 32x32 \) \
		\( -clone 0 -resize 48x48 \) \
		\( -clone 0 -resize 64x64 \) \
		-delete 0 -alpha off -colors 256 static/favicon.ico

new:
	hugo new "post/$$(date +%Y-%m-%d)-$(filter-out $@,$(MAKECMDGOALS)).md"

%:
	@:

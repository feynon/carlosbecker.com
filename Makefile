OS=$(shell uname -s)
.DEFAULT_GOAL := run
.SILENT: # silence!!

export PATH := ./bin:$(PATH)

setup:
	mkdir -p bin
ifeq ($(OS), Darwin)
	brew install hugo
	wget -O ./bin/htmltest https://github.com/wjdp/htmltest/releases/download/v0.8.0/htmltest-osx
else
	curl -sL https://raw.githubusercontent.com/goreleaser/godownloader/master/samples/godownloader-hugo.sh | sh
	wget -O ./bin/htmltest https://github.com/wjdp/htmltest/releases/download/v0.8.0/htmltest-linux
endif
	chmod +x ./bin/*

run:
	hugo server -w

build:
	rm -rf public
	echo "relativeurls = true" > ./tmp/config.toml
	cat config.toml >> ./tmp/config.toml
	hugo --config ./tmp/config.toml

lint:
	htmltest ./public

ci: build lint

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

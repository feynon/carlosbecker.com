OS=$(shell uname -s)
.DEFAULT_GOAL := run
.SILENT: # silence!!

export PATH := ./bin:$(PATH)

setup:
	mkdir -p bin
ifeq ($(OS), Darwin)
	brew install hugo || brew upgrade hugo
else
	curl -sL https://install.goreleaser.com/github.com/gohugoio/hugo.sh | bash
endif
	curl -sL https://install.goreleaser.com/github.com/ValeLint/vale.sh | bash
	curl -sL https://htmltest.wjdp.uk | bash
	chmod +x ./bin/*

run:
	hugo server -w

ci:
	# TODO: eventually check htmls as well
	$$(which vale) --glob='**/*.md' .
	rm -rf ./public || true
	hugo server -d ./public >/dev/null &
	sleep 2
	$$(which htmltest) ./public
	pkill hugo

avatar:
	wget -O static/avatar.jpg https://avatars0.githubusercontent.com/u/245435
	convert static/avatar.jpg \
		-bordercolor white -border 0 \
		\( -clone 0 -resize 16x16 \) \
		\( -clone 0 -resize 32x32 \) \
		\( -clone 0 -resize 48x48 \) \
		\( -clone 0 -resize 64x64 \) \
		-delete 0 -alpha off -colors 256 static/img/favicon.ico
	convert -resize x120 static/avatar.jpg static/img/apple-touch-icon.png

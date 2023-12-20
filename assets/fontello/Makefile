FONTELLO_HOST ?= https://fontello.com

open:
	curl --silent --show-error --fail --output .fontello --form \
		"config=@./config.json" ${FONTELLO_HOST}
	x-www-browser ${FONTELLO_HOST}/`cat .fontello` | open ${FONTELLO_HOST}/`cat .fontello`

save:
	rm -rf .fontello.src .fontello.zip
	curl --silent --show-error --fail --output .fontello.zip \
		${FONTELLO_HOST}/`cat .fontello`/get
	unzip .fontello.zip -d .fontello.src
	mkdir -p font css
	rsync -raP `find ./.fontello.src -maxdepth 1 -name 'fontello-*'`/font .
	mv `find ./.fontello.src -maxdepth 1 -name 'fontello-*'`/css/fontello-codes.css ./css
	mv `find ./.fontello.src -maxdepth 1 -name 'fontello-*'`/config.json .
	rm -rf .fontello.src .fontello.zip
	
clean:
	rm -rf .fontello.src .fontello.zip
	rm -rf font css

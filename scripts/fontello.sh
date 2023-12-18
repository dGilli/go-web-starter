#!/bin/bash
# open_fontello.sh
# Usage: ./open_fontello.sh session_id

CONFIG_FILE="../assets/fontello/config.json"
SESSION_ID=$(curl -X POST -F "config=@$CONFIG_FILE" https://fontello.com/)

if [ "$1" = "get" ]; then
    OUTPUT="../tmp/fontello.zip"
    curl -o $OUTPUT https://fontello.com/$SESSION_ID/get
    unzip -o $OUTPUT -d ../tmp; rm $OUTPUT

    TMP_DIR="../tmp/$(ls ../tmp | grep 'fontello-')"
    TARGET_DIR="../assets/fontello/"
    rsync -raP $TMP_DIR/font/ $TARGET_DIR/font/
    mv $TMP_DIR/css/fontello-codes.css $TARGET_DIR/css/
    mv $TMP_DIR/config.json $TARGET_DIR
    rm -rf $TMP_DIR
else
    open https://fontello.com/$SESSION_ID
fi

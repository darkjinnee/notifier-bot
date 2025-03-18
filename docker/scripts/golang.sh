#!/usr/bin/env sh

cd "$ROOT" && eval "$(echo "$APP_COMMAND" | sed 's/^\"\(.*\)\"$/\1/g')"

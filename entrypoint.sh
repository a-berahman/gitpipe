#!/bin/bash
set -e

case "$1" in
    bash)
        set -- "$@"
    ;;
    *)
        set -- ./gitpipe "$@"
    ;;
esac

exec "$@"

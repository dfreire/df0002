#!/bin/sh
HERE=`cd -P $(dirname $0) && pwd`

CACHE_DIR=$HERE/site/cache
mkdir -p $CACHE_DIR
rm -rf $CACHE_DIR/*

THUMBS_DIR=$HERE/thumbs
mkdir -p $THUMBS_DIR
rm -rf $THUMBS_DIR/*

# curl http://example.com/page-that-must-be-cached &> /dev/null

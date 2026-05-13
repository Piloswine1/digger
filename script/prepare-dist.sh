#!/bin/sh
set -eu

rm -rf ./inner/rest/dist/*
cp -r ./front/dist/* ./inner/rest/dist/

INPUT=./inner/rest/dist/index.html
OUTPUT=./inner/rest/dist/index.tmpl

sed \
  -e '1i\{{ define "dist/index.tmpl" }}' \
  -e '$a\{{ end }}' \
  -e 's|src="/assets/|src="{{ .prefix }}/assets/|g' \
  -e 's|href="/assets/|href="{{ .prefix }}/assets/|g' \
  -e 's|<div id="app"></div>|<script>window.__APP_CONFIG__={apiBaseUrl:"{{ .apiBaseUrl }}"}</script>\n    <div id="app"></div>|' \
  "$INPUT" > "$OUTPUT"

rm "$INPUT"

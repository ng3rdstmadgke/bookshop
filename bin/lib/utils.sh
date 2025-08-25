#!/bin/bash

function error {
  echo "[Error] $@" >&2
  exit 1
}

function info {
  echo "[Info] $@"
}
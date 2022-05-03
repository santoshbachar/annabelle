#!/bin/sh

name=$1

exec() {
  groupadd "$name"
}

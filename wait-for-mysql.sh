#!/bin/sh

set -e

host="$1"
shift
cmd="$@"

until mysql -h "127.0.0.1:33104" -u "root" -P 33104 -p example -c '\q'; do
  >&2 echo "MySQL is unavailable - sleeping"
  sleep 1
done

>&2 echo "MySQL is up - executing command"
exec $cmd
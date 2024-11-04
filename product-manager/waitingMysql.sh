#!/bin/sh
counter=0
max_retries=5

while ! ping -c 1 mysql > /dev/null 2>&1; do
    counter=$((counter + 1))
    if [ "$counter" -ge "$max_retries" ]; then
        echo "MySQL container did not respond after $max_retries attempts. Exiting."
        exit 1
    fi
    echo "Waiting for MySQL container to respond to ping... (Attempt $counter)"
    sleep 2
done

echo "MySQL container is responding to ping."
exec "$@"

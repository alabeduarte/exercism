FROM php:7.1.3-cli

COPY . /app
WORKDIR /app

RUN chmod +x phpunit.phar
CMD ["./phpunit.phar", "./hello-world/hello-world_test.php", "--colors=always"]

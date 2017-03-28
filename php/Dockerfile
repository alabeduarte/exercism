FROM php:7.1.3-cli

COPY . /app
WORKDIR /app

RUN chmod +x phpunit.phar
CMD ["./phpunit.phar", "./hamming/hamming_test.php", "--colors=always"]

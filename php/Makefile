.PHONY : install_and_run install run clean

full_run: install run

install:
	$(MAKE) clean
	docker build -t exercism-php .

run:
	./docker/restart-docker-image.sh exercism-php

clean:
	./docker/clean-docker-images.sh

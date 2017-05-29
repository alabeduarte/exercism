.PHONY : install_and_run install run clean

full_run: install run

install:
	$(MAKE) clean
	docker build -t exercism-go-lang .

run:
	./docker/restart-docker-image.sh exercism-go-lang

clean:
	./docker/clean-docker-images.sh

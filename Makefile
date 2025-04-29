start:
	docker compose up --build

stop:
	docker compose rm -v -f -s
	docker rmi ticket-booking-api


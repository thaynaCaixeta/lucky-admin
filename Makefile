.PHONY: docker-up docker-down docker-logs

docker-up:
	@docker-compose -f docker-compose-local.yml up -d

docker-down:
	@docker-compose -f docker-compose-local.yml down

docker-logs:
	@docker-compose -f docker-compose-local.yml logs -f

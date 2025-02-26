include api.mk

# COLOR VARIABLE
GREEN=\033[0;32m
RED=\033[0;31m
BLUE=\033[0;34m
LIGHT_BLUE=\033[1;34m
ORANGE=\033[0;33m
NOCOLOR=\033[0m

# STYLE VARIABLE
BLUE_TRIPLE_EQUALS=$(LIGHT_BLUE)===$(NOCOLOR)

# FUNCTION
define log_action
$(BLUE_TRIPLE_EQUALS) $(ORANGE)$(1)$(NOCOLOR) $(BLUE_TRIPLE_EQUALS)
endef
define log_action_end
$(GREEN_TRIPLE_RIGHT_ARROW) $(1) $(GREEN_TRIPLE_LEFT_ARROW)
endef
define check_env
	@if [ -z "$(ENV)" ]; then \
		echo "ERROR: ENV is required. Usage: make $(1) ENV=dev"; \
		exit 1; \
	fi
endef

# ENV VARIABLE
NAME := $(notdir $(CURDIR))
SLUG := $(subst -,_,$(NAME))
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
    ECHO_CMD = echo
else ifeq ($(UNAME_S),Darwin)
    ECHO_CMD = echo
else
    ECHO_CMD = echo -e
endif

.PHONY: setup-hooks
setup-hooks:
	@$(ECHO_CMD) "$(call log_action,Setting up Git Hooks)"
	chmod +x "./script/setup-hooks.sh"
	bash ./script/setup-hooks.sh
	@$(ECHO_CMD) "$(call log_action_end,Git Hooks setup completed)"

.PHONY: setup-env
setup-env:
	$(call check_env,$(MAKECMDGOALS))
	@$(ECHO_CMD) "$(call log_action,Setting up Environment Files)"
	chmod +x "./script/setup-env.sh"
	bash ./script/setup-env.sh $(ENV)
	@$(ECHO_CMD) "$(call log_action_end,Environment files setup completed)"

.PHONY: start
start:
	$(call check_env,$(MAKECMDGOALS))
	@$(ECHO_CMD) "$(call log_action,Generate App Network ($(NAME) | $(ENV)))"
	docker network ls | grep $(SLUG)_backend_$(ENV) || docker network create $(SLUG)_backend_$(ENV)
	@make -C database/sql DATABASES_SQL="$(DATABASES_SQL)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" run-database-sql
	@make -C database/cache DATABASES_CACHE="$(DATABASES_CACHE)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" run-database-cache
	@make -C broker BROKERS="$(BROKERS)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" run-broker
	@make -C application APPLICATIONS="$(APPLICATIONS)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" run-application

.PHONY: stop
stop:
	$(call check_env,$(MAKECMDGOALS))
	@make -C database/sql DATABASES_SQL="$(DATABASES_SQL)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" stop-database-sql
	@make -C database/cache DATABASES_CACHE="$(DATABASES_CACHE)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" stop-database-cache
	@make -C broker BROKERS="$(BROKERS)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" stop-broker
	@make -C application APPLICATIONS="$(APPLICATIONS)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" stop-application
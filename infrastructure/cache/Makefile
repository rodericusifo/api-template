SHELL := /bin/bash

# COLOR VARIABLE
GREEN=\033[0;32m
RED=\033[0;31m
BLUE=\033[0;34m
LIGHT_BLUE=\033[1;34m
ORANGE=\033[0;33m
NOCOLOR=\033[0m

# STYLE VARIABLE
BLUE_TRIPLE_EQUALS=$(LIGHT_BLUE)===$(NOCOLOR)
GREEN_TRIPLE_RIGHT_ARROW=$(GREEN)>>>$(NOCOLOR)
GREEN_TRIPLE_LEFT_ARROW=$(GREEN)<<<$(NOCOLOR)

# FUNCTION
define log_action
$(BLUE_TRIPLE_EQUALS) $(ORANGE)$(1)$(NOCOLOR) $(BLUE_TRIPLE_EQUALS)
endef
define log_action_end
$(GREEN_TRIPLE_RIGHT_ARROW) $(1) $(GREEN_TRIPLE_LEFT_ARROW)
endef

# ENV VARIABLE
AVAILABLE_CACHE := redis
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
    ECHO_CMD = echo
else ifeq ($(UNAME_S),Darwin)
    ECHO_CMD = echo
else
    ECHO_CMD = echo -e
endif

.PHONY: check-cache
check-cache:
	@$(ECHO_CMD) "$(call log_action,Check App Cache ($(NAME) | $(ENV)))"
	@for cache in $(CACHE); do \
		if [[ ! " $(AVAILABLE_CACHE) " =~ " $${cache} " ]]; then \
			$(ECHO_CMD) "$(RED)ERROR: Cache [$${cache}] is not available. Available Cache are: $(AVAILABLE_CACHE)$(NOCOLOR)"; \
			exit 1; \
		fi; \
	done
	@$(ECHO_CMD) "$(call log_action_end,All Caches are available)"

.PHONY: run-cache
run-cache: check-cache
	@$(ECHO_CMD) "$(call log_action,Start App Cache ($(NAME) | $(ENV)))"
	@for cache in $(CACHE); do \
		$(ECHO_CMD) "$(call log_action,Generate App Cache Volumes [$${cache}] ($(NAME) | $(ENV)))"; \
		docker volume ls | grep $(SLUG)_primary_$${cache}_data_$(ENV) || docker volume create --name $(SLUG)_primary_$${cache}_data_$(ENV); \
		docker volume ls | grep $(SLUG)_replica_$${cache}_data_$(ENV) || docker volume create --name $(SLUG)_replica_$${cache}_data_$(ENV); \
		$(ECHO_CMD) "$(call log_action,Start App Cache [$${cache}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) docker compose -p $(NAME)_cache --env-file $${cache}/env/$(ENV).env -f $${cache}/docker-compose.yaml up --build -d; \
	done
	@$(ECHO_CMD) "$(call log_action_end,Cache started successfully)"

.PHONY: stop-cache
stop-cache: check-cache
	@$(ECHO_CMD) "$(call log_action,Stop App Cache ($(NAME) | $(ENV)))"
	@for cache in $(CACHE); do \
		$(ECHO_CMD) "$(call log_action,Stop App Cache [$${cache}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) docker compose -p $(NAME)_cache --env-file $${cache}/env/$(ENV).env -f $${cache}/docker-compose.yaml down; \
	done
	@$(ECHO_CMD) "$(call log_action_end,Cache stopped successfully)"
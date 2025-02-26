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
AVAILABLE_DATABASES_CACHE := redis
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
    ECHO_CMD = echo
else ifeq ($(UNAME_S),Darwin)
    ECHO_CMD = echo
else
    ECHO_CMD = echo -e
endif

.PHONY: check-database-cache
check-database-cache:
	@$(ECHO_CMD) "$(call log_action,Check App Databases Cache ($(NAME) | $(ENV)))"
	@for database_cache in $(DATABASES_CACHE); do \
		if [[ ! " $(AVAILABLE_DATABASES_CACHE) " =~ " $${database_cache} " ]]; then \
			$(ECHO_CMD) "$(RED)ERROR: Database Cache [$${database_cache}] is not available. Available Databases Cache are: $(AVAILABLE_DATABASES_CACHE)$(NOCOLOR)"; \
			exit 1; \
		fi; \
	done
	@$(ECHO_CMD) "$(call log_action_end,All Databases Caches are available)"

.PHONY: run-database-cache
run-database-cache: check-database-cache
	@$(ECHO_CMD) "$(call log_action,Start App Databases Cache ($(NAME) | $(ENV)))"
	@for database_cache in $(DATABASES_CACHE); do \
		$(ECHO_CMD) "$(call log_action,Generate App Database Cache Volumes [$${database_cache}] ($(NAME) | $(ENV)))"; \
		docker volume ls | grep $(SLUG)_primary_$${database_cache}_data_$(ENV) || docker volume create --name $(SLUG)_primary_$${database_cache}_data_$(ENV); \
		docker volume ls | grep $(SLUG)_replica_$${database_cache}_data_$(ENV) || docker volume create --name $(SLUG)_replica_$${database_cache}_data_$(ENV); \
		$(ECHO_CMD) "$(call log_action,Start App Database Cache [$${database_cache}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) docker compose -p $(NAME)_database_cache --env-file $${database_cache}/env/$(ENV).env -f $${database_cache}/docker-compose.yaml up --build -d; \
	done
	@$(ECHO_CMD) "$(call log_action_end,Databases Cache started successfully)"

.PHONY: stop-database-cache
stop-database-cache: check-database-cache
	@$(ECHO_CMD) "$(call log_action,Stop App Databases Cache ($(NAME) | $(ENV)))"
	@for database_cache in $(DATABASES_CACHE); do \
		$(ECHO_CMD) "$(call log_action,Stop App Database Cache [$${database_cache}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) docker compose -p $(NAME)_database_cache --env-file $${database_cache}/env/$(ENV).env -f $${database_cache}/docker-compose.yaml down; \
	done
	@$(ECHO_CMD) "$(call log_action_end,Databases Cache stopped successfully)"
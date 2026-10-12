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
AVAILABLE_DATABASES_SQL := postgres
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
    ECHO_CMD = echo
else ifeq ($(UNAME_S),Darwin)
    ECHO_CMD = echo
else
    ECHO_CMD = echo -e
endif

.PHONY: check-database-sql
check-database-sql:
	@$(ECHO_CMD) "$(call log_action,Check App Databases SQL ($(NAME) | $(ENV)))"
	@for database_sql in $(DATABASES_SQL); do \
		if [[ ! " $(AVAILABLE_DATABASES_SQL) " =~ " $${database_sql} " ]]; then \
			$(ECHO_CMD) "$(RED)ERROR: Database SQL [$${database_sql}] is not available. Available Databases SQL are: $(AVAILABLE_DATABASES_SQL)$(NOCOLOR)"; \
			exit 1; \
		fi; \
	done
	@$(ECHO_CMD) "$(call log_action_end,All Databases SQL are available)"

.PHONY: run-database-sql
run-database-sql: check-database-sql
	@$(ECHO_CMD) "$(call log_action,Start App Databases SQL ($(NAME) | $(ENV)))"
	@for database_sql in $(DATABASES_SQL); do \
		$(ECHO_CMD) "$(call log_action,Generate App Database SQL Volumes [$${database_sql}] ($(NAME) | $(ENV)))"; \
		if [ "$${database_sql}" == "postgres" ]; then \
			chmod +x "postgres/primary/initdb/db.sh"; \
			chmod +x "postgres/primary/initdb/role.sh"; \
		fi; \
		docker volume ls | grep $(SLUG)_primary_$${database_sql}_data_$(ENV) || docker volume create --name $(SLUG)_primary_$${database_sql}_data_$(ENV); \
		docker volume ls | grep $(SLUG)_replica_$${database_sql}_data_$(ENV) || docker volume create --name $(SLUG)_replica_$${database_sql}_data_$(ENV); \
		$(ECHO_CMD) "$(call log_action,Start App Database SQL [$${database_sql}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) docker compose -p $(NAME)_database_sql --env-file $${database_sql}/env/$(ENV).env -f $${database_sql}/docker-compose.yaml up --build -d; \
	done
	@$(ECHO_CMD) "$(call log_action_end,Databases SQL started successfully)"

.PHONY: stop-database-sql
stop-database-sql: check-database-sql
	@$(ECHO_CMD) "$(call log_action,Stop App Databases SQL ($(NAME) | $(ENV)))"
	@for database_sql in $(DATABASES_SQL); do \
		$(ECHO_CMD) "$(call log_action,Stop App Database SQL [$${database_sql}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) docker compose -p $(NAME)_database_sql --env-file $${database_sql}/env/$(ENV).env -f $${database_sql}/docker-compose.yaml down; \
	done
	@$(ECHO_CMD) "$(call log_action_end,Databases SQL stopped successfully)"
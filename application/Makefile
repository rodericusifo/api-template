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
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
    ECHO_CMD = echo
else ifeq ($(UNAME_S),Darwin)
    ECHO_CMD = echo
else
    ECHO_CMD = echo -e
endif

.PHONY: gen-certs
gen-certs:
	@$(ECHO_CMD) "$(call log_action,Generate Certs Services ($(NAME) | $(ENV)))"
	chmod +x "./gen-certs.sh"
	bash ./gen-certs.sh
	@$(ECHO_CMD) "$(call log_action_end,All certificates generated successfully)"

.PHONY: gen-proto
gen-proto:
	@$(ECHO_CMD) "$(call log_action,Generate Proto Services ($(NAME) | $(ENV)))"
	@for application in $(APPLICATIONS); do \
		make -C $${application} gen-proto; \
	done
	@$(ECHO_CMD) "$(call log_action_end,All proto services generated successfully)"

.PHONY: run-application
run-application: gen-certs gen-proto
	@$(ECHO_CMD) "$(call log_action,Start App ($(NAME) | $(ENV)))"
	@for application in $(APPLICATIONS); do \
		$(ECHO_CMD) "$(call log_action,Start App [$${application}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) VERSION=$$(cat $${application}/VERSION) docker compose -p $(NAME)_application --env-file $${application}/env/$(ENV).env -f $${application}/docker-compose.yaml up --build -d; \
	done
	@$(ECHO_CMD) "$(call log_action_end,App started successfully)"

.PHONY: stop-application
stop-application:
	@$(ECHO_CMD) "$(call log_action,Stop App ($(NAME) | $(ENV)))"
	@for application in $(APPLICATIONS); do \
		$(ECHO_CMD) "$(call log_action,Stop App [$${application}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) VERSION=$(cat "$${application}/VERSION") docker compose -p $(NAME)_application --env-file $${application}/env/$(ENV).env -f $${application}/docker-compose.yaml down; \
	done
	@$(ECHO_CMD) "$(call log_action_end,App stopped successfully)"
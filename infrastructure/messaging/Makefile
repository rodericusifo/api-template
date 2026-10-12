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
AVAILABLE_MESSAGING := redpanda lavinmq
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
    ECHO_CMD = echo
else ifeq ($(UNAME_S),Darwin)
    ECHO_CMD = echo
else
    ECHO_CMD = echo -e
endif

.PHONY: check-messaging
check-messaging:
	@$(ECHO_CMD) "$(call log_action,Check App Messaging ($(NAME) | $(ENV)))"
	@for messaging in $(MESSAGING); do \
		if [[ ! " $(AVAILABLE_MESSAGING) " =~ " $${messaging} " ]]; then \
			$(ECHO_CMD) "$(RED)ERROR: Messaging [$${messaging}] is not available. Available Messaging are: $(AVAILABLE_MESSAGING)$(NOCOLOR)"; \
			exit 1; \
		fi; \
	done
	@$(ECHO_CMD) "$(call log_action_end,All Messaging are available)"

.PHONY: run-messaging
run-messaging: check-messaging
	@$(ECHO_CMD) "$(call log_action,Start App Messaging ($(NAME) | $(ENV)))"
	@for messaging in $(MESSAGING); do \
		$(ECHO_CMD) "$(call log_action,Generate App Messaging Volumes [$${messaging}] ($(NAME) | $(ENV)))"; \
		if [ "$${messaging}" == "redpanda" ]; then \
			docker volume ls | grep $(SLUG)_redpanda-0_data_$(ENV) || docker volume create --name $(SLUG)_redpanda-0_data_$(ENV); \
			docker volume ls | grep $(SLUG)_redpanda-1_data_$(ENV) || docker volume create --name $(SLUG)_redpanda-1_data_$(ENV); \
			docker volume ls | grep $(SLUG)_redpanda-2_data_$(ENV) || docker volume create --name $(SLUG)_redpanda-2_data_$(ENV); \
			docker volume ls | grep $(SLUG)_minio_data_$(ENV) || docker volume create --name $(SLUG)_minio_data_$(ENV); \
		elif [ "$${messaging}" == "lavinmq" ]; then \
			docker volume ls | grep $(SLUG)_lavinmq_data_$(ENV) || docker volume create --name $(SLUG)_lavinmq_data_$(ENV); \
			docker volume ls | grep $(SLUG)_etcd_data_$(ENV) || docker volume create --name $(SLUG)_etcd_data_$(ENV); \
    	else \
        	docker volume ls | grep $(SLUG)_$${messaging}_data_$(ENV) || docker volume create --name $(SLUG)_$${messaging}_data_$(ENV); \
    	fi; \
		$(ECHO_CMD) "$(call log_action,Start App Messaging [$${messaging}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) docker compose -p $(NAME)_messaging --env-file $${messaging}/env/$(ENV).env -f $${messaging}/docker-compose.yaml up --build -d; \
	done
	
	@$(ECHO_CMD) "$(call log_action_end,Messaging started successfully)"

.PHONY: stop-messaging
stop-messaging: check-messaging
	@$(ECHO_CMD) "$(call log_action,Stop App Messaging ($(NAME) | $(ENV)))"
	@for messaging in $(MESSAGING); do \
		$(ECHO_CMD) "$(call log_action,Stop App Messaging [$${messaging}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) docker compose -p $(NAME)_messaging --env-file $${messaging}/env/$(ENV).env -f $${messaging}/docker-compose.yaml down; \
	done
	@$(ECHO_CMD) "$(call log_action_end,Messaging stopped successfully)"
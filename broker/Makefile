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
AVAILABLE_BROKERS := redpanda lavinmq
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
    ECHO_CMD = echo
else ifeq ($(UNAME_S),Darwin)
    ECHO_CMD = echo
else
    ECHO_CMD = echo -e
endif

.PHONY: check-broker
check-broker:
	@$(ECHO_CMD) "$(call log_action,Check App Brokers ($(NAME) | $(ENV)))"
	@for broker in $(BROKERS); do \
		if [[ ! " $(AVAILABLE_BROKERS) " =~ " $${broker} " ]]; then \
			$(ECHO_CMD) "$(RED)ERROR: Broker [$${broker}] is not available. Available Brokers are: $(AVAILABLE_BROKERS)$(NOCOLOR)"; \
			exit 1; \
		fi; \
	done
	@$(ECHO_CMD) "$(call log_action_end,All Brokers are available)"

.PHONY: run-broker
run-broker: check-broker
	@$(ECHO_CMD) "$(call log_action,Start App Brokers ($(NAME) | $(ENV)))"
	@for broker in $(BROKERS); do \
		$(ECHO_CMD) "$(call log_action,Generate App Broker Volumes [$${broker}] ($(NAME) | $(ENV)))"; \
		if [ "$${broker}" == "redpanda" ]; then \
			docker volume ls | grep $(SLUG)_redpanda-0_data_$(ENV) || docker volume create --name $(SLUG)_redpanda-0_data_$(ENV); \
			docker volume ls | grep $(SLUG)_redpanda-1_data_$(ENV) || docker volume create --name $(SLUG)_redpanda-1_data_$(ENV); \
			docker volume ls | grep $(SLUG)_redpanda-2_data_$(ENV) || docker volume create --name $(SLUG)_redpanda-2_data_$(ENV); \
			docker volume ls | grep $(SLUG)_minio_data_$(ENV) || docker volume create --name $(SLUG)_minio_data_$(ENV); \
		elif [ "$${broker}" == "lavinmq" ]; then \
			docker volume ls | grep $(SLUG)_lavinmq_data_$(ENV) || docker volume create --name $(SLUG)_lavinmq_data_$(ENV); \
			docker volume ls | grep $(SLUG)_etcd_data_$(ENV) || docker volume create --name $(SLUG)_etcd_data_$(ENV); \
    	else \
        	docker volume ls | grep $(SLUG)_$${broker}_data_$(ENV) || docker volume create --name $(SLUG)_$${broker}_data_$(ENV); \
    	fi; \
		$(ECHO_CMD) "$(call log_action,Start App Broker [$${broker}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) docker compose -p $(NAME)_broker --env-file $${broker}/env/$(ENV).env -f $${broker}/docker-compose.yaml up --build -d; \
	done
	@$(ECHO_CMD) "$(call log_action_end,Brokers started successfully)"

.PHONY: stop-broker
stop-broker: check-broker
	@$(ECHO_CMD) "$(call log_action,Stop App Brokers ($(NAME) | $(ENV)))"
	@for broker in $(BROKERS); do \
		$(ECHO_CMD) "$(call log_action,Stop App Broker [$${broker}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) docker compose -p $(NAME)_broker --env-file $${broker}/env/$(ENV).env -f $${broker}/docker-compose.yaml down; \
	done
	@$(ECHO_CMD) "$(call log_action_end,Brokers stopped successfully)"
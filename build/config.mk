export DISQO_ENV=dev
# *if you change db or app ports match them in config/cluster/go-assessment-dev/config.json

export DISQO_MYSQL_CONTAINER=go-assessment-mysql
export DISQO_MYSQL_IMAGE=$(shell echo $$(git log -n 1 --pretty=format:%H -- ../docker/go-assessment-mysql | cut -c 1-7)-$(DISQO_MYSQL_CONTAINER))
export DISQO_MYSQL_HOST_PORT=1444
export DISQO_MYSQL_USER=root
export DISQO_MYSQL_INTERNAL_PORT=3306
export DISQO_MYSQL_ROOT_PASSWORD=password
export DISQO_MYSQL_DATABASE=badass_db

export DISQO_SERVICE_CONTAINER=go-assessment-app
export DISQO_SERVICE_INTERNAL_PORT=8080
export DISQO_SERVICE_HOST_PORT=8080
export DISQO_SERVICE_IMAGE=$(shell echo $$(git log -n 1 --pretty=format:%H -- ../docker/go-assessment-app | cut -c 1-7)-$(DISQO_SERVICE_CONTAINER))
export DISQO_BUILD_BINARY_DIR=artifacts
export DISQO_SERVICE_BINARY_DIR=go-assessment
export DAEMON=true

# Colors
export DISQO_BLACK=\033[0;30m
export DISQO_RED=\033[0;31m
export DISQO_GREEN=\033[0;32m
export DISQO_BROWN_ORANGE=\033[0;33m
export DISQO_BLUE=\033[0;34m
export DISQO_PURPLE=\033[0;35m
export DISQO_CYAN=\033[0;36m
export DISQO_LIGHT_GRAY=\033[0;37m
export DISQO_DARK_GRAY=\033[1;30m
export DISQO_LIGHT_RED=\033[1;31m
export DISQO_LIGHT_GREEN=\033[1;32m
export DISQO_YELLOW=\033[1;33m
export DISQO_LIGHT_BLUE=\033[1;34m
export DISQO_LIGHT_PURPLE=\033[1;35m
export DISQO_LIGHT_CYAN=\033[1;36m
export DISQO_WHITE=\033[1;37m
export DISQO_NC=\033[0m

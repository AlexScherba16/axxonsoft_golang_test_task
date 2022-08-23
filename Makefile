CONTAINER_NAME=axxonsoft_server

## COLORS
NO_COLOR=\033[0m
CYAN=\033[0;36m
GREEN=\033[32m
UNDER=\033[4m
BOLD=\033[1m

# build: build docker image
build:
	@echo "${CYAN}[build]\t->\tbuilding ${CONTAINER_NAME} executable ${NO_COLOR}"
	docker build -t $(CONTAINER_NAME) .
	@echo "${CYAN}[build]\t<-\tDone!${NO_COLOR}"

# run: run server
run:
	@echo "${CYAN}[run]\t->\trun ${CONTAINER_NAME}${NO_COLOR}"
	@docker run --rm -p 8080:8080 $(CONTAINER_NAME)
	@echo "${CYAN}[run]\t<-\tDone!${NO_COLOR}"

# send_fail_request: send_fail_request to local proxy server
send_fail_request:
	@echo "${CYAN}[send_fail_request]\t->\tlocalhost:8080/proxy/ ${GREEN}${UNDER}${BOLD}"
	@curl -X GET \
	-H "Content-type: application/json" \
	-H "Accept: application/json" \
	-d '{"param0":"pradeep"}' \
	"localhost:8080/proxy/"
	@echo "\n${CYAN}[send_fail_request]\t<-\tDone!${NO_COLOR}"

# ok_google_test: send google request to local proxy server
ok_google_test:
	@echo "${CYAN}[ok_google_test]\t->\tlocalhost:8080/proxy/ ${GREEN}${UNDER}${BOLD}"
	@curl -X GET \
	-H "Content-type: application/json" \
	-H "Accept: application/json" \
	-d '{"method":"GET","url":"http://google.com","headers": {"Authentication": "dccGFzc3dvcmQ="}}' \
	"localhost:8080/proxy/"
	@echo "\n${CYAN}[ok_google_test]\t<-\tDone!${NO_COLOR}"

# ok_httpbin_test: send httpbin/get request to local proxy server
ok_httpbin_test:
	@echo "${CYAN}[ok_httpbin_test]\t->\tlocalhost:8080/proxy/ ${GREEN}${UNDER}${BOLD}"
	@curl -X GET \
	-H "Content-type: application/json" \
	-H "Accept: application/json" \
	-d '{"method":"GET","url":"http://httpbin.org/get", \
	"headers": {\
		"Authentication": "Basic bG9naW46cGFzc3dvcmQ=",\
		"Sec-WebSocket-Extensions": "permessage-deflate; client_max_window_bits"\
		}}' \
	"localhost:8080/proxy/"
	@echo "\n${CYAN}[ok_httpbin_test]\t<-\tDone!${NO_COLOR}"

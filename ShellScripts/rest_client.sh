#!/bin/bash

#######################################
# Send HTTP request with specified parameters.
# Globals:
#   None
# Arguments:
#   $1 @{String} Save response body to this file
#   $2 @{String} HTTP Method
#   $3 @{String} URL without query
#   $4 @{Associative array} Headers
#   $5 @{Associative array} Queries
#   $6 @{String} Request body
# Outputs:
#   Writes executed curl command and status code to stdout.
#   Saves response body in specified ($1) file.
# Returns:
#   0 if request was succeeded, non-zero on error.
#######################################
http_request() {
  declare -r filename=$1
  declare -r method=$2
  declare url=$3
  declare -rn headers_ptr=$4
  declare -rn queries_ptr=$5
  declare -r req_body=$6

  cmd="curl -X ${method} -sS"

  # concatenate queries to URL
  declare -i i=0
  for key in "${!queries_ptr[@]}"; do
    key_value="${key}=${queries_ptr[${key}]}"
    if [ $i -eq 0 ]; then
      delimiter="?"
    else
      delimiter="&"
    fi
    url="${url}${delimiter}${key_value}"
    i=$((i + 1))
  done
  cmd="${cmd} \"${url}\""

  # concatenate headers to cmd
  for key in "${!headers_ptr[@]}"; do
    cmd="${cmd} -H '${key}: ${headers_ptr[${key}]}'"
  done

  # concatenate body to cmd
  if [ -n "${req_body}" ]; then
    cmd="${cmd} -d '${req_body}'"
  fi

  # concatenate others to cmd
  cmd="${cmd} -w '%{http_code}\n'"
  cmd="${cmd} -o ${filename}"
  echo "INFO: send HTTP request with this command: ${cmd}"

  status_code=$(eval "$cmd")

  echo "INFO: status code: ${status_code}"
  if [ "${status_code}" -ge 200 ] && [ "${status_code}" -le 299 ]; then
    return 0
  fi
  return 1
}

#######################################
# [Example] Send GET request
#######################################
example_get() { 
  declare -rA headers=()
  declare -rA queries=()
  declare -r resp_body_filename="example_get_request.json"

  http_request "${resp_body_filename}" "GET" "https://httpbin.org/get" headers queries
  if [ $? -ne 0 ]; then
    exit 1
  fi
  cat ${resp_body_filename} | jq ".url"
}
example_get

#######################################
# [Example] Send POST request
#######################################
example_post() { 
  declare -rA headers=(
    ["content-type"]="application/json"
    ["accept"]="application/json"
  )
  declare -rA queries=(
    ["key1"]="value1"
    ["key2"]="value2"
  )
  declare -r body=$(jq --null-input \
    --arg key1 "value1" \
    --arg key2 "value2" \
    '{
      "key1": $key1,
      "key2": $key2
    }'
  )
  declare -r resp_body_filename="example_post_request.json"

  http_request "${resp_body_filename}" "POST" "https://httpbin.org/post" headers queries "${body}"
  if [ $? -ne 0 ]; then
    exit 1
  fi
  cat ${resp_body_filename} | jq ".url"
}
example_post

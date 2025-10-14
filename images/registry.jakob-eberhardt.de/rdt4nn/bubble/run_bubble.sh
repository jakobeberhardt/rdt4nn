#!/bin/bash

# Bubble configuration loader script
# This script loads configuration from bubble.conf file and exports environment variables

CONFIG_FILE="bubble.conf"

if [ -f "$CONFIG_FILE" ]; then
    echo "Loading configuration from $CONFIG_FILE"
    
    # Read configuration file and export environment variables
    while IFS='=' read -r key value; do
        # Skip comments and empty lines
        if [[ ! $key =~ ^[[:space:]]*# && -n $key ]]; then
            # Remove any trailing comments
            value=$(echo "$value" | cut -d'#' -f1 | sed 's/[[:space:]]*$//')
            # Export the variable
            export "$key"="$value"
            echo "Set $key=$value"
        fi
    done < "$CONFIG_FILE"
else
    echo "Configuration file $CONFIG_FILE not found, using defaults"
fi

# Run the bubble program with any passed arguments
exec ./bubble "$@"
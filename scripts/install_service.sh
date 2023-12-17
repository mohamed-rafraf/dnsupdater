#!/bin/bash

# Define variables for download
BINARY_URL="https://github.com/mohamed-rafraf/dnsupdater/releases/download/v1.0.0/dnsupdater"  # Replace with the actual binary URL
BINARY_NAME="dnsupdater"
BINARY_DESTINATION="/usr/local/bin/$BINARY_NAME"
SERVICE_NAME="$BINARY_NAME"
SERVICE_FILE="/etc/systemd/system/$SERVICE_NAME.service"

# Default values
DEFAULT_DOMAIN="securinets.tn"

# Ask user for environment variables with defaults
read -p "Input the domain [$DEFAULT_DOMAIN]: " DOMAIN
DOMAIN=${DOMAIN:-$DEFAULT_DOMAIN}

DEFAULT_SUBDOMAIN="moataz"

read -p "Input the sub-domain [$DEFAULT_SUBDOMAIN] : " SUBDOMAIN
SUBDOMAIN=${SUBDOMAIN:-$DEFAULT_SUBDOMAIN}

DEFAULT_CHECK_INTERVAL="2s"

read -p "Input the interval (e.g., 10m, 1h) [$DEFAULT_CHECK_INTERVAL]: " CHECK_INTERVAL
CHECK_INTERVAL=${CHECK_INTERVAL:-$DEFAULT_CHECK_INTERVAL}

read -p "Input the api key: " API_KEY

read -p "Input the email: " EMAIL

DEFAULT_FILE_PATH="/var/lib/dnsupdater"
read -p "Input the directory that contains dnsupdater data [$DEFAULT_FILE_PATH]" FILE_PATH
FILE_PATH=${FILE_PATH:-$DEFAULT_FILE_PATH}

FILE_PATH=$FILE_PATH/ip.dat

echo $FILE_PATH

# Check if the script is running as root
if [ "$(id -u)" -ne 0 ]; then
    echo "This script must be run as root."
    exit 1
fi

# Download the binary
echo "Downloading binary from $BINARY_URL..."
wget -O "$BINARY_DESTINATION" "$BINARY_URL"
chmod +x "$BINARY_DESTINATION"

# Create the systemd service file
echo "Creating systemd service file at $SERVICE_FILE..."
cat > "$SERVICE_FILE" <<EOF
[Unit]
Description="DNS Updater for updating DNS records in CloudFlare"
After=network.target

[Service]
ExecStart=$BINARY_DESTINATION
Restart=on-failure
Environment="DOMAIN=$DOMAIN"
Environment="SUBDOMAIN=$SUBDOMAIN"
Environment="CHECK_INTERVAL=$CHECK_INTERVAL"
Environment="API_KEY=$API_KEY"
Environment="EMAIL=$EMAIL"
Environment="FILE_PATH=$FILE_PATH"
# Add additional Environment="KEY=value" lines if needed

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd to apply new changes
echo "Reloading systemd daemon..."
systemctl daemon-reload

# Enable and start the service
echo "Enabling and starting $SERVICE_NAME service..."
systemctl enable "$SERVICE_NAME"
systemctl start "$SERVICE_NAME"

echo "Installation completed successfully."
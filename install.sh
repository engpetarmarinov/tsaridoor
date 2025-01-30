#!/bin/bash
set -e
SERVICE_NAME="tsaridoor.service"
SYSTEMD_PATH="/etc/systemd/system/"
TEMP_PATH="/home/$TSARIDOOR_USERNAME"

# Copy the service file
echo "Copying $SERVICE_NAME to $TSARIDOOR_IP..."
scp "$SERVICE_NAME" "$TSARIDOOR_USERNAME@$TSARIDOOR_IP:$TEMP_PATH"

# SSH to set permissions, enable, and start the service
echo "Setting up systemd service on $TSARIDOOR_IP..."
ssh "$TSARIDOOR_USERNAME@$TSARIDOOR_IP" <<EOF
  sudo mv $TEMP_PATH/$SERVICE_NAME $SYSTEMD_PATH
  sudo systemctl daemon-reload
  sudo systemctl enable $SERVICE_NAME
EOF

echo "Service deployment complete!"

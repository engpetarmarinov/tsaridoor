#!/bin/bash
set -e
SERVICE_NAME="tsaridoor.service"
echo "Copying files..."
scp -r users.csv ./tsaridoor-hotrelease ./home.html ./static  "${TSARIDOOR_USERNAME}"@"${TSARIDOOR_IP}":/home/"${TSARIDOOR_USERNAME}"/tsaridoor

echo "Hotreleasing: Copy the new version and restart ${SERVICE_NAME}..."
ssh "${TSARIDOOR_USERNAME}"@"${TSARIDOOR_IP}"  <<EOF
  mkdir -p /home/$TSARIDOOR_USERNAME/tsaridoor/
  cd /home/$TSARIDOOR_USERNAME/tsaridoor/
  sudo systemctl stop $SERVICE_NAME
  mv tsaridoor-hotrelease tsaridoor
  chmod +x tsaridoor
  sudo systemctl start $SERVICE_NAME
EOF

echo "Done. Press any key to exit."
read -r
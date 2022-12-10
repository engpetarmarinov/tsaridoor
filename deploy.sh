#!/bin/bash
echo "Copying files..."
scp -r users.csv ./tsaridoor-hotrelease ./home.html ./static  "${TSARIDOOR_USERNAME}"@"${TSARIDOOR_IP}":/home/"${TSARIDOOR_USERNAME}"/tsaridoor

echo "Hotreleasing: Copy the new version and restart tsaridoor.service..."
ssh "${TSARIDOOR_USERNAME}"@"${TSARIDOOR_IP}" "cd /home/${TSARIDOOR_USERNAME}/tsaridoor/ && sudo systemctl stop tsaridoor.service && mv tsaridoor-hotrelease tsaridoor && chmod +x tsaridoor && sudo systemctl start tsaridoor.service"
echo "Done. Press any key to exit."
read -r

# tsaridoor
A simple go web app that unlocks my door with a Raspberry Pi and a relay

Create csv file with usernames and sha256 passwords, e.g. for test/test:
```bash
cat <<EOF > middleware/users.csv
username,password
test,9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08
EOF
```

# Build the binary
```bash
./build.sh
```

# Install systemd service unit
```bash
TSARIDOOR_USERNAME={username} TSARIDOOR_IP={ip} ./install.sh
```

# Deploy
```bash
TSARIDOOR_USERNAME={username} TSARIDOOR_IP={ip} ./deploy.sh
```

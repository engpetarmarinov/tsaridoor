# tsaridoor
A simple go web app that unlocks my door with a Raspberry Pi and a relay

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

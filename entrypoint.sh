#!/bin/bash
cat << EOF > .env
PORT=3000
IPV4=$IPV4
IPV6=$IPV6
REFRESH=$REFRESH
LASTHEARD=$LASTHEARD
MREFDFILE=/var/log/mrefd.xml
EMAIL=$EMAIL
EOF

./gomrefdash
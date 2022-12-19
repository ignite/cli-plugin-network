set -x
h=$(cat cli_last_sync)
git --git-dir=../cli/.git diff ${h} HEAD -- ignite/services/network | patch -p3
git --git-dir=../cli/.git diff ${h} HEAD -- ignite/cmd/network* | patch -p2
git --git-dir=../cli/.git rev-parse HEAD > cli_last_sync

## Docker Network:

### Remove the old network (optional):

```
docker network rm gate
```
⚠️ Only do this if no other running containers depend on it.

### Create it again with IPAM config:

```
docker network create \
--driver=bridge \
--subnet=172.20.0.0/16 \
--gateway=172.20.0.1 \
gate
```
You can also add:

* --ip-range=... if you want to restrict dynamic IPs.
* --attachable if needed for connecting to it manually.

#### ✅ **You can use `init_network.sh` in the scripts to run above commands.**
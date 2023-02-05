# Instana

## Start Instana Agent

```sh
# Launch and access multipass VM instance.
multipass launch -n instana && multipass shell instana

# Install instana-agent
curl -o setup_agent.sh https://setup.instana.io/agent \
     && chmod 700 ./setup_agent.sh \
     && sudo ./setup_agent.sh -a someaccesstoken -t dynamic \
     -e ingress-green-saas.instana.io:443 -s

# Check instana-agent status
sudo systemctl status instana-agent
...
     Status: "Agent is available"

sudo lsof -i -P -n | grep LISTEN
...
java      4144            root  125u  IPv6  37386      0t0  TCP 127.0.0.1:42699 (LISTEN)
```

## Start Sample App:

```sh
cd cmd/instana
go build

# copy go binary to vm instance
multipass transfer -vvvv  instana instana:instana

# get shell on the vm
multipass shell instana

# make binary executable
chmod 0700 instana

export INSTANA_AGENT_ENDPOINT=127.0.0.1
export INSTANA_AUTO_PROFILE=true
export INSTANA_DEBUG=true
# overwrite service.name
export INSTANA_SERVICE_NAME=fib
./run.sh
# Your server is live!
# Try to navigate to: http://127.0.0.1:3000/fib?n=6
```

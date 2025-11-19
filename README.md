# Edge Device Simulator: IoT simulator for task offload analysis and experiments on Cloud and Edge
Simulates a device on the edge of a 5G network that offloads a task or executes it locally. The simulator was conceived to be used
on O-RAN 5G networks but can be expanded to different scenarios.

## Usage
The simulator can be run locally or in a docker container. The docker image can be built using the `Dockerfile` in the root directory.

## Configuration
The simulator can be configured via command line parameters. The available parameters are:
```
./edge-sim --help                                     
Usage of ./edge-sim:
  -arrival-rate float
    	Arrival rate of workloads in requests per second. (default 2)
  -callback string
    	Callback URL to send results to. (default "http://localhost:8080")
  -duration duration
    	Time in seconds to run the simulation. (default 1m0s)
  -scenario int
    	Scenario to run:
    	0 - Local processing
    	1 - Cloud processing
    	2 - Hybrid edge with xApp (default 2)
  -workload-mean int
    	Mean of workload sizes. (default 25)
  -workload-std-var int
    	Standard deviation of workload sizes. (default 3)

```

Sensible defaults are provided for all parameters, but they can be changed as needed.

For correct RNG, the environment variables `BASE_SEED` and `DEVICE_ID` must be set. These are set via docker-compose
but can also be set manually.

## Randomness
The simulator uses a pseudo-random number generator (PRNG) to generate the task payload. For reproducibility, the source 
is created using the [PCG](https://pkg.go.dev/math/rand/v2#NewPCG) method, which uses a base seed and the device ID.

For more information on the PRNG, see the [PCG website](https://www.pcg-random.org).

The task generation rate is modeled as an exponential distribution, where the rate can be configured via the parameter `--arrival-rate`.

The workload size is modeled as a log-normal distribution, where the mean and standard deviation can be configured via the parameters `--workload-mean` and `--workload-std-var`.

## Scenarios
Below are described the scenarios supported by the simulator. It is important to note that the tasks
are simply started (be it locally, on the cloud or on the edge), and the results are received by a callback server,
configured via the parameter `--callback-server`. For O-RAN simulations, the server [xApp Callback Server](https://github.com/RafaelRochaS/xapp-task-callback-server) 
can be used.


### 0: Local Processing
The first scenario is the simplest one, where the device executes the task locally. It will incur
the least amount of distance latency, at the cost of lower processing power (and higher energy consumption).

The scenario is activated by using the parameter `--scenario 0`. Alternatively, the docker compose file `scenario-0-compose.yml`
can be used to start the devices with this scenario.

### 1: Cloud Processing
In this scenario, the device offloads the task to the cloud, where it is executed. All tasks are
sent to the cloud, with no intelligence in the offloading. The idea is to execute the tasks with more processing power,
at the cost of increased latency.

The scenario is activated by using the parameter `--scenario 1`. Alternatively, the docker compose file `scenario-1-compose.yml`
can be used to start the devices with this scenario.

**Note:** Scenario 1 instantiates a pod in a Kubernetes cluster, and as such requires a kubeconfig file to be present, pointing
to a valid cluster. The kubeconfig file can be created by running e.g. `gcloud container clusters get-credentials <cluster-name> --zone <zone>`, in a [GCP GKE cluster](https://cloud.google.com/kubernetes-engine?hl=en).
Since such files are credentials, they are not included in this repository.

### 2: Hybrid Processing
The last (current) scenario is the most intelligent one, where the device checks the size of the task and offloads it to the cloud
if it is larger than a certain threshold. 

However, instead of simply pushing the task to the cloud, a request is made to a
MEC server controller. In the O-RAN scenario, an [xApp](https://github.com/RafaelRochaS/xapp-mec-go) located in the NearRT RIC is used, which
will then decide whether to execute the task on the edge server (based on available resources) or on the cloud.

The idea is to balance the execution of the task to better utilize local, edge and cloud resources, resulting in optimal performance
and latency.

The scenario is activated by using the parameter `--scenario 2`. Alternatively, the docker compose file `scenario-2-compose.yml`
can be used to start the devices with this scenario.
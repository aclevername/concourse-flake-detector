# concourse-flake-detector

A tool to find flakey jobs in a Concourse pipeline. Flakey in this context is when a build fails when it has otherwise passed with the same set of inputs. This is currently restricted to only checking changes to git resources.

In order to build the flake-detect checkout the repo and run `cd cmd/flake-detector && go build .`
## Usage
##### Flags
`-url <concourse url>` **_Required_** - The URL of the concourse.

`-pipeline <pipeline name>`  **_Required_** - The name of the pipeline to be scanned.

`-team <team name>` Optional- Needed if your pipeline belongs to a concourse team.

`-count <number of builds to scan>` Optional- Only scan the last x amount of builds. Currently max's and defaults at 100.

`-bearer <bearer authentication token>` Optional- Needed if your pipeline is not public. To obtain a bearer token go to `<ci-url>/api/v1/teams/<team>/auth/token`or run the `flake-detector` without the flag configured and it will prompt you with the URL to go to obtain a token.

`-debug` Optional- Configure true if you want debug information about what endpoints are being hit.

`-insecure-tls` Optional- TLS accepts any certificate presented by the server and any host name in that certificate. Not recommended.



 ##### Examples

`./flake-detector -team main -pipeline main -url https://ci.concourse.ci`

```
+------------------------+--------+--------+
|          NAME          | BUILDS | FLAKES |
+------------------------+--------+--------+
| fly                    |     74 |      0 |
| atc                    |     74 |      9 |
| baggageclaim           |     74 |      0 |
| blackbox               |     74 |      2 |
| bosh-testflight        |     74 |      1 |
| bin-testflight         |     74 |      0 |
| bosh-deploy            |     74 |      2 |
| bin-smoke              |     74 |      1 |
| bin-docker             |     74 |      1 |
| shipit                 |      8 |      0 |
| virtualbox-box         |     13 |      2 |
| virtualbox-testflight  |     11 |      1 |
| release-virtualbox-box |      7 |      2 |
| topgun                 |     74 |      7 |
+------------------------+--------+--------+
```

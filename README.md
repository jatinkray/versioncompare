[![Go](https://github.com/jatinkray/versioncompare/actions/workflows/go.yml/badge.svg)](https://github.com/jatinkray/versioncompare/actions/workflows/go.yml)
# Description
This app is developed using golang to build a cron utility for comparing deployed application url core versions.


## local build and run 

For local development, your system required to have following prequisites
- golang minimum version 12

```shell
./run.sh
```

## Production run 
Using executable binary

```shell
wget 

export APP_CRON_SCHEDULE="0 */1 * * * *"
export TEST_URL=""
export PRODUCTION_URL=""
versioncompare
```

Using docker 

```shell
docker run -e APP_CRON_SCHEDULE="0 */1 * * * *" -e TEST_URL="" -e PRODUCTION_URL="" -it jatinkray/versioncompare
```

### CRON schedule
The cron format used by the Quartz Scheduler, commonly used for scheduled jobs


### CRON Expression Format
A cron expression represents a set of times, using 6 space-separated fields.

Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes         | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
Note: Month and Day-of-week field values are case insensitive. "SUN", "Sun", and "sun" are equally accepted.
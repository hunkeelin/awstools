### Introduction ###
Quick reboot tools and instance info mapping tools I wrote for practice. 


### Usage ###
I've included the binary in the repo. It should work on any linux-base systems.

```
./awstools -h
age of ./awstools:
  -config string
        location of the config file (default "config")
  -nametag string
        the nametag of the instance you want
  -nodry
        by default everything is dryrun so you don't trip yourself =) (default true)
  -region string
        the region you are running on (default "us-west-2")
./awstools -nametag foo.somecompany.com
The ip address : 10.123.1.21
rebooting instance:  i-0358OIdjflkajc411e
DryRunOperation: Request would have succeeded, but DryRun flag is set.
        status code: 412, request id: aslkdfj-6asdf-4fbaasdf9a81-asdfsadf371a
```


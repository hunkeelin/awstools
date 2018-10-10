### Introduction ###
Quick reboot tools and instance info mapping tools I wrote for practice. 


### Usage ###
I've included the binary in the repo. It should work on any linux-base systems.

```
Usage of ./awstools:
  -bindaddr string
        the bind address for server option
  -bindport string
        the bind port for the server option (default "2018")
  -config string
        location of the config file (default "config")
  -dry
        Do dryrun
  -nametag string
        the nametag of the instance you want
  -region string
        the region you are running on (default "us-west-2")
  -server
        whether to host an api server
```

#### cli ####
```
./awstools -nametag foo.somecompany.com
The ip address : 10.123.1.21
rebooting instance:  i-0358OIdjflkajc411e
DryRunOperation: Request would have succeeded, but DryRun flag is set.
        status code: 412, request id: aslkdfj-6asdf-4fbaasdf9a81-asdfsadf371a
```

#### server ####
There's a `-server` option for this software where it will host a http server to recieve apicalls. the payload is as follows:
```
{
    "action":$action // for now only "reboot" is allowed
    "nametag":$nametag // basically the nametag of the host you want to reboot
}
```

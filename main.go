package main
import (
    "flag"
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

var (
    config = flag.String("config","config","location of the config file")
    region = flag.String("region","us-west-2","the region you are running on")
    instancetag = flag.String("nametag","","the nametag of the instance you want")
    dryrun = flag.Bool("dry",false,"Do dryrun")
)
func main() {
    flag.Parse()
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
        Config: aws.Config {
            Region: aws.String(*region),
        },
    }))
    ec2Svc := ec2.New(sess)
    ip,id,err := getIpInstancefromTag(*instancetag,ec2Svc)
    if err != nil {
        panic(err)
    }
    fmt.Println("The Private ip address for natting :", ip)
    fmt.Println("rebooting instance: ",id)
    err = rebootIn(ec2Svc,*dryrun,id)
    if err != nil {
        panic(err)
    }
}

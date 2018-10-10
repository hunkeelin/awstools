package main
import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/aws/session"
)

func reboot(nametag string) error{
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
        Config: aws.Config {
            Region: aws.String(*region),
        },
    }))
    ec2Svc := ec2.New(sess)
    _,id,err := getIpInstancefromTag(nametag,ec2Svc)
    if err != nil {
        fmt.Println("unable to get instance "+nametag)
        return err
    }
    input := &ec2.RebootInstancesInput{
        InstanceIds: []*string{
            aws.String(id),
        },
        DryRun: aws.Bool(false),
    }
    _, err = ec2Svc.RebootInstances(input)
    if err != nil {
        return err
    }
    return nil
}

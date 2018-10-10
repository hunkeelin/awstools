package main
import (
    "fmt"
    "sync"
    "errors"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func getIpInstancefromTag(tag string,svc *ec2.EC2) (ip string,id string,e error) {
    var toreturnip,toreturnid string
    result, err := svc.DescribeInstances(nil)
    if err != nil {
        return toreturnip,toreturnid,err
    } else {
        sema := make(chan struct{},4)
        wg := sync.WaitGroup{}
        for _,i := range result.Reservations {
            for _,m := range i.Instances {
                sema <- struct{}{}
                wg.Add(1)
                go func(){
                    for _,j := range m.Tags {
                        check := *j.Key
                        if check == "Name" {
                            nametag := *j.Value
                            if nametag == tag {
                                Ip := *m.PrivateIpAddress
                                Id := *m.InstanceId
                                toreturnip = Ip
                                toreturnid = Id
                            }
                        }
                    }
                    <-sema
                    wg.Done()
                }()
            }
            wg.Wait()
        }
    }
    if toreturnip == "" && toreturnid == "" {
        return toreturnip,toreturnid,errors.New("Tag is invalid")
    }
    return toreturnip,toreturnid,nil
}
func rebootIn(svc *ec2.EC2,dry bool,id string) error{
    input := &ec2.RebootInstancesInput{
        InstanceIds: []*string{
            aws.String(id),
        },
        DryRun: aws.Bool(dry),
    }
    _, err := svc.RebootInstances(input)
    awsErr, ok := err.(awserr.Error)
    if ok && awsErr.Code() == "DryRunOperation"{
        fmt.Println(awsErr)
        return nil
    }
    if err != nil {
        return err
    }
    return nil
}

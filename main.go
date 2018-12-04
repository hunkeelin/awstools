package main
import (
    "flag"
    "fmt"
    "net/http"
    "github.com/hunkeelin/mtls/klinserver"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

var (
    config = flag.String("config","config","location of the config file")
    region = flag.String("region","us-west-2","the region you are running on")
    instancetag = flag.String("nametag","","the nametag of the instance you want")
    dryrun = flag.Bool("dry",false,"Do dryrun")
    server = flag.Bool("server",false,"whether to host an api server")
    addr = flag.String("bindaddr","","the bind address for server option")
    port = flag.String("bindport","2018","the bind port for the server option")
    rebootvm = flag.Bool("reboot",false,"whether to reboot the box")
)
func main() {
    flag.Parse()
    if *server {
        con := http.NewServeMux()
        con.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            mainHandler(w, r)
        })
        j := &klinserver.ServerConfig {
            BindAddr: *addr,
            BindPort: *port,
            ServeMux: con,
        }
        panic(klinserver.Server(j))
    } else {
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
        if *rebootvm {
            err = rebootIn(ec2Svc,*dryrun,id)
            if err != nil {
                panic(err)
            }
        }
    }
}

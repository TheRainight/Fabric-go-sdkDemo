package main

import (
    "os"
    "fmt"
    "github.com/wangz/sdkInit"
    "github.com/wangz/web"
    "github.com/wangz/service"
    "github.com/wangz/web/controller"


   
)

const (
    configFile = "config.yaml"
    initialized = false
    SimpleCC = "simplecc"
)

func main() {

    initInfo := &sdkInit.InitInfo{

        ChannelID: "wzchannel",
        ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/wangz/fixtures/artifacts/channel.tx",

        OrgAdmin:"Admin",
        OrgName:"FBI",
        OrdererOrgName: "orderer.wangz.com",

        ChaincodeID: SimpleCC,
        ChaincodeGoPath: os.Getenv("GOPATH"),
        ChaincodePath: "github.com/wangz/chaincode/",
        UserName:"User1",
    }

    sdk, err := sdkInit.SetupSDK(configFile, initialized)
    if err != nil {
        fmt.Printf(err.Error())
        return
    }

    defer sdk.Close()

    err = sdkInit.CreateChannel(sdk, initInfo)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
     serviceSetup := service.ServiceSetup{
        ChaincodeID:SimpleCC,
        Client:channelClient,
    }
    fmt.Println(channelClient)
    app := controller.Application{
        Fabric: &serviceSetup,
    }
    web.WebStart(&app)
   

    // msg, err := serviceSetup.SetInfo("AAA", "200")
    // if err != nil {
    //     fmt.Println(err)
    // } else {
    //     fmt.Println(msg)
    // }



}
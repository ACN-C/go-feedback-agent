## Feedback Agent
## Features

- CPU metric
- RAM metric
- TCP connections metric
- Read/reload from config
- Halt/Down/Normal Status States

## Prerequisites

* Go v1.9 or later
* Windows

## Build (powershell) 

```
$env:GOOS="windows"
$env:GOARCH="amd64"
go build -ldflags "-X main.Version=$env:VERSION -X main.Build=$env:BUILD" -v -o ./bin/windows64/LBCPUMon.exe ./src
```
 
## XML

```
<xml>
  <Cpu>
    <ImportanceFactor value="1" />
    <ThresholdValue value="100" />
  </Cpu>
  <Ram>
    <ImportanceFactor value="0" />
    <ThresholdValue value="100" />
  </Ram>
  <TCPService>
    <Name value="HTTP" />
    <IPAddress value="*" />
    <Port value="80" />
    <MaxConnections value="0" />
    <ImportanceFactor value="0" />
  </TCPService>
  <ReadAgentStatusFromConfig value="False" />
  <ReadAgentStatusFromConfigInterval value="5" />
  <AgentStatus value="Normal" />
  <Interval value="10" />
  <Port value="3333" />
  <ReturnIdle value="true" />
  <LogLevel value="INFO" />
</xml>
```

## Install
```
sc create FeedBackService binPath= "full\path\to\LBCPUMon.exe"
```
 


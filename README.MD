# grafana-sync

Keep your grafana dashboards in sync.

## Table of Contents

- [grafana-sync](#grafana-sync)
  - [Table of Contents](#table-of-contents)
  - [Installing](#installing)
  - [Getting Started](#getting-started)
    - [Pull dashboards](#pull-dashboards)
    - [Pull folder](#pull-folder)
    - [Pull notifications](#pull-notifications)
    - [Pull datasources](#pull-datasources)
    - [Push dashboards](#push-dashboards)
    - [Push folders](#push-folders)
    - [Push notifications](#push-notifications)
    - [Push datasources](#push-datasources)
  - [Global parameters](#global-parameters)
  - [Contributing](#contributing)
  - [License](#license)

## Installing

Download the latest binary from [releases](https://github.com/mpostument/grafana-sync/releases)

## Getting Started

### Pull dashboards

```shell
# Save all dashboards to directory
grafana-sync pull-dashboards --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="dashboards" --url http://127.0.0.1:3000

# Save all dashboards from specific folder to local directory using folder name
grafana-sync pull-dashboards --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="dashboards" --url http://127.0.0.1:3000 --folderName="TestFolder"

# Save all dashboards from specific folder to local directory using folder id
grafana-sync pull-dashboards --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="dashboards" --url http://127.0.0.1:3000 --folderId=1

# Save dashboards with specific tags to directory
grafana-sync pull-dashboards --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="dashboards" --url http://127.0.0.1:3000 --tag=export
```

### Pull folder

```shell
grafana-sync pull-folders --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="folders" --url http://127.0.0.1:3000
```

### Pull notifications

```shell
grafana-sync pull-notifications --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="notifications" --url http://127.0.0.1:3000
```

### Pull datasources

```shell
grafana-sync pull-datasources --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="datasources" --url http://127.0.0.1:3000
```

### Push dashboards

```shell
# Push all dashboards to general directory
grafana-sync push-dashboards --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="dashboards" --url http://127.0.0.1:3000

# Push dashboards to grafana in custom folder by folder name
grafana-sync push-dashboards --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="dashboards" --url http://127.0.0.1:3000 --folderName="TestFolder"

# Push folders to grafana in custom folder by folder id
grafana-sync push-folders --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="dashboards" --url http://127.0.0.1:3000 --folderId=1
```

### Push folders

```shell
grafana-sync push-folders --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="folders" --url http://127.0.0.1:3000
```

### Push notifications

```shell
grafana-sync push-notifications --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="notifications" --url http://127.0.0.1:3000
```

### Push datasources

```shell
grafana-sync push-datasources --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="datasources" --url http://127.0.0.1:3000
```

## Global parameters

`directory` - Directory where to save dashboards. Default `.`  
`tag` - Dashboard tag to read. Supported only with `pull` option. Default `""`  
`apikey` - Grafana api key, need to be editor or admin. Default `""`.  
Api key can be stored in `$HOME/.grafana-sync.yaml` as `apikey: <ApiKey>`  
`url` - Grafana Url with port. Default `http://localhost:3000`  

## Contributing

1. Fork it
2. Download your fork to your PC ( `git clone https://github.com/your_username/grafana-sync && cd grafana-sync` )
3. Create your feature branch ( `git checkout -b my-new-feature` )
4. Make changes and add them ( `git add .` )
5. Commit your changes ( `git commit -m 'Add some feature'` )
6. Push to the branch ( `git push origin my-new-feature` )
7. Create new pull request

## License

grafana-sync is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/mpostument/grafana-sync/blob/master/LICENSE)

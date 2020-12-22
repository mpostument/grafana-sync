# grafana-sync

Keep your grafana dashboards in sync.

## Table of Contents

- [grafana-sync](#grafana-sync)
  - [Table of Contents](#table-of-contents)
  - [Installing](#installing)
  - [Getting Started](#getting-started)
    - [Pull](#pull)
      - [Save all dashboards to directory](#save-all-dashboards-to-directory)
      - [Save dashboards with specific tags to directory](#save-dashboards-with-specific-tags-to-directory)
    - [Push](#push)
      - [Push dashboards to grafana](#push-dashboards-to-grafana)
  - [Global parameters](#global-parameters)
  - [Contributing](#contributing)
  - [License](#license)

## Installing

Download the latest binary from [releases](https://github.com/mpostument/grafana-sync/releases)

## Getting Started

### Pull

#### Save all dashboards to directory

`grafana-sync pull`  
Example:
 `grafana-sync pull --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="dashboard" --url http://127.0.0.1:3000`

#### Save dashboards with specific tags to directory

`grafana-sync pull`  
Example:
 `grafana-sync pull --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="dashboard" --url http://127.0.0.1:3000 --tag=export`  

### Push

#### Push dashboards to grafana

`grafana-sync push`  
Example:
 `grafana-sync push --apikey="eyJrIjoiOWJYTktGNFlCbFVMOG1LY3d6ekN4Mmw4MFgyYU44a1UiLCJuIjoiY29icmEiLCJpZCI6MX0=" --directory="dashboard" --url http://127.0.0.1:3000`

## Global parameters

`directory` - Directory where to save dashboards. Default `.`  
`tag` - Dashboard tag to read. Supported only with `pull` option. Default `""`  
`apikey` - Grafana api key, need to be editor or admin. Default `""`. Api key can be stored in `$HOME/.grafana-sync.yaml` as `apikey: <ApiKey>`
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

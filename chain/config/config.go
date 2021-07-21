package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type Config tomlConfig

type tomlConfig struct {
	Title string
	Currency currency
	Chain chainConfig
	System systemConfig
	Icb icbConfig
	General generalConfig
	P2p p2pConfig
	Database databaseConfig
	Validators validatorConfig
}

type currency struct {
	Name string
	Symbol string
	NamePlur string `toml:"name_plur"`
	SymbolPlur string `toml:"symbol_plur"`
	MinUnitName string `toml:"min_unit_name"`
	MinValue float32 `toml:"min_value"`
	RegaliosInRegal int `toml:"regalios_in_regal"`
	Decimal int
	CoinbaseAddress string `toml:"coinbase_address"`
	LiquidityAddress string `toml:"liquidity_address"`
}

type chainConfig struct {
	Version int
	GenesisFile string `toml:"genesis_file"`
	GenerateNewKeys bool `toml:"generate_new_keys"`
	BlocksPerHour int `toml:"blocks_per_hour"`
	TxFeesPercent float32 `toml:"tx_fees_percent"`
	MaxTxFee float32 `toml:"max_tx_fee"`
	MinTxFee float32 `toml:"min_tx_fee"`
	MaxBlockSize int `toml:"max_block_size"`

}

type systemConfig struct {
	ICBEnabled bool `toml:"icb_enabled"`
	DebugMode bool `toml:"debug_enabled"`
	DevMode bool `toml:"dev_mode"`
	TestMode bool `toml:"test_mode"`
	ProdMode bool `toml:"prod_mode"`
	NetworkType string `toml:"network_type"`
	MaxCpus int `toml:"max_cpus"`
	MinCpus int `toml:"min_cpus"`
	MaxMem int `toml:"max_mem"`
	MinMem int `toml:"min_mem"`
	RPCPort int `toml:"rpc_port"`
	PrivKey string `toml:"priv_key"`
	PubKey string `toml:"pub_key"`
	LogPath string `toml:"log_path"`
	LogFile string `toml:"log_file"`
	LogLevel string `toml:"log_level"`
	DisplayError bool `toml:"display_error"`
}

type icbConfig struct {
	TotalSupply uint32 `toml:"total_supply"`
	TokenName string `toml:"token_name"`
	EthTokenAddress string `toml:"eth_token_address"`
	TokenMinBid float32 `toml:"token_min_bid"`
	Website string
	Email string
	Twitter string
	Telegram string
	Coinmarketcap string
	Coingecko string
	Bitcointalk string
	Reddit string
	WhitepaperURL string `toml:"whitepaper_url"`
	Github string
	Objective int
	StartDate string `toml:"start_date"`
	MaxDate string `toml:"max_date"`
}

type generalConfig struct {
	Domain string
	Name string
	Email string
	Version string
}

type p2pConfig struct {
	BootstrapFile string `toml:"bootstrap_file"`
	CustomBootstrappers []string `toml:"custom_bootstrappers"`
	Port int
	MinPeers int `toml:"min_peers"`
	MaxPeers int `toml:"max_peers"`
	Timeout int
}


type databaseConfig struct {
	InMemoryMax int `toml:"in_memory_max"`
	Localpath string
	Testpath string
	Livepath string
	Dbnames []string
}

type validatorConfig struct {
	MinValidatorsRequired int `toml:"min_validators_required"`
	MaxValidators int `toml:"max_validators"`
	MinStake float32 `toml:"min_stake"`
	VotingPowerPerStakeUnit float32 `toml:"voting_power_per_stake_unit"`
	MaxVotingPower float32 `toml:"max_voting_power"`
	MaxStake int64 `toml:"max_stake"`
	MaxSubValidators int `toml:"max_sub_validators"`
	MaxTimeToValidate int64 `toml:"max_time_to_validate"`
	StakeReward int64 `toml:"stake_reward"`
	RequiredForQuorumPercent int `toml:"required_for_quorum_percent"`
}

func (c Config) Parse() *tomlConfig {
	var config tomlConfig
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		logrus.Errorln("an error occured while loading regalio's configuration file: ", err)
		return nil
	}

	return &config

}

var ChainConfig *tomlConfig

func GetConfiguration() error {

	var c Config
	tConf := c.Parse()
	if tConf == nil {
		return errors.New("cannot load configuration file")
	}
	ChainConfig = tConf
	return nil
}

func init() {
	err := GetConfiguration()
	if err != nil {
		panic(err)
	}

}
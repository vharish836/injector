package bridge

type GetInfoReq struct{}

type GetInfoRsp struct {
	Version         string  `json:"version"`
	Nodeversion     int     `json:"nodeversion"`
	Protocolversion int     `json:"protocolversion"`
	Chainname       string  `json:"chainname"`
	Description     string  `json:"description"`
	Protocol        string  `json:"protocol"`
	Port            int     `json:"port"`
	Setupblocks     int     `json:"setupblocks"`
	Nodeaddress     string  `json:"nodeaddress"`
	Burnaddress     string  `json:"burnaddress"`
	Incomingpaused  bool    `json:"incomingpaused"`
	Miningpaused    bool    `json:"miningpaused"`
	Walletversion   int     `json:"walletversion"`
	Balance         float64 `json:"balance"`
	Walletdbversion int     `json:"walletdbversion"`
	Reindex         bool    `json:"reindex"`
	Blocks          int     `json:"blocks"`
	Timeoffset      int     `json:"timeoffset"`
	Connections     int     `json:"connections"`
	Proxy           string  `json:"proxy"`
	Difficulty      float64 `json:"difficulty"`
	Testnet         bool    `json:"testnet"`
	Keypoololdest   int     `json:"keypoololdest"`
	Keypoolsize     int     `json:"keypoolsize"`
	Paytxfee        float64 `json:"paytxfee"`
	Relayfee        float64 `json:"relayfee"`
	Errors          string  `json:"errors"`
}

type GetBlockChainParamsReq struct{}

type GetBlockChainParamsRsp struct {
	ChainProtocol            string  `json:"chain-protocol"`
	ChainDescription         string  `json:"chain-description"`
	RootStreamName           string  `json:"root-stream-name"`
	RootStreamOpen           bool    `json:"root-stream-open"`
	ChainIsTestnet           bool    `json:"chain-is-testnet"`
	TargetBlockTime          int     `json:"target-block-time"`
	MaximumBlockSize         int     `json:"maximum-block-size"`
	DefaultNetworkPort       int     `json:"default-network-port"`
	DefaultRPCPort           int     `json:"default-rpc-port"`
	AnyoneCanConnect         bool    `json:"anyone-can-connect"`
	AnyoneCanSend            bool    `json:"anyone-can-send"`
	AnyoneCanReceive         bool    `json:"anyone-can-receive"`
	AnyoneCanReceiveEmpty    bool    `json:"anyone-can-receive-empty"`
	AnyoneCanCreate          bool    `json:"anyone-can-create"`
	AnyoneCanIssue           bool    `json:"anyone-can-issue"`
	AnyoneCanMine            bool    `json:"anyone-can-mine"`
	AnyoneCanActivate        bool    `json:"anyone-can-activate"`
	AnyoneCanAdmin           bool    `json:"anyone-can-admin"`
	SupportMinerPrecheck     bool    `json:"support-miner-precheck"`
	AllowArbitraryOutputs    bool    `json:"allow-arbitrary-outputs"`
	AllowP2ShOutputs         bool    `json:"allow-p2sh-outputs"`
	AllowMultisigOutputs     bool    `json:"allow-multisig-outputs"`
	SetupFirstBlocks         int     `json:"setup-first-blocks"`
	MiningDiversity          float64 `json:"mining-diversity"`
	AdminConsensusUpgrade    float64 `json:"admin-consensus-upgrade"`
	AdminConsensusAdmin      float64 `json:"admin-consensus-admin"`
	AdminConsensusActivate   float64 `json:"admin-consensus-activate"`
	AdminConsensusMine       float64 `json:"admin-consensus-mine"`
	AdminConsensusCreate     float64 `json:"admin-consensus-create"`
	AdminConsensusIssue      float64 `json:"admin-consensus-issue"`
	LockAdminMineRounds      int     `json:"lock-admin-mine-rounds"`
	MiningRequiresPeers      bool    `json:"mining-requires-peers"`
	MineEmptyRounds          float64 `json:"mine-empty-rounds"`
	MiningTurnover           float64 `json:"mining-turnover"`
	FirstBlockReward         int     `json:"first-block-reward"`
	InitialBlockReward       int     `json:"initial-block-reward"`
	RewardHalvingInterval    int     `json:"reward-halving-interval"`
	RewardSpendableDelay     int     `json:"reward-spendable-delay"`
	MinimumPerOutput         int     `json:"minimum-per-output"`
	MaximumPerOutput         int64   `json:"maximum-per-output"`
	MinimumRelayFee          int     `json:"minimum-relay-fee"`
	NativeCurrencyMultiple   int     `json:"native-currency-multiple"`
	SkipPowCheck             bool    `json:"skip-pow-check"`
	PowMinimumBits           int     `json:"pow-minimum-bits"`
	TargetAdjustFreq         int     `json:"target-adjust-freq"`
	AllowMinDifficultyBlocks bool    `json:"allow-min-difficulty-blocks"`
	OnlyAcceptStdTxs         bool    `json:"only-accept-std-txs"`
	MaxStdTxSize             int     `json:"max-std-tx-size"`
	MaxStdOpReturnsCount     int     `json:"max-std-op-returns-count"`
	MaxStdOpReturnSize       int     `json:"max-std-op-return-size"`
	MaxStdOpDropsCount       int     `json:"max-std-op-drops-count"`
	MaxStdElementSize        int     `json:"max-std-element-size"`
	ChainName                string  `json:"chain-name"`
	ProtocolVersion          int     `json:"protocol-version"`
	NetworkMessageStart      string  `json:"network-message-start"`
	AddressPubkeyhashVersion string  `json:"address-pubkeyhash-version"`
	AddressScripthashVersion string  `json:"address-scripthash-version"`
	PrivateKeyVersion        string  `json:"private-key-version"`
	AddressChecksumValue     string  `json:"address-checksum-value"`
	GenesisPubkey            string  `json:"genesis-pubkey"`
	GenesisVersion           int     `json:"genesis-version"`
	GenesisTimestamp         int     `json:"genesis-timestamp"`
	GenesisNbits             int     `json:"genesis-nbits"`
	GenesisNonce             int     `json:"genesis-nonce"`
	GenesisPubkeyHash        string  `json:"genesis-pubkey-hash"`
	GenesisHash              string  `json:"genesis-hash"`
	ChainParamsHash          string  `json:"chain-params-hash"`
}

type GetRuntimeParamsReq struct{}

type GetRuntimeParamsRsp struct {
	Port                 int     `json:"port"`
	Reindex              bool    `json:"reindex"`
	Rescan               bool    `json:"rescan"`
	Txindex              bool    `json:"txindex"`
	Autocombineminconf   int     `json:"autocombineminconf"`
	Autocombinemininputs int     `json:"autocombinemininputs"`
	Autocombinemaxinputs int     `json:"autocombinemaxinputs"`
	Autocombinedelay     int     `json:"autocombinedelay"`
	Autocombinesuspend   int     `json:"autocombinesuspend"`
	Autosubscribe        string  `json:"autosubscribe"`
	Handshakelocal       string  `json:"handshakelocal"`
	Bantx                string  `json:"bantx"`
	Lockblock            string  `json:"lockblock"`
	Hideknownopdrops     bool    `json:"hideknownopdrops"`
	Maxshowndata         int     `json:"maxshowndata"`
	Miningrequirespeers  bool    `json:"miningrequirespeers"`
	Mineemptyrounds      float64 `json:"mineemptyrounds"`
	Miningturnover       float64 `json:"miningturnover"`
	Lockadminminerounds  int     `json:"lockadminminerounds"`
	Gen                  bool    `json:"gen"`
	Genproclimit         int     `json:"genproclimit"`
}

type SetRuntimeParamReq struct {
	Param string
	Value interface{}
}

type SetRuntimeParamRsp struct{}

type HelpReq struct{}

type HelpRsp string

type StopReq struct{}

type StopRsp string

type AddMultiSigAddressReq struct {
	NRequired int      `json:"nrequired"`
	Addresses []string `json:"addresses"`
}

type AddMultiSigAddressRsp string

type GetAddressesReq struct{}

type GetAddressesRsp []string

type GetNewAddressReq struct{}

type GetNewAddressRsp string

type ImportAddressReq []struct {
	Address string `json:"address"`
	Label   string `json:"label,omitempty"`
	Rescan  bool   `json:"rescan,omitempty"`
}

type ImportAddressRsp struct{}

type ListAddressesReq struct{}

type ListAddressesRsp []struct {
	Address string `json:"address"`
	Ismine  bool   `json:"ismine"`
}

type CreateKeyPairsReq struct{}

type CreateKeyPairsRsp []struct {
	Address string `json:"address"`
	Pubkey  string `json:"pubkey"`
	Privkey string `json:"privkey"`
}

type CreateMultiSigReq struct {
	NRequired int      `json:"nrequired"`
	Addresses []string `json:"addresses"`
}

type CreateMultiSigRsp struct {
	Address      string `json:"address"`
	RedeemScript string `json:"redeemScript"`
}

type ValidateAddressReq string

type ValidateAddressRsp struct {
	Isvalid      bool   `json:"isvalid"`
	Address      string `json:"address"`
	Ismine       bool   `json:"ismine"`
	Isscript     bool   `json:"isscript"`
	Pubkey       string `json:"pubkey"`
	Iscompressed bool   `json:"iscompressed"`
	Account      string `json:"account"`
}

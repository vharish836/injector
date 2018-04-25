package injector

type Getinforeq struct{}

type Getinforsp struct {
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
	ChainProtocol            string  `json:"chainprotocol"`
	ChainDescription         string  `json:"chaindescription"`
	RootStreamName           string  `json:"rootstreamname"`
	RootStreamOpen           bool    `json:"rootstreamopen"`
	ChainIsTestnet           bool    `json:"chainistestnet"`
	TargetBlockTime          int     `json:"targetblocktime"`
	MaximumBlockSize         int     `json:"maximumblocksize"`
	DefaultNetworkPort       int     `json:"defaultnetworkport"`
	DefaultRPCPort           int     `json:"defaultrpcport"`
	AnyoneCanConnect         bool    `json:"anyonecanconnect"`
	AnyoneCanSend            bool    `json:"anyonecansend"`
	AnyoneCanReceive         bool    `json:"anyonecanreceive"`
	AnyoneCanReceiveEmpty    bool    `json:"anyonecanreceiveempty"`
	AnyoneCanCreate          bool    `json:"anyonecancreate"`
	AnyoneCanIssue           bool    `json:"anyonecanissue"`
	AnyoneCanMine            bool    `json:"anyonecanmine"`
	AnyoneCanActivate        bool    `json:"anyonecanactivate"`
	AnyoneCanAdmin           bool    `json:"anyonecanadmin"`
	SupportMinerPrecheck     bool    `json:"supportminerprecheck"`
	AllowArbitraryOutputs    bool    `json:"allowarbitraryoutputs"`
	AllowP2ShOutputs         bool    `json:"allowp2shoutputs"`
	AllowMultisigOutputs     bool    `json:"allowmultisigoutputs"`
	SetupFirstBlocks         int     `json:"setupfirstblocks"`
	MiningDiversity          float64 `json:"miningdiversity"`
	AdminConsensusUpgrade    float64 `json:"adminconsensusupgrade"`
	AdminConsensusAdmin      float64 `json:"adminconsensusadmin"`
	AdminConsensusActivate   float64 `json:"adminconsensusactivate"`
	AdminConsensusMine       float64 `json:"adminconsensusmine"`
	AdminConsensusCreate     float64 `json:"adminconsensuscreate"`
	AdminConsensusIssue      float64 `json:"adminconsensusissue"`
	LockAdminMineRounds      int     `json:"lockadminminerounds"`
	MiningRequiresPeers      bool    `json:"miningrequirespeers"`
	MineEmptyRounds          float64 `json:"mineemptyrounds"`
	MiningTurnover           float64 `json:"miningturnover"`
	FirstBlockReward         int     `json:"firstblockreward"`
	InitialBlockReward       int     `json:"initialblockreward"`
	RewardHalvingInterval    int     `json:"rewardhalvinginterval"`
	RewardSpendableDelay     int     `json:"rewardspendabledelay"`
	MinimumPerOutput         int     `json:"minimumperoutput"`
	MaximumPerOutput         int64   `json:"maximumperoutput"`
	MinimumRelayFee          int     `json:"minimumrelayfee"`
	NativeCurrencyMultiple   int     `json:"nativecurrencymultiple"`
	SkipPowCheck             bool    `json:"skippowcheck"`
	PowMinimumBits           int     `json:"powminimumbits"`
	TargetAdjustFreq         int     `json:"targetadjustfreq"`
	AllowMinDifficultyBlocks bool    `json:"allowmindifficultyblocks"`
	OnlyAcceptStdTxs         bool    `json:"onlyacceptstdtxs"`
	MaxStdTxSize             int     `json:"maxstdtxsize"`
	MaxStdOpReturnsCount     int     `json:"maxstdopreturnscount"`
	MaxStdOpReturnSize       int     `json:"maxstdopreturnsize"`
	MaxStdOpDropsCount       int     `json:"maxstdopdropscount"`
	MaxStdElementSize        int     `json:"maxstdelementsize"`
	ChainName                string  `json:"chainname"`
	ProtocolVersion          int     `json:"protocolversion"`
	NetworkMessageStart      string  `json:"networkmessagestart"`
	AddressPubkeyhashVersion string  `json:"addresspubkeyhashversion"`
	AddressScripthashVersion string  `json:"addressscripthashversion"`
	PrivateKeyVersion        string  `json:"privatekeyversion"`
	AddressChecksumValue     string  `json:"addresschecksumvalue"`
	GenesisPubkey            string  `json:"genesispubkey"`
	GenesisVersion           int     `json:"genesisversion"`
	GenesisTimestamp         int     `json:"genesistimestamp"`
	GenesisNbits             int     `json:"genesisnbits"`
	GenesisNonce             int     `json:"genesisnonce"`
	GenesisPubkeyHash        string  `json:"genesispubkeyhash"`
	GenesisHash              string  `json:"genesishash"`
	ChainParamsHash          string  `json:"chainparamshash"`
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
	ParameterName  string      `json:"parametername"`
	ParameterValue interface{} `json:"parametervalue"`
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

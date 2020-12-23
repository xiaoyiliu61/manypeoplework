package entity

type RPCGetBlockChainInfo struct {
	/*private String chain;
	private int blocks;
	private int  headers;
	private String  bestblockhash;
	private double  difficulty;
	private long  mediantime;
	private double   verificationprogress;
	private boolean   initialblockdownload;
	private String  chainwork;
	private long  size_on_disk;
	private boolean  pruned;
//
	private int  pruneheight;
	private boolean   automatic_pruning;
	private long   prune_target_size;*/
     Automatic_pruning   bool  `json:"automatic_pruning"`
     Bestblockhash  string `json:"bestblockhash"`
     Chain   string `json:"chain"`
     Blocks  int64 `json:"blocks"`
     Bip65    Bip  `json:"bip65"`
}

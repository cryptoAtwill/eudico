TENDERMINT_PATH="~/Projects/tendermint"
NODE_0="127.0.0.1:26657"
NODE_1="127.0.0.1:26660"

NODE_0_PATH="~/.eudico-node0/"
NODE_1_PATH="~/.eudico-node1/"

NODE_0_API="1234"
NODE_1_API="1235"

NODE_0_KEY="/Users/alpha/Projects/tendermint/build/node0/config/priv_validator_key.json"
NODE_1_KEY="/Users/alpha/Projects/tendermint/build/node1/config/priv_validator_key.json"

#rm -rf ./eudico
#make eudico

rm -rvf ~/.eudico
rm -rvf ~/.eudico-node0
rm -rvf ~/.eudico-node1
rm -rvf ~/.eudico-node2
rm -rf ~/Projects/tendermint/build/node*

rm -rf ./term*.log
touch ./term1.log
touch ./term2.log
touch ./term3.log
touch ./term4.log


tmux new-session -d -s "tendermint" \; \
  split-window -t "tendermint:0" -v -l 66% \; \
  split-window -t "tendermint:0.1" -v -l 50% \; \
  split-window -t "tendermint:0.1" -h \; \
  split-window -t "tendermint:0.3" -h \; \
  \
  send-keys -t "tendermint:0.0" "./eudico tendermint application -addr=tcp://0.0.0.0:26650 & " Enter \; \
  send-keys -t "tendermint:0.0" "./eudico tendermint application -addr=tcp://0.0.0.0:26651 & " Enter \; \
  send-keys -t "tendermint:0.0" "./eudico tendermint application -addr=tcp://0.0.0.0:26652 & " Enter \; \
  send-keys -t "tendermint:0.0" "./eudico tendermint application -addr=tcp://0.0.0.0:26653 & " Enter \; \
  send-keys -t "tendermint:0.0" "cd $TENDERMINT_PATH; make localnet-start" Enter \; \
  \
  send-keys -t "tendermint:0.1" ";
        export EUDICO_PATH=$NODE_0_PATH
        ./eudico tspow daemon --genesis=./testdata/pow.gen --api=$NODE_0_API > term1.log 2>&1 &
            tail -f term1.log" Enter \; \
  send-keys -t "tendermint:0.2" "
        export EUDICO_PATH=$NODE_0_PATH
        ./eudico wait-api;
        ./scripts/wait-for-it.sh -t 0 $NODE_0 -- sleep 1;
        ./eudico wallet import-tendermint-key --as-default -path=$NODE_0_KEY;
        ./eudico tspow miner --default-key > term2.log 2>&1 &
            tail -f term2.log" Enter \; \
 send-keys -t "tendermint:0.3" ";
         export EUDICO_PATH=$NODE_1_PATH
         ./eudico tspow daemon --genesis=./testdata/pow.gen --api=$NODE_1_API > term3.log 2>&1 &
             tail -f term3.log" Enter \; \
 send-keys -t "tendermint:0.4" "
         export EUDICO_PATH=$NODE_1_PATH
         ./eudico wait-api;
         ./scripts/wait-for-it.sh -t 0 $NODE_1 -- sleep 1;
         ./eudico wallet import-tendermint-key --as-default -path=$NODE_1_KEY" Enter \; \
attach-session -t "tendermint:0.4"


# export EUDICO_TENDERMINT_RPC=http://"127.0.0.1:26657"
# export EUDICO_TENDERMINT_RPC=http://"127.0.0.1:26660"
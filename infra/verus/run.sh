#!/bin/bash
ADDRESS=$1
KEY=$2
# Start
mkdir "vrsctest"
verusd \
  -chain="vrsctest" \
  -ac_halving=7900004 \
  -datadir=./vrsctest \
  -port=10000 \
  -rpcport=10001 &> vrsctest0.log &

sleep 10

verusd \
  -chain="vrsctest" \
  -ac_halving=7900004 \
  -mineraddress=$ADDRESS \
  -minetolocalwallet=0 \
  -gen \
  -genproclimit=1 \
  -connect=127.0.0.1:10000

sleep 10

echo "VERUS_ADDRESS=$ADDRESS"

# Import the address
verus -chain=vrsctest importprivkey $KEY "" true

# Simulate mining
while :
do
    sleep 100
done
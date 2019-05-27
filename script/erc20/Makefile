contract:=erc20

default:go

sol:
	solc --abi ${contract}.sol | awk '/JSON ABI/{x=1;next}x' > ${contract}.json
	solc --bin ${contract}.sol | awk '/Binary:/{x=1;next}x' > ${contract}.bin

go:
	abigen --bin=${contract}.bin --abi=${contract}.json --pkg=${contract} --out=${contract}.go

clean:
	rm ${contract}.json
	rm ${contract}.bin
	rm ${contract}.go
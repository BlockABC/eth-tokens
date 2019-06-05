## 编译
运行make即可

## 运行
工程分成两个部分，erc20以及erc721
### ERC20
如果需要从my easy wallet上拉取数据执行
```bash
./eth_token -erc20 -g
```
如果从ether scan拉取数据执行：
```bash
./eth_token -erc20 -e
```
仅仅从本地tokens目录构建则执行
```bash
./eth_token -erc20
```

### ERC721
如果需要从ETHer Scan上拉取前50个NFT代币，则执行
```bash
./eth_token -nft -build
```
仅仅从本地tokens目录构建则执行
```bash
./eth_token -nft
```
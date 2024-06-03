const ethers = require("ethers")
const fs = require("fs-extra")
async function main(){
    /**
     * 0.8.19往上版本会报错提示 VM Exception while processing transaction: invalid opcode
     */
    //HTTP://10.90.24.158:7545172.22.48.1:7545
        const provider = new ethers.providers.JsonRpcProvider(`https://sepolia.infura.io/v3/477d001c968c47d2b1529e1ccac39b0a`);
        // const provider = new ethers.providers.InfuraProvider(`https://mainnet.infura.io/v3/477d001c968c47d2b1529e1ccac39b0a`);

        // const provider = new ethers.providers.JsonRpcProvider(`http://172.22.48.1:7545`);

            const wallet = new ethers.Wallet("7e5a65a881069a63b430c2430dbb67f45b30f5b90f3c47813d0e0cb4928ef5f2",provider);
    const abi = fs.readFileSync("./SimpleStorage_sol_SimpleStorage.abi","utf8");
    const binary = fs.readFileSync("./SimpleStorage_sol_SimpleStorage.bin","utf8");
    const contractFactory = new ethers.ContractFactory(abi,binary,wallet);
    console.log('deploying,please wait...');
    //设置gas
    const overrides = {
		gasPrice: 10000000000, // Can set this >= to the number read from Ganache window
		gasLimit: 6721975, // Use the same gasLimit as read from Ganache window (or a bit higher if still having issue)
	};
    const contract = await contractFactory.deploy(overrides);
    //交易回执
    const deploymentReceipt = await contract.deployTransaction.wait(1);
    console.log(contract);
    console.log(deploymentReceipt)
};

main().then(()=> process.exit(0))
    .catch((error)=>{
        console.error(error)
        process.exit(1)
});
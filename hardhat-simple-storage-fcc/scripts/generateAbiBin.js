const fs = require("fs");
const path = require("path");
const solc = require("solc");

async function main() {
  // 获取合约文件夹路径
  const contractsPath = "./contracts"; // 修改为你的合约文件夹路径
  // 读取合约文件夹中的所有文件
  fs.readdirSync(contractsPath).forEach((file) => {
    // 获取文件路径
    const filePath = path.join(contractsPath, file);
    // 读取合约文件内容
    const contractSource = fs.readFileSync(filePath, "utf8");
    // 编译合约
    const input = {
      language: "Solidity",
      sources: {
        [file]: {
          content: contractSource,
        },
      },
      settings: {
        outputSelection: {
          "*": {
            "*": ["*"],
          },
        },
      },
    };
    const compiledContract = JSON.parse(solc.compile(JSON.stringify(input)));
    // 获取合约的名称
    const contractName = path.basename(file, ".sol");
    // 获取合约的 ABI
    const abi = compiledContract.contracts[file][contractName].abi;
    // 获取合约的字节码
    const bytecode = compiledContract.contracts[file][contractName].evm.bytecode.object;
    // 生成 ABI 文件
    fs.writeFileSync(`${filePath}.abi`, JSON.stringify(abi, null, 2));
    // 生成 BIN 文件
    fs.writeFileSync(`${filePath}.bin`, bytecode);
    console.log(`ABI and BIN files generated for ${contractName}`);
  });
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

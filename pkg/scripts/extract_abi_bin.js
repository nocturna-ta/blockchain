const fs = require('fs')
const path = require('path')

const contractJsonPath = path.join(__dirname, '..', 'truffle', 'build', 'contracts', 'SimpleStorage.json')

const contractJson = JSON.parse(fs.readFileSync(contractJsonPath, 'utf8'))

const abi = JSON.stringify(contractJson.abi, null, 2)
const byteCode = contractJson.bytecode

fs.writeFileSync("../contracts/SimpleStorage.abi", abi, 'utf-8');
fs.writeFileSync("../contracts/SimpleStorage.bin", byteCode, 'utf-8');


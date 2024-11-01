const fs = require('fs')
const path = require('path')
const {exec} = require('child_process')

const abiPath = path.resolve(__dirname, '..', 'contracts', 'SimpleStorage.abi')
const binPath = path.resolve(__dirname, '..', 'contracts', 'SimpleStorage.bin')

const outputPath = path.resolve(__dirname, '..', 'contracts', 'simple_storage.go')
const packageName = 'contracts'

exec(`abigen --abi ${abiPath} --bin=${binPath} --pkg ${packageName} --type SimpleStorage --out ${outputPath}`, (error, stdout ,stderr) =>{
    if(error){
        console.error(`Error generating go bindings : ${error}` )
        return;
    }
    console.log('Go bindings generated successfully')
})
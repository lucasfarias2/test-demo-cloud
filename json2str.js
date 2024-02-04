const fs = require('fs');

function jsonToEnvString(filePath) {
    let rawData = fs.readFileSync(filePath);
    let data = JSON.parse(rawData);
    return JSON.stringify(data).replace(/\n/g, '\\n');
}

let escapedJsonString = jsonToEnvString('./service-account.json');
console.log(escapedJsonString);

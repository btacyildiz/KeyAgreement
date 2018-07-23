var readline = require('readline-sync');
const p = 11
const q = 5
const a = 3
const participantCount = 3

var id = parseInt(readline.question("Enter ID:"))
var r = parseInt(readline.question("Enter private key:"))

console.log("Calculating Z...")
var z = Math.pow(a,r) % p
console.log("id: " + id + " z: " + z)
var zStr = readline.question("Enter Z List: ")
var zMap = JSON.parse(zStr)
console.log(zMap['Z1'])
console.log("Calculating X...")
var before = id - 1
var after = id + 1 
if((id - 1) == 0){
    before = participantCount
}
if((id + 1) > participantCount){
    after = 1
}
var ZBefore = "Z" + before
var ZAfter = "Z" + after
console.log("Before: (" + ZBefore + ": "  + parseInt(zMap[ZBefore]) +") After: ("+ ZAfter + ": " + parseInt(zMap[ZAfter])+")")
console.log("Exp: " + Math.pow( (parseInt(zMap[ZAfter]) / parseInt(zMap[ZBefore]))  , r))
var x = Math.pow( (parseInt(zMap[ZAfter]) / parseInt(zMap[ZBefore])) , r) % p
console.log("id: " + id + " x: "+ x)

var xMap = JSON.parse(readline.question("Enter X List: "))
console.log("Calculating Session Key...")

before = id - 1
if (before == 0){
    before = participantCount
} 
var beforeZ = zMap[("Z"+before)]
var multi = Math.pow(beforeZ, (participantCount*r))
var j = id
for(var i= participantCount -1 ; i>=1 ; i--){
    console.log("i= " + i + " j= " + j + " X"+j+"^"+i)
    var x = xMap["X" + j]
    multi = multi * Math.pow(x, i)
    j++
    if(j>3){
        j = 1
    }
}
console.log("ID:" + id +" Key: " + (multi%p))





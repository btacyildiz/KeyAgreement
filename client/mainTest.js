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
var x = Math.pow( parseInt(zMap[ZAfter]), r) % p
var y = Math.pow( parseInt(zMap[ZBefore]), r) % p

console.log("id: " + id + " x: "+ x + " y: "+ y)

var xMap = JSON.parse(readline.question("Enter X List: "))
console.log("Calculating Session Key...")

before = id - 1
if (before == 0){
    before = participantCount
} 
var beforeZ = zMap[("Z"+before)]
console.log("P: " + p +" R: " + r +  " PCount: " + participantCount + " Z: "+ beforeZ )

console.log("Pow: " +  Math.pow(beforeZ, (participantCount*r)))
var multi = Math.pow(beforeZ, participantCount) % p
multi = (Math.pow(multi, r) % p )
console.log("Multi X : " + multi)
var muiltiY = 1
var i = id
for(var n= participantCount -1 ; n>=1 ; n--){
    console.log("n= " + n + " i= " + i + " X"+i+"^"+n)
    var x = xMap["X" + i]
    var y = xMap["Y" + i]
    console.log( " X"+ n +  ":" +  + (Math.pow(x, n) % p))
    console.log( " Y"+ n +  ":" +  + (Math.pow(y, n) % p))

    multi = multi * (Math.pow(x, n) % p)
    muiltiY = muiltiY * (Math.pow(y, n) % p)
    i++
    if(i>participantCount){
        i = 1
    }
}
console.log("MultiX: " + multi % p  + " MultiY: " + muiltiY % p )
var sec1 = (multi%p)
var sec2 = (muiltiY % p)

while(sec1 % sec2 != 0){
    sec1 += p
}

console.log("ID:" + id +" Key: " + sec1 / sec2)





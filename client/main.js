const p = 7
const q = 3
const a = 2
const participantCount = 3

// Generate Z 
var id = process.argv[2]
var r = process.argv[3]
var command = process.argv[4]
if(command == "-calcz"){
   var z = Math.pow(a,r) % p
   console.log("id: " + id + " z: " + z)
}else if(command == "-calcx"){
   var zMap = JSON.parse(process.argv[5])
   var before = id -1
   var after = id + 1 
   if((id - 1) == 0){
     before = participantCount
   }
   if((id + 1) > participantCount){
     after = 1
   }
   var ZBefore = "Z" + before
   var ZAfter = "Z" + after
   var X = Math.pow( (parseInt(zMap[ZAfter]) / parseInt(zMap[ZBefore])) , r) % p
   console.log("id: " + id + " x: "+ X)
}else if(command == "-calckey"){
    var zMap = JSON.parse(process.argv[5])
    var xMap = JSON.parse(process.argv[6])
    var before = id - 1
    if (before == 0){
        before = participantCount
    } 
    var beforeZ = zMap[("Z"+before)]
    var multi = Math.pow(beforeZ, (participantCount*r))
    var j = 1
    for(var i= participantCount -1 ; i>=1 ; i--){
        var x = xMap["X" + j]
        multi = multi * Math.pow(x, i)
        j++
    }
    console.log("Key: " + (multi%p))
}






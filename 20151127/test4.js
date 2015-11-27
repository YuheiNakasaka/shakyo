var stream = new kamo.Stream();

var mappedStream = stream.map(function(m){
  return m * 2;
});

mappedStream.subscribe(function(m) {
  console.log(m);
})

var filterdStream = stream.filter(function(m){
  return m % 2 == 1;
})

filterdStream.subscribe(function(m){
  console.log(m)
});

var scannedStream = stream.scan(0, function(previusMessage, message) {
  return previusMessage + message;
});

scannedStream.subscribe(function(m) {
  console.log(m);
})

stream.publish(1);
stream.publish(2);
stream.publish(3);

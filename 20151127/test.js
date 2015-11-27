var stream = new kamo.Stream();

stream.subscribe(function(m){
  console.log(m);
})

stream.publish("test");

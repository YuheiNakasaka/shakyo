var codes = [
  38,
  38,
  40,
  40,
  37,
  39,
  37,
  39,
  66,
  65
]

kamo.Stream.fromEventHandlerSetter(window, 'onkeyup')
.map(function(e){return e.keyCode})
.scan([], function(q,code){return q.concat([code]).slice(-codes.length);})
.filter(function(q){return q.length === codes.length})
.filter(function(q){return JSON.stringify(q) === JSON.stringify(codes)})
.subscribe(function(){alert('Conguratulation!')});
// var stream = new kamo.Stream();

// window.onlick = function(e) {
//   stream.publish(e);
// }

var stream = kamo.Stream.fromEventHandlerSetter(window, 'onclick');
var circle;
var speed = 30;
var radius = 100;

// 初期化でcallされる
function setup() {
  createCanvas(500, 500);
  circle = new MyCircle();
}

// 毎フレームcallされる
function draw() {
  background(255);
  circle.move();
}

function MyCircle() {
  this.x = width/2;
  this.y = height/2;
  this.degree = 0;
  this.rad = 0;
  this.radius = radius;

  this.move = function() {
    this.degree += 3;
    this.rad = this.degree * Math.PI / 180;
    this.x = width/2 + this.radius * Math.cos(this.rad);
    this.y = height/2 + this.radius * Math.sin(this.rad);

    ellipse(this.x, this.y, 30, 30);
  }
}
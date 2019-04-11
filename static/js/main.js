var ws;

function setup() {
    createCanvas(windowWidth, windowHeight);
    background(0);

    ws = new WebSocket("ws://localhost:8000/ws")
    ws.onopen = function() {
        console.log("socket connected");
    }

    ws.onmessage = function(event) {
        var msg = JSON.parse(event.data);
        fill(0, 255, 255);
        noStroke();
        ellipse(msg.x,msg.y, 20, 20);
    }

    ws.onclose = function() {
        console.log("socket closed");
    }
}


function draw() {
    var wury = "wuriyanto.com";
    var year = "2019"
    textSize(20);
    fill(255, 255, 0);
    text(wury, 10, 30);
    textSize(18);
    fill(255, 255, 0);
    text(year, 10, 50);
}

function mouseDragged() {
    fill(255, 0, 255);
    noStroke();
    ellipse(mouseX,mouseY, 20, 20);
    sendMouse(mouseX,mouseY);
  }
  
  
function sendMouse(xpos, ypos) {
    var data = {
        x: xpos,
        y: ypos
    };
    ws.send(JSON.stringify(data));
}
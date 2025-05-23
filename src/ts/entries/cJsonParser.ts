const NUMBER_OF_POINTS = 20;

class Vec2 {
    x: number;
    y: number;

    constructor(x = 0, y = 0) {
        this.x = x;
        this.y = y;
    }

    add(vec: Vec2) {
        this.x += vec.x;
        this.y += vec.y;
    }
}

type Point = {
    pos: Vec2;
    vel: Vec2;
};

let element: HTMLDivElement;

// returns random number from 0 to n
function rand(n: number) {
    return Math.round(Math.random() * n);
}

function initPoints(): Point[] {
    let points: Point[] = [];

    for(let i = 0; i < NUMBER_OF_POINTS; i++) {
        const pos = new Vec2();
        const maxPos = 100;
        pos.x = rand(maxPos);
        pos.y = rand(maxPos);

        const vel = new Vec2();
        const minVel = 0.1;
        const weight = 0.1;
        // the first part gives us a number from -1 to 1
        vel.x = (1 - rand(2)) * weight + minVel;
        vel.y = (1 - rand(2)) * weight + minVel;

        points.push({
            pos,
            vel,
        });
    }

    return points;
}

const points: Point[] = initPoints();

function updatePoints() {
    for(const point of points) {
        point.pos.add(point.vel);

        if(point.pos.x < 0) {
            point.pos.x = 0;
            point.vel.x *= -1;
        } else if(point.pos.x > 100) {
            point.pos.x = 100;
            point.vel.x *= -1;
        }

        if(point.pos.y < 0) {
            point.pos.y = 0;
            point.vel.y *= -1;
        } else if(point.pos.y > 100) {
            point.pos.y = 100;
            point.vel.y *= -1;
        }
    }
}

// sets the polygon style (read more: https://developer.mozilla.org/en-US/docs/Web/CSS/basic-shape/polygon)
function updatePolygon() {
    let pointsString = "";

    for(let i = 0; i < points.length; i++) {
        if(i > 0) {
            pointsString += ",";
        }

        const point = points[i];
        pointsString += `${point.pos.x}% ${point.pos.y}%`;
    }

    element.style.clipPath = `polygon(${pointsString})`;
}

function loop() {
    if(!element) return;

    updatePoints();
    updatePolygon();

    requestAnimationFrame(loop);
}

window.onload = () => {
    const id = "c-json-parser-bg";
    element = document.getElementById(id) as HTMLDivElement;

    if(element) {
        loop();
    } else {
        console.error(`Element with id ${id} doesn't exist`);
    }
};

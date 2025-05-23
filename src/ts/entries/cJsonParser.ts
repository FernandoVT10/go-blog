const NUMBER_OF_POLYGONS = 20;

let element: HTMLDivElement;

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

function initPoints(): Point[] {
    let points: Point[] = [];

    for(let i = 0; i < NUMBER_OF_POLYGONS; i++) {
        const pos = new Vec2();
        pos.x = Math.floor(Math.random() * 100);
        pos.y = Math.floor(Math.random() * 100);

        const vel = new Vec2();
        vel.x = (1 - Math.round(Math.random() * 2)) * 0.1 + 0.1;
        vel.y = (1 - Math.round(Math.random() * 2)) * 0.1 + 0.1;

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

    if(!element) {
        console.error(`Element with id ${id} doesn't exist`);
    }

    requestAnimationFrame(loop);
};

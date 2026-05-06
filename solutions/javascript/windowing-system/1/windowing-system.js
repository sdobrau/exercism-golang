// @ts-check

/**
 * Implement the classes etc. that are needed to solve the
 * exercise in this file. Do not forget to export the entities
 * you defined so they are available for the tests.
 */

export class Size {
    constructor(width = 80, height = 60) {
	this.width = width;
	this.height = height;
    }

    resize(newWidth, newHeight) {
	this.width = newWidth;
	this.height = newHeight;
    }
}

export class Position {
    constructor(x = 0, y = 0) {
	this.x = x
	this.y = y
    }
    move(newX, newY) {
	this.x = newX;
	this.y = newY;
    }
}

export class ProgramWindow {
    constructor() {
	this.screenSize = new Size(800, 600)
	this.size = new Size()
	this.position = new Position()
    }
    resize(sizeArg) {
	if (sizeArg.width <= 0) {
	    sizeArg.width = 1
	}
	if (sizeArg.height <= 0) {
	    sizeArg.height = 1
	}
	// eg. if y 20 and height 600 > need 580 (600 - y)
	if (this.position.y + sizeArg.height > 600) {
	    sizeArg.height = 600 - this.position.y
	}
	// e.g. if x 300 and width 600 > 800, need 500 (800-x)
	if (this.position.x + sizeArg.width > 800) {
	    sizeArg.width = 800 - this.position.x
	}
	// then, set
	this.size.width = sizeArg.width
	this.size.height = sizeArg.height
    }
    move(positionArg) {
	if (positionArg.x <= 0) {
	    positionArg.x = 0
	}
	if (positionArg.y < 0) {
	    positionArg.y = 0
	}
	// e.g. if 250width,100height + move to 600,200
	// would move to 550(800-width),200 as width was too large
	if (positionArg.x + this.size.width > 800) {
	    positionArg.x = 800 - this.size.width
	}
	
	if (positionArg.y + this.size.height > 600) {
	    positionArg.y = 600 - this.size.height
	}
	// otherwise just set
	this.position.x = positionArg.x
	this.position.y = positionArg.y
    }
}

export function changeWindow(programWindow) {
    // taken from SleeplessByte

    // Move to the top-left corner first so it can always resize
    programWindow.move(new Position())
    
    programWindow.move(new Position(100, 150))
    programWindow.resize(new Size(400, 300))
    return programWindow
}


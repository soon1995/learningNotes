class Snake {
    head: HTMLElement
    bodies: HTMLCollection
    element: HTMLElement

    constructor() {
        // querySelector只会选第一个
        this.element = document.getElementById('snake')!
        this.head = document.querySelector('#snake > div')!
        this.bodies = this.element.getElementsByTagName('div')
    }

    get X() {
        return this.head.offsetLeft
    }

    get Y() {
        return this.head.offsetTop
    }

    set X(v: number) {

        if (this.X === v) {
            return
        }

        if (v < 0 || v > 290) {
            throw new Error('蛇撞墙了')
        }

        if (this.bodies[1] && (this.bodies[1] as HTMLElement).offsetLeft === v) {
            if (v > this.X) {
                v = this.X - 10
            } else {
                v = this.X + 10
            }
        }

        this.moveBody()

        this.head.style.left = v + 'px'

        this.checkHeadBody()
    }

    set Y(v: number) {

        if (this.Y === v) {
            return
        }

        if (v < 0 || v > 290) {
            throw new Error('蛇撞墙了')
        }

        if (this.bodies[1] && (this.bodies[1] as HTMLElement).offsetTop === v) {
            if (v > this.Y) {
                v = this.Y - 10
            } else {
                v = this.Y + 10
            }
        }

        this.moveBody()

        this.head.style.top = v + 'px'

        this.checkHeadBody()
    }

    addBody() {
        this.element.insertAdjacentHTML("beforeend", "<div></div>")
    }

    moveBody() {
        for (let i = this.bodies.length - 1; i > 0; i--) {
            let X = (this.bodies[i - 1] as HTMLElement).offsetLeft;
            let Y = (this.bodies[i - 1] as HTMLElement).offsetTop;

            (this.bodies[i] as HTMLElement).style.left = X + 'px';
            (this.bodies[i] as HTMLElement).style.top = Y + 'px';
        }
    }

    checkHeadBody() {
        for (let i = 1; i < this.bodies.length; i++) {
            let bd = this.bodies[i] as HTMLElement
            if (this.X === bd.offsetLeft && this.Y === bd.offsetTop) {
                throw new Error("撞到自己了")
            } 
        }
    }
}

export default Snake
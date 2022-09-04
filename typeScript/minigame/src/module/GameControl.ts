// 引入其他的类
import ScorePanel from "./ScorePanel"
import Food from "./Food"
import Snake from "./Snake"

// 控制其他所有类
class GameControl {
    snake: Snake
    food: Food
    scorePanel: ScorePanel
    direction: string = 'Right'
    isLive = true

    constructor() {
        this.snake = new Snake()
        this.food = new Food()
        this.scorePanel = new ScorePanel(10,1)

        this.init()
    }

    private init() {
        // cannot, this means document itself instead of GameControl
        //document.addEventListener('keydown', this.keydownHandler)
        document.addEventListener('keydown', this.keydownHandler.bind(this))
        this.move()
    }

    private keydownHandler(event: KeyboardEvent) {
        let key = event.key
        this.direction = event.key
    }

    // move according to direction
    private move() {
        let X = this.snake.X
        let Y = this.snake.Y
        switch (this.direction) {
            case "ArrowUp":
            case "Up":
                Y -= 10
                break
            case "ArrowDown":
            case "Down":
                Y += 10
                break
            case "ArrowLeft":
            case "Left":
                X -= 10
                break
            case "ArrowRight":
            case "Right":
                X += 10
                break
        }

        this.checkEat(X, Y)
        
        try {
            this.snake.X = X
            this.snake.Y = Y
        } catch (e) {
            // snake die
            alert((e as Error).message + " - GAME OVER")
            this.isLive = false
        }

        // enable automatic run
        this.isLive && setTimeout(this.move.bind(this), 300 - (this.scorePanel.level - 1) * 30)
    }

    // check if snake ate food
    private checkEat(X: number, Y: number) {
        if (X === this.food.X && Y === this.food.Y) {
            this.scorePanel.addScore()
            this.food.change()
            this.snake.addBody()
        }
    }
}

export default GameControl
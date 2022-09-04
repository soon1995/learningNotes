class ScorePanel {
    score = 0
    level = 1
    scoreEle: HTMLElement
    levelEle: HTMLElement

    // 设置一个变量设置数目
    maxLevel: number;
    upScore: number;

    constructor(maxLevel: number = 10, upScore: number = 10) {
        this.scoreEle = document.getElementById("score")!
        this.levelEle = document.getElementById("level")!
        this.maxLevel = maxLevel
        this.upScore = upScore
    }   
    
    // add score
    addScore() {
        this.scoreEle.innerHTML = ++this.score + ''

        if (this.score % this.upScore === 0) {
            this.levelUp()
        }
    }

    // level up
    private levelUp() {
        if (this.level < this.maxLevel) {
            this.levelEle.innerHTML = ++this.level + ''
        }
    }

}

// const scorePanel = new ScorePanel()
// for (let i = 0; i < 200; i++) {
//     scorePanel.addScore()
// }

export default ScorePanel
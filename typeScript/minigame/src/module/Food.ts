// 定义food
class Food {
    // 元素
    private element: HTMLElement;

    constructor() {
        this.element = document.getElementById("food")!
    }

    // get x and y
    get X() {
        return this.element.offsetLeft
    }

    get Y() {
        return this.element.offsetTop
    }

    change() {
        // 生成随机的位子
        // 最小0, 最大 290
        // 蛇移动一次一格10px，坐标必须是整10
        // random()不包括0 和1, round()过后包括0也包括1了
        let left = Math.round(Math.random() * 29) * 10
        let top = Math.round(Math.random() * 29) * 10

        this.element.style.left = left +'px'
        this.element.style.top = top + 'px'
    }
}

// const food = new Food()
// console.log(food.X, food.Y);
// food.change()
// console.log(food.X, food.Y);

export default Food;
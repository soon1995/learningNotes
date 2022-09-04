# TypeScript

This is a learning note, recorded from [尚硅谷TypeScript教程（李立超老师TS新课）](https://www.bilibili.com/video/BV1Xy4y1v7S2?p=1&vd_source=a788bdd4d7cdd9dfe02852346d523cb9)



Start : 3/9/2022

End : 5/9/2022

## Setting up

1. Download Node.js

   https://nodejs.org/en/

2. use npm download

   npm i -g typescript 

   **Tips: ** -g stands for global

3. 01_helloTS.ts

   ```ts
   console.log('hello TS');
   ```

4. cmd

   ```bash
   tsc 01_helloTS.ts
   ```

5. noticed a js file generated





## Type

> **Difference with js:** once set a type for a variable, that variable must follow this set type after.

```js
wrong:
let a: number; // let a = 10 / let a: number = 10
a = 10;
a = 'hello'; << wrong! Failed to compile too
```

> **Alias**
>
> ```typescript
> type MyType = 1 | 2 | 3
> let a: MyType
> let b: MyType
> ```



1. number

2. boolean

3. string

4. 字面量 (let a: ? | ?)

5. any (default), not recommended to use

6. unknown

   > Difference between any and unknown:
   >
   > ```typescript
   > let e: any;
   > let f: unknown;
   > let a: string;
   > a = e; // ok
   > a = f; // Not ok, type unknown cannot be assigned to string
   > 
   > if (typeof f === "string") {
   >     a = f; // ok
   > }
   > 
   > s = e as string; // assertion, use when we know this is string, if the type not suit => undefined
   > s = <string>e; // same as above
   > ```

7. void, usually use for return value of function

   > ok to return `undefined` or`null`

8. never

   > never return
   >
   > used to throw new Error('');

9. object, not recommended to use

   > {name:'孙悟空'}
   >
   > ```typescript
   > let a: object;
   > a = {}; // ok
   > a = function(){}; // ok
   > ```

10. array

    > [1,2,3]
    >
    > ```typescript
    > let e: string[]; // or let e: Array<string>
    > e = ['a','b','c'];
    > ```

11. tuple (like array, but the length is fixed, we called 元组)

    > [4,5]
    >
    > ```typescript
    > let h: [string, string]
    > h = ['hello', 'abc']
    > ```

12. enum

    > enum(A,B)
    >
    > ```typescript
    > enum Gender{
    >     Male,
    >     Female
    > }
    > 
    > let i : {name: string, gender: Gender}
    > i = {
    >     name: '孙悟空',
    >     gender: Gender.Male
    > }
    > 
    > console.log(i.gender === Gender.Male);
    > ```

13. Function

    ```typescript
    let d: (a: number, b: number) => number;
    
    d = function(n1: number, n2: number) {
        return n1 + n2;
    }
    console.log(d(1,2)); // 3
    ```

    



>  **Example** 

> **function**

```typescript
function sum(a: number, b: number): number {
    return a + b;
}
```

> **let a: ?**
>
> once used a: ?, the a cannot be modified to other value. Thus, normally we use it to fixed type or value

```typescript
let b: "male" | "female";
b = "male"; // ok
b = "hello"; // Not ok

let c: boolean | string;
c = true; // ok
c = 'hello'; // ok
c = 10; // Not ok

----
function
----
function a(a,b): string | boolean {
    return "hello"
}
```



> **How to declare an object if type object is not recommended?**
>
> **Tips: **? means can be null

```typescript
let b: {name: string, age?: 18}; // Use me
b = {name: '孙悟空'} // if age without ?, then this will be fault
b = {name: 123} // Not ok

let c: object
c = {name: '孙悟空'}
c = {name: 123} // ok
```



> **How to declare an object if I do not know the fields in that?**

```typescript
let c: {name: string, [propName: string]: any} // propName is not fixed string, can be [abc: string]
c = {name: '猪八戒', a: 'haha', b: 123}

```



## TsConfig

```bash
tsc xxx.ts -w // Watch input files. So there is no need to compile every time after making modification to that ts file, ctrl + c to quit
```



> **To watch all ts files**
>
> - To have config file --> tsconfig.json
> - `tsc -w`

tsconfig.json (yes, it has nothing inside)

```bash
{

}
```



> **Include**
>
> only compile selected ts
>
> ** 任意目录
>
> \* 任意文件

```json
{
    "include": [
        "./src/**/*"
    ]
}
```



> **Exclude**
>
> Default: ["node_modules", "bower_components", "jspm_packages"]

```json
{
    "exclude": [
        "./src/hello/**/*"
    ]
}
```



>  **Extend**
>
> 定义被继承的配置文件

```json
{
    "extends": "./configs/base"
}
```



> **Files**
>
> 指定被编译的文件，只有需要编译的文件少时才会用到

```json
{
    "files": [
        "core.ts",
        "sys.ts",
        "types.ts"
    ]
}
```



### **Compiler Options**

#### **Target**

> select the outcome version of ES, default ES3
>
> **Tips: **to know which options can be chosen, put the invalid string and see the error information

```json
{
    "compilerOptions": {
    	"target": "es6"
    }
}
```



#### **Module**

> select modular standard eg
>
> 'none', 'commonjs', 'amd', 'system', 'umd', 'es6', '**es2015**', 'es2020', 'es2022', 'esnext', 'node16', 'nodenext'

```json
{
    "compilerOptions": {
    	"module": "es2015"
    }
}
```

**Example**

```js
m.js
-----
export const hi = "hillo"

app.ts
------
import { hi } from "m.js"

console.log(hi)

app.js
------
import { hi } from "m.js";
console.log(hi);
```



#### **Lib**

> seldom used, we use default one
>
> select the library that to be used in this project
>
> eg. if none select `"lib":[]`, then the 提示for the sentence will be disable too. eg, document.getElementById(..), there will not be tips for input "doc".



#### **Out Directory**

> Where compiled js to be at. Default same place with ts

```json
{
    "compilerOptions": {
    	"outDir": "./dist"
    }
}
```



#### **Out File**

> Compile all ts compiled result to ONE js
>
> Only **amd** and **system** modules are supported

```json
{
    "compilerOptions": {
    	"outFile": "./dist/app.js"
    }
}
```



#### **Allow Js**

> Whether compile the js in the target also, default **false**

```json
{
    "compilerOptions": {
    	"allowJs": false
    }
}
```



#### **CheckJs**

> Whether to check js using the logic of ts, eg check types in js. Default **false**

```json
{
    "compilerOptions": {
    	"checkJs": false
    }
}
```



#### **Remove Comments**

> Options to have comment in compiled js. Default **false**

```json
{
    "compilerOptions": {
    	"removeComments": true
    }
}
```



#### **No Emit**

> do not generate js. Default **false**
>
> Seldom use, used when we need to check syntax but do not want to compile to js



#### No Emit On Error

> whether not to generate js files if there is error. Default **false**



#### Always Strict

> 用来设置编译后的文件是否使用严格模式("use strict" in js). Default **false**



#### No Implicit Any

> no default any for sentences. Default **false**

```typescript
when true
---------
function fn(a, b) {} // not pass
```



#### No Implicit This

> 不允许不明确类型的this. Default **false**



#### Strict Null Checks

> need to do if (x != null) {...} before the the used of null possible variable. Default **false**

```typescript
to solve
--------
if (box1 !== null) {
    box1.addEventListener('click', function() {
        alert('hello')
    })
}

or

box1?.addEventListener('click', function() {
    alert('hello')
})

or : when confirmed have
class Food {
    element: HTMLElement;

    constructor() {
        this.element = document.getElementById("food")!; // ! here
    }
}
```



#### Strict

> 所有的严格的开关，default **false**. **建议 true**





## Webpack

1. `npm init -y`
2. `npm i -D webpack webpack-cli typescript ts-loader` // ts-loader -> 整合ts 和 webpack， 有它typescript compiler才能在webpack使用
3. create new config file: `webpack.config.js`



### Config

`webpack.config.js`

```js
// 引入一个包
const path = require('path');

// code all webpack config here
module.exports = {
    // 指定入口文件
    entry: "./src/index.ts",
    
    // 指定文件打包文件
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: "bundle.js"
    },
    
    // 指定webpack打包时要使用的模块
    module: {
        // 指定要加载的规则
        rules: [
            {
                // test指定的是规则生成的文件, /是路径有这个名的
                test : /\.ts$/,
                // 要使用的loader
                use: 'ts-loader',
                exclude: /node-modules/
            }
        ]
    },
    
    // 用来设置引用模块，否则报错
    resolve: {
    	extensions: ['.ts','.js']
	}
}
```



**Tips: **remember to insert `"build": "webpack"` in package.json scripts. so that when we `npm run build` the webpack will run.



### Plugins for webpack

> **html-webpack-plugin**
>
> webpack中html插件，用来自动创建html文件。

1. `npm i -D html-webpack-plugin`

2. webpack.config.js

```js
const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
...

module.exports = {
   ...
   plugins: [
       ...
       new HtmlWebpackPlugin({
           // title:'TS测试' <- this will change the index.html title
           template: './src/index.html' //<- this will use index page we create.
       }),
   ]
}
```



> **webpack-dev-server**
>
> webpack的开发服务器， 当save了自动reflect in webpage

1. `npm i -D webpack-dev-server`

2. package.json

   ```json
   {
     ...
     "scripts": {
       ...
       "start": "webpack serve --open chrome.exe"
     },
     ...
   }
   
   ```

3. `npm start`

   - noticed whenever save in entry ts file, the webpage will reflect



> **clean-webpack-plugin**
>
> webpack中的清除插件，每次构建都会先清除目录,

1. `npm i -D clean-webpack-plugin`

2. webpack.config.js

   ```js
   ...
   const { CleanWebpackPlugin } = require('clean-webpack-plugin');
   
   // code all webpack config here
   module.exports = {
       ...
       plugins: [
           ...
           new CleanWebpackPlugin()
       ]
   }
   ```



### Babel and core-js for browser compatility

1. `npm i -D @babel/core @babel/preset-env babel-loader core-js`

   - babel-loader -> 结合babel 和 Webpack
   - core-js -> 让老版本浏览器用到新版本的技术, 东西比较多，babel里设置按需加载。例如：babel不能处理Promise, ，所以core-js会生成自己准备的Promise

2. webpack.config.js

   ```js
   // 引入一个包
   const path = require('path');
   const HtmlWebpackPlugin = require('html-webpack-plugin');
   const { CleanWebpackPlugin } = require('clean-webpack-plugin');
   
   // code all webpack config here
   module.exports = {
       // 指定入口文件
       entry: "./src/index.ts",
       
       // 指定文件打包文件
       output: {
           path: path.resolve(__dirname, 'dist'),
           filename: "bundle.js"
       },
       
       // 指定webpack打包时要使用的模块
       module: {
           // 指定要加载的规则
           rules: [
               {
                   // test指定的是规则生成的文件, /是路径有这个名的
                   test : /\.ts$/,
                   // 要使用的loader, 后面的先执行
                   use: [
                       'babel-loader',
                       'ts-loader'
                   ],
                   exclude: /node-modules/
               }
           ]
       },
       
       plugins: [
           new CleanWebpackPlugin(),
           new HtmlWebpackPlugin({
               // title:'TS测试' <- this will change the index.html title
               template: './src/index.html' //<- this will use index page we create.
           })
       ],
       
       // 用来设置引用模块，否则报错
       resolve: {
       	extensions: ['.ts','.js']
   	}
   }
   ```

3. 因为webpack生成的js会用箭头函数，在IE11并不适用，我们可以:

   ```js
   ...
   module.exports ={
       ...
       output: {
           ...
           environment: {
               arrowFunction: false,
               const: false
           }
       }
       ...
   }
   ```



### **Result**

webpack.config.js

```js
// 引入一个包
const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');

// code all webpack config here
module.exports = {
    // 指定入口文件
    entry: "./src/index.ts",
    
    // 指定文件打包文件
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: "bundle.js",
        environment: {
            arrowFunction: false,
            const: false
        }
    },
    
    // 指定webpack打包时要使用的模块
    module: {
        // 指定要加载的规则
        rules: [
            {
                // test指定的是规则生成的文件, /是路径有这个名的
                test : /\.ts$/,
                // 要使用的loader, 后面的先执行
                use: [
                    //配置babel
                    {
                        // 指定加载器
                        loader:'babel-loader',
                        options: {
                            // 设置自定义环境
                            presets: [
                                [
                                    // 指定环境的插件
                                    '@babel/preset-env',
                                    // 配置信息
                                    {   
                                        // 要兼容的目标浏览器
                                        targets: {
                                            'chrome':'58',
                                            'ie':'11'
                                        },
                                        // 3 是版本,请看package.json
                                        'corejs':'3',
                                        // 使用corejs的方式, usage是按需加载
                                        'useBuiltIns':'usage',
                                    }
                                ]
                            ]
                        }
                    },
                    'ts-loader'
                ],
                exclude: /node-modules/
            }
        ]
    },
    
    plugins: [
        new CleanWebpackPlugin(),
        new HtmlWebpackPlugin({
            // title:'TS测试' <- this will change the index.html title
            template: './src/index.html' //<- this will use index page we create.
        })
    ],
    
    // 用来设置引用模块，否则报错
    resolve: {
    	extensions: ['.ts','.js']
	}
}
```



package.json

```json
{
  "name": "part3",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "build": "webpack",
    "start": "webpack serve --open chrome.exe"
  },
  "keywords": [],
  "author": "",
  "license": "ISC",
  "devDependencies": {
    "@babel/core": "^7.18.13",
    "@babel/preset-env": "^7.18.10",
    "babel-loader": "^8.2.5",
    "clean-webpack-plugin": "^4.0.0",
    "core-js": "^3.25.0",
    "html-webpack-plugin": "^5.5.0",
    "ts-loader": "^9.3.1",
    "typescript": "^4.8.2",
    "webpack": "^5.74.0",
    "webpack-cli": "^4.10.0",
    "webpack-dev-server": "^4.10.1"
  }
}

```



## OOP

### Class

```ts
class Person {
    // we can not to declare the field by below methods:
    // class Person {
    //      constructor(public name: string, public age: number) {
    //          //no need to write anything here
    //      }    
    // }
    name: string
    age: number

    // 构造函数
    // constructor() {
    //     this.name = "Wong"
    //     this.age = 18
    // }

    constructor(name: string, age: number) {
        this.name = name
        this.age = age
    }

    // 静态属性，不用new也可用
    static hobby = "fishing"

    // 可读不可改
    readonly height = 180

    // 方法
    sayHello(): void {
        console.log("Hello")
    }

    // 静态方法() 
    static sayByebye() {
        console.log("Bye bye");
    }
}

const per = new Person("孙悟空", 18)
console.log(per)
console.log(per.name, per.age)
console.log(Person.hobby);

// 修改
per.age = 22
per.sayHello = function() {
    console.log(" what ");
}

// 方法调用
per.sayHello()
Person.sayByebye()

```

**Tips**

```typescript
class ScorePanel {
    maxLevel: number;
	
    // if no param inserted => use 10, else use param
    constructor(maxLevel: number = 10) {
        this.maxLevel = maxLevel
    }   
}
```





### Inheritance

```typescript
// Inheritance 继承
class Teacher extends Person{

    subject: string
    
    constructor(name: string, age: number, subject: string) {
        super(name, age)
        this.subject = subject
    }

    sayHello(): void {
        // use super function
        super.sayHello()
        
        // or `${this.name}...`
        console.log(this.name + ": Hello students")
    }
}

const teacher = new Teacher("Tan", 33, "Math")
teacher.sayHello()

```



### Abstract Class

```typescript
// Abstract Class
abstract class Transport {
    name: string
    
    constructor(name: string) {
        this.name = name
    }

    abstract move(): void
}

class Car extends Transport {
    move(): void {
        console.log(" 80 km/hr");
    }
}
```



### Interface

```typescript

// Interface
// 定义一个类的结构，应该包含什么属性和方法, 和type MyType={..} 很像
// 不能有实际的值
interface MyInterface {
    name: string
    saySomething(): void
}

class MyClass implements MyInterface {
    name: string

    constructor(name: string) {
        this.name = name
    }

    saySomething(): void {
        console.log("haha");      
    }
}
```



### Encapsulation

```typescript
// Encapsulation
// private - only for it self
// public (Default) - all 
// protected - for self and children
class Account {
    private _name: string
    private _balance: number
    private _age: number

    constructor(name: string, age: number) {
        this._name = name
        this._balance = 0
        this._age = age
    } 

    public expense(money: number) {
        if (this._balance >= money) {
            this._balance -= money
        }
        console.log("not enough balance");
    }

    public income(money: number) {
        this._balance += money
    }

    //TS methods/////////////////////
    // use account.age to call
    get age() {
        return this._age
    }

    // use account.age = xxx to call
    set age(value: number) {
        if (value >= 0) {
            this._age = value
        }
    }

}

const customerA = new Account("cA", 99)
customerA.expense(1000)
customerA.income(1000)
customerA.age = 33
console.log(customerA.age);
console.log(customerA)

```



### 泛型

```typescript
// 泛型
function fn<T, K>(a: T, b: K): T {
    console.log(b);
    return a
}

// 小范围
interface Inter {
    length: number
}

function fn2<T extends Inter>(a: T): number {
    return a.length
}

// 泛型for class
class MyClass2<T> {
    name: T
    constructor(name: T) {
        this.name = name
    }
}

const mc = new MyClass2<string>("myclass 2")
```



## Mini Practice

> **贪食蛇**

![image-20220904150740635](C:\coding\typeScript\images\image-20220904150740635.png)

### Technique

> Webpack
>
> Babel
>
> CSS (Less) - `npm i -D less less-loader css-loader style-loader`
>
> - webpack.config.js
>
>   ```js
>   ...
>     
>   // code all webpack config here
>   module.exports = {
>       ...
>       // 指定webpack打包时要使用的模块
>       module: {
>           // 指定要加载的规则
>           rules: [
>               ...
>               //设置less文件的处理
>               {
>                   test: /\.less$/,
>                   use: [
>                       "style-loader",
>                       "css-loader",
>                       "less-loader"
>                   ]
>               }
>           ]
>       },
>       ...
>   }
>   ```
>
> CSS compatibility (postcss) - `npm i -D postcss postcss-loader postcss-preset-env`



### Problem faced

> document.addKeyListener(.., this.fn), some of the part in this.fn not functioning:

```js
document.addEventListener('keydown', this.fn.bind(this))
```



## Reference for future

### 立即函数

```js
(function() {
    ...
})()
```

### **ShortCut for if**

> If "isLive", then console.log

```typescript
this.isLive && console.log(xx)
```


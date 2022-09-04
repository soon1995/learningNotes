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
            },
            //设置less文件的处理
            {
                test: /\.less$/,
                use: [
                    "style-loader",
                    "css-loader",
                    {
                        loader: "postcss-loader",
                        options: {
                            postcssOptions: {
                                plugins: [
                                    [
                                        "postcss-preset-env",
                                        {
                                            browsers: 'last 2 versions'
                                        }
                                    ]
                                ]
                            }
                        }
                    }
                    ,
                    "less-loader"
                ]
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
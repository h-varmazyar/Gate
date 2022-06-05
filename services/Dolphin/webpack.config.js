const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const TerserPlugin = require('terser-webpack-plugin');

module.exports = {
    entry: './scripts/script.js',
    mode: 'production',
    output: {
        path: __dirname + '/public/assets',
        filename: 'script.js',
    },
    module: {
        rules: [
            {
                test: /\.(js)$/,
                loader: 'babel-loader',
            },
            {
                test: /\.scss$/,
                use: [
                    {
                        loader: MiniCssExtractPlugin.loader,
                        options: {
                            publicPath: '/',
                        },
                    },
                    'css-loader',
                    'sass-loader',
                ],
            }
        ],
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: 'style.css',
            ignoreOrder: false,
        }),
    ],
    optimization: {
        minimizer: [new TerserPlugin({
            extractComments: false,
        })],
    },
};
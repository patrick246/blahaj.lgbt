const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin')
const CompressionWebpackPlugin = require('compression-webpack-plugin')

module.exports = {
	mode: 'production',
	entry: './src/index.tsx',
	devtool: 'source-map',
	output: {
		path: path.resolve(__dirname, 'dist'),
		filename: '[name].[contenthash].js',
		clean: true,
		assetModuleFilename: 'assets/[name]-[hash][ext][query]'
	},
	resolve: {
		extensions: ['.tsx', '.ts', '.js'],
	},
	plugins: [
		new HtmlWebpackPlugin({
			title: 'Blahaj',
		}),
		new CompressionWebpackPlugin({
			algorithm: 'gzip',
		})
	],
	module: {
		rules: [
			{
				test: /\.tsx?/i,
				use: 'ts-loader',
				exclude: /node_modules/,
			},
			{
				test: /\.(png|svg|jpg|jpeg|gif|woff|woff2)$/i,
				type: 'asset/resource',
			},
			{
				test: /\.css/i,
				use: ['style-loader', 'css-loader'],
			}
		]
	},
	optimization: {
		usedExports: true,
		splitChunks: {
			chunks: 'all'
		}
	},
	performance: { hints: false },
	devServer: {
		proxy: {
			'/api': 'http://localhost:8080'
		},
		historyApiFallback: true,
		port: 3000,
	}
};
/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		'./pages/**/*.{js,ts,jsx,tsx,mdx}',
		'./components/**/*.{js,ts,jsx,tsx,mdx}',
		'./app/**/*.{js,ts,jsx,tsx,mdx}',
	],
	important: '#__next',
	theme: {
		extend: {
			container: {
				center: true,
			},
			screens: {
				'2xl': '1440px',
			},
			boxShadow: {
				30: ' 0px 14px 28px rgba(0, 0, 0, 0.25),0px 10px 10px rgba(0, 0, 0, 0.22) ',
				38: ' 0px 10px 15px -3px rgba(0, 0, 0, 0.1),0px 4px 6px -2px rgba(0, 0, 0, 0.05) ',
			},
			fontSize: {
				heading: '2.8rem',
				'heading-1': '2.4rem',
				'heading-2': '2.2rem',
				'heading-3': '1.8rem',
				'heading-4': '1.6rem',
				'heading-allcap-lg': '1.8rem',
				'heading-allcap-md': '1.6rem',
				'body-xl': '2.1rem',
				'body-lg': '1.8rem',
				'body-md': '1.6rem',
				'body-sm': '1.4rem',
				'body-xsm': '1.2rem',
			},
			fontWeight: {
				bold: 700,
				regular: 400,
				light: 300,
			},
			spacing: {
				compact: '0.8rem',
				cozy: '1.6rem',
				common: '2rem',
				comfortable: '3.2rem',
			},
			colors: {
				primary: {
					580: {
						DEFAULT: '#BFCBA8',
						80: '#CCD5B9',
						60: '#D9E0CB',
						40: '#E5EADC',
						20: '#F2F5EE',
					},
					625: {
						DEFAULT: '#5B8A72',
						80: '#78A68E',
						60: '#99BCAA',
						40: '#bbd2c7',
						20: '#dde9e3',
					},
					5555: {
						DEFAULT: '#56776C',
						80: '#71998c',
						60: '#95b2a8',
						40: '#b8ccc5',
						20: '#dce5e2',
					},
					418: {
						DEFAULT: '#464F41',
						80: '#6a7762',
						60: '#8e9c87',
						40: '#b4bdaf',
						20: '#d9ded7',
					},
				},
				secondary: {
					376: '#8DBF2E',
					186: '#CC2647',
					1495: '#FF8B28',
					116: '#F2CE1B',
					100: '#F6EB69',
				},
				'action-link': '#01a049',
				neutral: {
					gray: {
						1: '#FFFFFF',
						2: '#F8F8F8',
						3: '#eeeeee',
						4: '#dcdcdc',
						5: '#cecece',
						6: '#BABABA',
						7: '#A0A0A0',
						8: '#999999',
						9: '#666666',
						10: '#333333',
					},
				},
				status: {
					error: {
						DEFAULT: '#E90000',
						mid: '#FF807F',
						light: '#ffe7e7',
					},
					success: {
						DEFAULT: '#53AF3B',
						light: '#DDEFD8',
					},
					warning: {
						DEFAULT: '#fb9716',
						light: '#fdeede',
					},
				},
			},
		},
	},
	plugins: [],
}

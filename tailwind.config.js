/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./public/views/*.html"],
	theme: {
		extend: {
			fontFamily: {
				sans: ["Iosevka Aile Iaso", "sans-serif"],
				mono: ["Iosevka Curly Iaso", "monospace"],
				serif: ["Iosevka Etoile Iaso", "serif"],
			},
			colors: {
				"light-brown": "#272727",
				"dark-gray": "#151515",
				"light-gray": "#bbb",
			},
			keyframes: {
				typing: {
					'0%': {
						width: "0%",
						opacity: '1'
					},
					'100%': {
						width: "100%",
						opacity: '1'
					}
				},
				blink: {
					'0%': {
						opacity: 0
					},
  					'40%': {
						opacity: 0
					},
  					'50%': {
						opacity: 1
					},
  					'90%': {
						opacity: 1
					},
  					'100%': {
						opacity: 0
					},
				}
			},
			animation: {
				"typing-1": "typing 1s 1s steps(20, end) forwards",
				"typing-2": "typing 1s 4s steps(20, end) forwards",
				"typing-3": "typing 1s 7s steps(20, end) forwards",
				"typing-4": "typing 1s 8s steps(20, end) forwards",

				"blink-1": "blink 1s 1s 2 forwards",
				"blink-2": "blink 1s 4.5s 2 forwards", 
				"blink-4": "blink 1s 7.5s infinite",
			}
		},
	},
	plugins: [],
}

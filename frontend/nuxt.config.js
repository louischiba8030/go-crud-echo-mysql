export default {
	axios: {
		proxy: true,
	},
	buildModules: [
		"@nuxtjs/vuetify"
	],
	modules: [
		"@nuxtjs/axios",
		"@nuxtjs/proxy"
	],
	proxy: {
		'/api/': {
			target: 'http://localhost:8003',
			//pathRewrite: {'^/api': '/'}
		}
	},
	components: true
}

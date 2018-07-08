module.exports = {
  pwa: {
    workboxPluginMode: 'InjectManifest',
    name: 'RetroMemo',
    workboxOptions: {
      "globDirectory": "public/",
      "globPatterns": [
          "**/*.{html,ttf,woff,woff2,ico,js,css}"
      ],
      "swSrc": "./src/sw.js",
      "swDest": "./public/sw.js",
    }
  },

  lintOnSave: true
}

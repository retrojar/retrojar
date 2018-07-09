const CACHE = 'cache-update-and-refresh'

self.addEventListener('install', evt => {
  evt.waitUntil(caches.open(CACHE).then(cache => {
    cache.addAll([
      './controlled.html',
      './img'
    ])
  }))
})

self.addEventListener('fetch', evt => {
  evt.respondWith(fromCache(evt.request))
  evt.waitUntil(
    update(evt.request)
      .then(refresh)
  )
})

function fromCache (request) {
  return caches.open(CACHE).then(cache => cache.match(request))
}

function update (request) {
  return caches.open(CACHE).then(cache => {
    return fetch(request).then(response => {
      return cache.put(request, response.clone()).then(() => response)
    })
  })
}

function refresh (response) {
  return self.clients.matchAll().then(clients => {
    clients.forEach(client => {
      let message = {
        type: 'refresh',
        url: response.url,
        eTag: response.headers.get('ETag')
      }
      client.postMessage(JSON.stringify(message))
    })
  })
}

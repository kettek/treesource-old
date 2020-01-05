window.jsApp = {
  title: 'treesource',
  el: null,
  commandIndex: 1,
  pendingCallbacks: {},
}

jsApp.handleEvent = function(name, payload) {
  if (payload.CommandIndex) {
      if (jsApp.pendingCallbacks[payload.CommandIndex]) {
          jsApp.pendingCallbacks[payload.CommandIndex](payload)
          delete jsApp.pendingCallbacks[payload.CommandIndex]
      }
  }
}

jsApp.acquireCommandIndex = function() {
    return jsApp.commandIndex++
}

jsApp.sendMessage = function(eventName, eventData, eventCallback) {
  eventData.CommandIndex = jsApp.acquireCommandIndex()
  jsApp.pendingCallbacks[eventData.CommandIndex] = eventCallback
  app.handleEvent(eventName, eventData)
}

jsApp.setupMenu = function() {
  document.getElementById('jsApp__menu__init').onclick = function(e) {
    jsApp.sendMessage('init', {test: 123}, function(data) {
        alert(JSON.stringify(data))
    })
  }
  document.getElementById('jsApp__menu__sync').onclick = function(e) {
    jsApp.sendMessage('sync', {Resync: true}, function(data) {
      alert(JSON.stringify(data))
    })
  }
}

jsApp.setupTabs = function() {
  let tabs = document.getElementById('jsApp__tabs').getElementsByTagName('button')
  for (let i = 0; i < tabs.length; i++) {
    tabs[i].onclick = function(e) {
      let tabs = document.getElementById('jsApp__tabs').getElementsByTagName('button')
      for (let i = 0; i < tabs.length; i++) {
        if (tabs[i].className == 'selected') {
          tabs[i].className = ''
        }
      }
      let sections = document.getElementsByTagName('section')
      for (let i = 0; i < sections.length; i++) {
        sections[i].className = ''
        if (sections[i].id.split('__section')[1] == e.target.id.split('__tab')[1]) {
          sections[i].className = 'selected'
        }
      }
      e.target.className = 'selected'
    }
  }
}

jsApp.setupSearch = function() {
  let searchButton = document.getElementById('jsApp__sectionSearch__button')
  let searchInput = document.getElementById('jsApp__sectionSearch__input')
  searchButton.onclick = function(e) {
    app.handleEvent('search', {SearchString: searchInput.value})
  }
}

jsApp.init = function() {
  jsApp.el = document.getElementById('jsApp')
  jsApp.setupMenu()
  jsApp.setupTabs()
  jsApp.setupSearch()
}

if (/Linux|MSIE|Trident/.test(window.navigator.userAgent)) {
  jsApp.init()
} else {
  window.addEventListener('DOMContentLoaded', jsApp.init)
}

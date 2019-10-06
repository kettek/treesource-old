window.jsApp = {
  title: 'treesource',
  el: null,
}

jsApp.handleEvent = function(name, payload) {
  alert(name)
  alert(payload)
}

jsApp.setupMenu = function() {
  document.getElementById('jsApp__menu__init').onclick = function(e) {
    app.handleEvent('init', {})
  }
  document.getElementById('jsApp__menu__sync').onclick = function(e) {
    app.handleEvent('sync', {})
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

jsApp.init = function() {
  jsApp.el = document.getElementById('jsApp')
  jsApp.setupMenu()
  jsApp.setupTabs()
}

window.addEventListener('DOMContentLoaded', jsApp.init)
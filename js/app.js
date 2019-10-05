const jsApp = {
  title: 'treesource',
  el: null,
}

jsApp.init = () => {
  jsApp.el = createElement('div')
  jsApp.el.id = 'jsApp'
  document.getElementById('body').appendChild(jsApp.el)
}

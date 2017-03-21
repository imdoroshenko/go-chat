const socket = new WebSocket('ws://127.0.0.1:3000')
socket.addEventListener('open', (e) => {
  console.log(e)
})

const messageTpl = document.createElement('div')

function main() {
  const form = document.querySelector('#form'),
        message = document.querySelector('#message'),
        feed = document.querySelector('#feed'),
        channel = document.querySelector('#channel')

  channel.addEventListener('change', (e) => {
    changeChannel(feed, channel)
  })

  form.addEventListener('submit', (e) => {
    console.log('submit', message.value, channel.value)
    sendMessage(message, channel)
    e.preventDefault()
  })

  socket.addEventListener('message', (e) => {
    console.log('message', e)
    renderMessage(feed, e.data)
  })
}

function changeChannel(feed, channel) {
  while(feed.firstChild){
    feed.removeChild(feed.firstChild)
  }
  socket.send(JSON.stringify({
    resource: 'chn',
    method: 'join',
    channel: {
      name: channel.value
    }
  }))
}

function renderMessage(container, data) {
  const messageNode = messageTpl.cloneNode(true),
        message = JSON.parse(data)
  switch (message.resource) {
    case 'msg':
      messageNode.appendChild(document.createTextNode(message.message.value))
      container.appendChild(messageNode)
      break
    case 'chn':
      messageNode.appendChild(document.createTextNode("Channel changed to " + message.channel.name))
      container.appendChild(messageNode)
      break
  }
  container.scrollTop = container.scrollHeight
}

function sendMessage(message, channel) {
  socket.send(JSON.stringify({
    resource: 'msg',
    method: 'send',
    message: {
      channel: channel.value,
      value: message.value
    }
  }))
  message.value = ''
}

function Robot() {
  this._inProgress = false
}

Robot.prototype.start = function() {

  if (this._inProgress) {
    return
  }
  this._inProgress = true

  ;(function lambda(r){
    if (r._inProgress) {
      sendMessage({value: Math.random().toString(36).substring(7)}, {value:'default'})
      window.setTimeout(lambda, 10, r)
    }
  })(this)
}


Robot.prototype.stop = function() {
  this._inProgress = false
}

const robot = new Robot()

